package domain

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/math/decimal"
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

func (i *Income) GetCursor(chainCode string) *deposit.IncomeCursor {
	cursor, err := i.depositRepo.GetIncomeCursor(chainCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cursor = deposit.NewIncomeCursor(&deposit.IncomeCursorDTO{ChainCode: chainCode})
		}
	}
	return cursor
}

func (i *Income) SaveCursor(cursor *deposit.IncomeCursor) error {
	return i.depositRepo.SaveIncomeCursor(cursor)
}

func (i *Income) ScanDeposits(chainCode string, xxx interface{}) error {

	var deps []*deposit.Deposit

	deps, err := i.scanDeposits(chainCode, xxx)
	if err != nil {
		return err
	}

	deps, err = i.filterDeposits(deps)
	if err != nil {
		return err
	}

	err = i.saveDeposits(deps)

	return err
}

func (i *Income) scanDeposits(chainCode string, xxx interface{}) ([]*deposit.Deposit, error) {
	var deps []*deposit.Deposit

	cApi, ok := i.onchainSvc.GetChainApi(onchain.Code(chainCode))
	if !ok {
		return deps, fmt.Errorf("chain %s not found", chainCode)
	}

	txs, err := cApi.ScanTxn(xxx)
	if err != nil || len(txs) == 0 {
		return deps, err
	}

	assets, err := i.chainRepo.GetChainAssets(chainCode)
	if err != nil {
		return nil, err
	}
	aITOCodePrecession := make(map[string][2]interface{})
	for _, asset := range assets {
		aITOCodePrecession[asset.Identity()] = [2]interface{}{asset.Code(), asset.Precession()}
	}

	for _, txn := range txs {
		var assetCodePrecession [2]interface{}
		assetCodePrecession, ok = aITOCodePrecession[txn.Identity]
		if !ok {
			continue
		}
		precession, _ := assetCodePrecession[1].(int64)
		amount := deposit.NewAmountVO(txn.Identity, txn.Amount, precession, decimal.Decimal{}).ToAmount()
		depDTO := &deposit.DepositDTO{
			Id:         i.depositRepo.NextId(),
			Chain:      string(txn.Chain),
			Asset:      assetCodePrecession[0].(string),
			TxHash:     txn.TxHash,
			Sender:     txn.Sender,
			Receiver:   txn.Receiver,
			Memo:       txn.Memo,
			Identity:   txn.Identity,
			AmountRaw:  txn.Amount,
			Precession: precession,
			Amount:     amount,
			VOut:       txn.VOut,
			Status:     deposit.DepositStatusPending,
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
