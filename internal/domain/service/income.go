package service

import (
	"context"
	"fmt"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/internal/xchain"
)

type Income struct {
	accountRepository wallet.AccountRepository
	chainRepository   chainasset.ChainRepository
	depositRepository deposit.DepositRepository
	onChainSvc        *xchain.Service
}

func NewIncome(accountRepository wallet.AccountRepository, chainRepository chainasset.ChainRepository, depositRepository deposit.DepositRepository, onChainSvc *xchain.Service) *Income {
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

	deps, err := i.scanDeposits(ctx, chainCode, cursor)
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

func (i *Income) scanDeposits(ctx context.Context, chainCode chainasset.ChainCode, cursor *xchain.Cursor) ([]*deposit.Deposit, error) {
	var deps []*deposit.Deposit

	cApi, ok := i.onChainSvc.GetChainApi(xchain.Chain(chainCode))
	if !ok {
		return deps, fmt.Errorf("chain %s not found", chainCode)
	}

	txs, err := cApi.ScanTransfers(ctx, cursor)
	if err != nil || len(txs) == 0 {
		return deps, err
	}

	assets, err := i.chainRepository.GetChainAssets(chainCode)
	if err != nil {
		return nil, err
	}
	identity2asset := make(map[string]*chainasset.Asset)
	for _, asset := range assets {
		identity2asset[asset.Identity()] = asset
	}

	for _, txn := range txs {
		asset, exist := identity2asset[string(txn.CoinValue.Identity)]
		if !exist {
			continue
		}
		amount := asset.CalculateAmount(txn.CoinValue.Amount)
		depDTO := &deposit.DepositDTO{
			Id:       i.depositRepository.NextId(),
			Chain:    string(txn.Chain),
			TxHash:   txn.TxHash,
			VOut:     txn.VOut,
			Receiver: string(txn.Recipient.Address),
			Memo:     string(txn.Recipient.Memo),
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

func (i *Income) incomeCursor(chainCode chainasset.ChainCode) *xchain.Cursor {
	return &xchain.Cursor{}
}

func (i *Income) saveIncomeCursor(onChainCursor *xchain.Cursor) error {
	return nil
}
