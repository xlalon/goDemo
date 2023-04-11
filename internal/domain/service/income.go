package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/xlalon/golee/internal/domain/model/account"
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/onchain"
)

type Income struct {
	accountRepository account.AccountRepository
	chainRepository   chainasset.ChainRepository
	depositRepository deposit.DepositRepository
	onChainSvc        *onchain.Service
}

func NewIncome(accountRepository account.AccountRepository, chainRepository chainasset.ChainRepository, depositRepository deposit.DepositRepository, onChainSvc *onchain.Service) *Income {
	return &Income{
		accountRepository: accountRepository,
		chainRepository:   chainRepository,
		depositRepository: depositRepository,
		onChainSvc:        onChainSvc,
	}
}

func (i *Income) ScanDeposits(ctx context.Context, chainCode chainasset.ChainCode) error {

	var deps []*deposit.Deposit

	cursor := i.incomeCursor(chainCode)

	deps, err := i.scanDeposits(ctx, cursor)
	if err != nil {
		return err
	}

	deps, err = i.filterDeposits(deps)
	if err != nil {
		return err
	}

	if err = i.saveDeposits(deps); err != nil {
		return err
	}

	if err = i.saveIncomeCursor(cursor); err != nil {
		return err
	}

	return nil
}

func (i *Income) scanDeposits(ctx context.Context, cursor *onchain.Cursor) ([]*deposit.Deposit, error) {
	var deps []*deposit.Deposit

	cApi, ok := i.onChainSvc.GetChainApi(cursor.Chain)
	if !ok {
		return deps, fmt.Errorf("chain %s not found", cursor.Chain)
	}

	txs, err := cApi.ScanTxn(ctx, cursor)
	if err != nil || len(txs) == 0 {
		return deps, err
	}

	assets, err := i.chainRepository.GetChainAssets(chainasset.ChainCode(cursor.Chain))
	if err != nil {
		return nil, err
	}
	identity2asset := make(map[string]*chainasset.Asset)
	for _, asset := range assets {
		identity2asset[asset.Identity()] = asset
	}

	for _, txn := range txs {
		asset, exist := identity2asset[txn.CoinValue.Identity]
		if !exist {
			continue
		}
		amount := asset.CalculateAmount(txn.CoinValue.Amount)
		depDTO := &deposit.DepositDTO{
			Id:       i.depositRepository.NextId(),
			Chain:    string(txn.TxnId.Chain),
			TxHash:   txn.TxnId.TxHash,
			VOut:     txn.TxnId.VOut,
			Receiver: txn.Receiver.Address,
			Memo:     txn.Receiver.Memo,
			Asset:    string(asset.Code()),
			Amount:   amount,

			Sender:    txn.Sender,
			Height:    txn.Height,
			Timestamp: txn.Timestamp,
			Comment:   txn.Comment,

			Status: deposit.DepositStatusPending,
		}

		deps = append(deps, deposit.DepositFactory(depDTO))
	}

	return deps, nil
}

func (i *Income) filterDeposits(deps []*deposit.Deposit) ([]*deposit.Deposit, error) {
	filters := []Filterer{
		NewAmountFilter(i.chainRepository),
		NewAccountFilter(i.accountRepository),
	}
	result := deps
	var err error
	for _, filter := range filters {
		result, err = filter.Satisfied(result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (i *Income) saveDeposits(deps []*deposit.Deposit) error {
	for _, dep := range deps {
		if err := i.depositRepository.Save(dep); err != nil {
			return err
		}
	}
	return nil
}

func (i *Income) incomeCursor(chainCode chainasset.ChainCode) *onchain.Cursor {
	address, label := "", onchain.AccountDeposit
	chainConf, exist := i.onChainSvc.GetChainConfig(onchain.Code(chainCode))
	if exist {
		address = chainConf.DepositAddress
	}
	cursor, err := i.depositRepository.GetIncomeCursor(string(chainCode), address, string(label))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cursor = deposit.NewIncomeCursor(
				string(chainCode),
				0,
				address,
				string(label),
				"",
				string(onchain.CursorDirectionAsc),
				0)
		}
		return nil
	}
	return onchain.NewCursor(
		onchain.Code(chainCode),
		cursor.Height(),
		cursor.Address(),
		onchain.Label(cursor.Label()),
		cursor.TxHash(),
		onchain.Direction(cursor.Direction()),
		cursor.Index(),
	)
}

func (i *Income) saveIncomeCursor(onChainCursor *onchain.Cursor) error {
	return i.depositRepository.SaveIncomeCursor(
		deposit.NewIncomeCursor(
			string(onChainCursor.Chain),
			onChainCursor.Height,
			onChainCursor.Account.Address,
			string(onChainCursor.Account.Label),
			onChainCursor.TxHash,
			string(onChainCursor.Direction),
			onChainCursor.Index))
}
