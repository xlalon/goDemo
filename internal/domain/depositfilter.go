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

	chainAddresses := make(map[string][]string)
	for _, dep := range deps {
		chainCode := dep.Chain()
		if _, ok := chainAddresses[chainCode]; !ok {
			chainAddresses[chainCode] = []string{}
		}
		chainAddresses[chainCode] = append(chainAddresses[chainCode], dep.Receiver())
	}
	// retrieve accounts from repo
	chainAddressesSelf := make(map[string]map[string]bool)
	for chainCode, addresses := range chainAddresses {
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

	chainIdentities := make(map[string][]string)
	for _, dep := range deps {
		if _, ok := chainIdentities[dep.Chain()]; !ok {
			chainIdentities[dep.Chain()] = []string{}
		}
		chainIdentities[dep.Chain()] = append(chainIdentities[dep.Chain()], dep.Identity())
	}
	// retrieve identity from repo
	assetsSelf, err := af.chainRepository.GetAssets()
	if err != nil {
		return nil, err
	}
	chainAssetMinAmount := make(map[string]map[string]decimal.Decimal)
	chainIdentitiesSelf := make(map[string]map[string]bool)
	for _, asset := range assetsSelf {
		if _, ok := chainIdentitiesSelf[asset.Chain()]; !ok {
			chainIdentitiesSelf[asset.Chain()] = make(map[string]bool)
		}
		if asset.Setting() != nil {
			chainAssetMinAmount[asset.Chain()][asset.Code()] = asset.Setting().MinDepositAmount().Div(decimal.NewFromInt(100))
		}
		chainIdentitiesSelf[asset.Chain()][asset.Identity()] = true
	}
	// filter deps that identity not our asset
	var result []*deposit.Deposit
	for _, dep := range deps {
		if depositAmount, ok := chainAssetMinAmount[dep.Chain()][dep.Asset()]; ok && depositAmount.GreaterThan(dep.Amount()) {
			continue
		}
		if _, ok := chainIdentitiesSelf[dep.Chain()][dep.Identity()]; !ok {
			continue
		}
		result = append(result, dep)
	}

	return result, nil
}
