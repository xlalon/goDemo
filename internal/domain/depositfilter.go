package domain

import (
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Filterer interface {
	IsValid(*deposit.Deposit) bool
	Satisfied([]*deposit.Deposit) ([]*deposit.Deposit, error)
}

type FilterSpecification struct{}

func (f *FilterSpecification) IsValid(*deposit.Deposit) bool {
	return true
}

func (f *FilterSpecification) Satisfied(deps []*deposit.Deposit) ([]*deposit.Deposit, error) {
	return deps, nil
}

type AccountFilter struct {
	FilterSpecification
	walletRepository wallet.WalletRepository
}

func NewAccountFilter(walletRepository wallet.WalletRepository) *AccountFilter {
	return &AccountFilter{
		walletRepository: walletRepository,
	}
}

func (af *AccountFilter) Satisfied(deps []*deposit.Deposit) ([]*deposit.Deposit, error) {

	chainAddresses := make(map[string]map[string]bool)
	for _, dep := range deps {
		chainCode := dep.Chain()
		if _, ok := chainAddresses[chainCode]; !ok {
			chainAddresses[chainCode] = make(map[string]bool)
		}
		chainAddresses[chainCode][dep.Receiver()] = true
	}
	// retrieve accounts from repo
	chainAddressesSelf := make(map[string]map[string]bool)
	for chainCode, addressesM := range chainAddresses {
		var addresses []string
		for address, _ := range addressesM {
			addresses = append(addresses, address)
		}
		accounts, err := af.walletRepository.GetAccountsByChainAddresses(chainCode, addresses)
		if err != nil {
			return nil, err
		}
		if _, ok := chainAddressesSelf[chainCode]; !ok {
			chainAddressesSelf[chainCode] = make(map[string]bool)
		}
		for _, account := range accounts {
			chainAddressesSelf[chainCode][account.Address()] = true
		}
	}
	// filter deps that receiver not our account
	var result []*deposit.Deposit
	for _, dep := range deps {
		if _, ok := chainAddressesSelf[dep.Chain()][dep.Receiver()]; !ok {
			continue
		}
		result = append(result, dep)
	}

	return result, nil
}

type AmountFilter struct {
	FilterSpecification
	chainRepository chainasset.ChainRepository
}

func NewAmountFilter(chainRepository chainasset.ChainRepository) *AmountFilter {
	return &AmountFilter{
		chainRepository: chainRepository,
	}
}

func (af *AmountFilter) Satisfied(deps []*deposit.Deposit) ([]*deposit.Deposit, error) {
	// filter deps that amount not satisfied

	// retrieve assets from repo
	assets, err := af.chainRepository.GetAssets()
	if err != nil {
		return nil, err
	}
	// calculate asset dust amount
	dustAmounts := make(map[chainasset.ChainCode]map[chainasset.AssetCode]decimal.Decimal)
	for _, asset := range assets {
		if asset.Setting() == nil {
			continue
		}
		if _, ok := dustAmounts[asset.Chain()]; !ok {
			dustAmounts[asset.Chain()] = make(map[chainasset.AssetCode]decimal.Decimal)
		}
		dustAmounts[asset.Chain()][asset.Code()] = asset.DustAmount()
	}
	var result []*deposit.Deposit
	for _, dep := range deps {
		cc, as := chainasset.ChainCode(dep.Chain()), chainasset.AssetCode(dep.Asset())
		if dustDepAmount, ok := dustAmounts[cc][as]; ok && dep.Amount().LessThan(dustDepAmount) {
			continue
		}
		result = append(result, dep)
	}

	return result, nil
}
