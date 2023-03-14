package deposit

import (
	"fmt"

	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/service/chain"
	"github.com/xlalon/golee/internal/service/deposit/conf"
	"github.com/xlalon/golee/internal/service/deposit/domain"
	"github.com/xlalon/golee/internal/service/deposit/repository"
	"github.com/xlalon/golee/internal/service/deposit/repository/dao"
	"github.com/xlalon/golee/internal/service/wallet"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Income struct {
	repo       repository.DepositRepository
	onchainSvc *onchain.Service
	chainSvc   *chain.Service
	walletSvc  *wallet.Service
}

func NewIncome(conf *conf.Config, onchainSvc *onchain.Service, chainSvc *chain.Service, walletSvc *wallet.Service) *Income {
	return &Income{
		repo:       dao.New(conf),
		onchainSvc: onchainSvc,
		chainSvc:   chainSvc,
		walletSvc:  walletSvc,
	}
}

func (i *Income) ScanDeposits(chainCode string, xxx interface{}) error {

	var deps []*domain.Deposit

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

func (i *Income) scanDeposits(chainCode string, xxx interface{}) ([]*domain.Deposit, error) {
	var deps []*domain.Deposit

	cApi, ok := i.onchainSvc.GetChainApi(onchain.Code(chainCode))
	if !ok {
		return deps, fmt.Errorf("chain %s not found", chainCode)
	}

	txs, err := cApi.ScanTxn(xxx)
	if err != nil {
		return deps, err
	}

	for _, txn := range txs {
		asset, err1 := i.chainSvc.GetAssetByIdentity(string(txn.Chain), txn.Identity)
		if asset == nil || err1 != nil {
			continue
		}
		amount := domain.NewAmountVO(txn.Identity, txn.Amount, asset.Precession, decimal.Decimal{}).ToAmount()
		depDTO := &domain.DepositDTO{
			Chain:      string(txn.Chain),
			Asset:      asset.Code,
			TxHash:     txn.TxHash,
			Sender:     txn.Sender,
			Receiver:   txn.Receiver,
			Memo:       txn.Memo,
			Identity:   txn.Identity,
			AmountRaw:  txn.Amount,
			Precession: asset.Precession,
			Amount:     amount,
			VOut:       txn.VOut,
		}

		dep := domain.NewDeposit(0, *domain.NewDepositItem(depDTO), domain.DepositStatusPending)

		deps = append(deps, dep)
	}

	return deps, nil
}

func (i *Income) filterDeposits(deps []*domain.Deposit) ([]*domain.Deposit, error) {
	filters := []Filterer{
		NewAmountFilter(i.chainSvc),
		NewAccountFilter(i.walletSvc),
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

func (i *Income) saveDeposits(deps []*domain.Deposit) error {
	for _, dep := range deps {
		err := i.repo.CreateDeposit(
			&domain.DepositDTO{
				Chain:     dep.GetChain(),
				Asset:     dep.GetAsset(),
				TxHash:    dep.GetTxHash(),
				Sender:    dep.GetSender(),
				Receiver:  dep.GetReceiver(),
				Memo:      dep.GetMemo(),
				Identity:  dep.GetIdentity(),
				Amount:    dep.GetAmount(),
				AmountRaw: dep.GetAmountRaw(),
				VOut:      dep.GetVOut(),
				Status:    dep.GetStatus(),
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
