package domain

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/internal/onchain"
)

type Income struct {
	chainRepo   chainasset.ChainRepository
	depositRepo deposit.DepositRepository
	walletRepo  wallet.WalletRepository

	onchainSvc *onchain.Service
}

func NewIncome(chainRepo chainasset.ChainRepository, depositRepo deposit.DepositRepository, walletRepo wallet.WalletRepository) *Income {
	return &Income{
		chainRepo:   chainRepo,
		depositRepo: depositRepo,
		walletRepo:  walletRepo,

		onchainSvc: onchain.NewService(),
	}
}

func (i *Income) ScanDeposits(chainCode chainasset.ChainCode) error {

	var deps []*deposit.Deposit

	cursor := i.incomeCursor(chainCode)

	deps, err := i.scanDeposits(cursor)
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
		fmt.Println(err)
		return err
	}

	return nil
}

func (i *Income) scanDeposits(cursor *onchain.Cursor) ([]*deposit.Deposit, error) {
	var deps []*deposit.Deposit

	cApi, ok := i.onchainSvc.GetChainApi(cursor.Chain)
	if !ok {
		return deps, fmt.Errorf("chain %s not found", cursor.Chain)
	}

	txs, err := cApi.ScanTxn(cursor)
	if err != nil || len(txs) == 0 {
		return deps, err
	}

	assets, err := i.chainRepo.GetChainAssets(chainasset.ChainCode(cursor.Chain))
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
			Id:       i.depositRepo.NextId(),
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
		NewAmountFilter(i.chainRepo),
		NewAccountFilter(i.walletRepo),
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
		if err := i.depositRepo.Save(dep); err != nil {
			return err
		}
	}
	return nil
}

func (i *Income) incomeCursor(chainCode chainasset.ChainCode) *onchain.Cursor {
	address, label := "", onchain.AccountDeposit
	chainConf, exist := i.onchainSvc.GetChainConfig(onchain.Code(chainCode))
	if exist {
		address = chainConf.DepositAddress
	}
	cursor, err := i.depositRepo.GetIncomeCursor(string(chainCode), address, string(label))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cursor = deposit.NewIncomeCursor(string(chainCode), 0, "", "", "", "", 0)
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
	return i.depositRepo.SaveIncomeCursor(
		deposit.NewIncomeCursor(
			string(onChainCursor.Chain),
			onChainCursor.Height,
			onChainCursor.Account.Address,
			string(onChainCursor.Account.Label),
			onChainCursor.TxHash,
			string(onChainCursor.Direction),
			onChainCursor.Index))
}
