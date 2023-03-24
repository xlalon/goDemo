package model

//
//import (
//	"errors"
//	"fmt"
//	domain2 "github.com/xlalon/golee/internal/service/chainasset/domain"
//	"gorm.io/gorm"
//
//	"github.com/xlalon/golee/internal/onchain"
//	"github.com/xlalon/golee/internal/service/deposit/domain"
//	"github.com/xlalon/golee/internal/service/wallet"
//	"github.com/xlalon/golee/pkg/math/decimal"
//)
//
//type Income struct {
//	depositRepo domain.DepositRepository
//
//	onchainSvc *onchain.Service
//	chainSvc   *domain2.Service
//	walletSvc  *wallet.Service
//}
//
//func NewIncome(depositRepo domain.DepositRepository, chainSvc *domain2.Service, walletSvc *wallet.Service) *Income {
//	return &Income{
//		depositRepo: depositRepo,
//
//		onchainSvc: onchain.NewService(),
//		chainSvc:   chainSvc,
//		walletSvc:  walletSvc,
//	}
//}
//
//func (i *Income) GetCursor(chainCode string) *domain.IncomeCursor {
//	cursor, err := i.depositRepo.GetIncomeCursor(chainCode)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			cursor = domain.NewIncomeCursor(&domain.IncomeCursorDTO{ChainCode: chainCode})
//		}
//	}
//	return cursor
//}
//
//func (i *Income) SaveCursor(cursor *domain.IncomeCursor) error {
//	return i.depositRepo.SaveIncomeCursor(cursor)
//}
//
//func (i *Income) ScanDeposits(chainCode string, xxx interface{}) error {
//
//	var deps []*domain.Deposit
//
//	deps, err := i.scanDeposits(chainCode, xxx)
//	if err != nil {
//		return err
//	}
//
//	deps, err = i.filterDeposits(deps)
//	if err != nil {
//		return err
//	}
//
//	err = i.saveDeposits(deps)
//
//	return err
//}
//
//func (i *Income) scanDeposits(chainCode string, xxx interface{}) ([]*domain.Deposit, error) {
//	var deps []*domain.Deposit
//
//	cApi, ok := i.onchainSvc.GetChainApi(onchain.Code(chainCode))
//	if !ok {
//		return deps, fmt.Errorf("chain %s not found", chainCode)
//	}
//
//	txs, err := cApi.ScanTxn(xxx)
//	if err != nil || len(txs) == 0 {
//		return deps, err
//	}
//
//	assets, err := i.chainSvc.GetChainAssets(chainCode)
//	if err != nil {
//		return nil, err
//	}
//	aITOCodePrecession := make(map[string][2]interface{})
//	for _, asset := range assets {
//		aITOCodePrecession[asset.Identity] = [2]interface{}{asset.Code, asset.Precession}
//	}
//
//	for _, txn := range txs {
//		var assetCodePrecession [2]interface{}
//		assetCodePrecession, ok = aITOCodePrecession[txn.Identity]
//		if !ok {
//			continue
//		}
//		precession, _ := assetCodePrecession[1].(int64)
//		amount := domain.NewAmountVO(txn.Identity, txn.Amount, precession, decimal.Decimal{}).ToAmount()
//		depDTO := &domain.DepositDTO{
//			Id:         i.depositRepo.NextId(),
//			Chain:      string(txn.Chain),
//			Asset:      assetCodePrecession[0].(string),
//			TxHash:     txn.TxHash,
//			Sender:     txn.Sender,
//			Receiver:   txn.Receiver,
//			Memo:       txn.Memo,
//			Identity:   txn.Identity,
//			AmountRaw:  txn.Amount,
//			Precession: precession,
//			Amount:     amount,
//			VOut:       txn.VOut,
//			Status:     domain.DepositStatusPending,
//		}
//
//		deps = append(deps, domain.DepositFactory(depDTO))
//	}
//
//	return deps, nil
//}
//
//func (i *Income) filterDeposits(deps []*domain.Deposit) ([]*domain.Deposit, error) {
//	filters := []Filterer{
//		NewAmountFilter(i.chainSvc),
//		NewAccountFilter(i.walletSvc),
//	}
//	result := deps
//	var err error
//	for _, filter := range filters {
//		result, err = filter.Satisfied(result)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return result, nil
//}
//
//func (i *Income) saveDeposits(deps []*domain.Deposit) error {
//	for _, dep := range deps {
//		if err := i.depositRepo.Save(dep); err != nil {
//			return err
//		}
//	}
//	return nil
//}
