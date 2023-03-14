package deposit

import (
	"github.com/xlalon/golee/internal/service/chain"
	"github.com/xlalon/golee/internal/service/deposit/domain"
	"github.com/xlalon/golee/internal/service/wallet"
)

type Filterer interface {
	IsValid(*domain.Deposit) bool
	Satisfied([]*domain.Deposit) ([]*domain.Deposit, error)
}

type FilterSpecification struct{}

func (f *FilterSpecification) IsValid(*domain.Deposit) bool {
	return true
}

func (f *FilterSpecification) Satisfied(deps []*domain.Deposit) ([]*domain.Deposit, error) {
	return deps, nil
}

type AccountFilter struct {
	FilterSpecification
	walletSvc *wallet.Service
}

func NewAccountFilter(walletSvc *wallet.Service) *AccountFilter {
	return &AccountFilter{
		walletSvc: walletSvc,
	}
}

func (af *AccountFilter) Satisfied(deps []*domain.Deposit) ([]*domain.Deposit, error) {

	chainAddresses := make(map[string][]string)
	for _, dep := range deps {
		chainCode := dep.GetChain()
		if _, ok := chainAddresses[chainCode]; !ok {
			chainAddresses[chainCode] = []string{}
		}
		chainAddresses[chainCode] = append(chainAddresses[chainCode], dep.GetReceiver())
	}
	// retrieve accounts from repo
	chainAddressesSelf := make(map[string]map[string]bool)
	for chainCode, addresses := range chainAddresses {
		accounts, err := af.walletSvc.GetAccountsByChainAddresses(chainCode, addresses)
		if err != nil {
			return nil, err
		}
		if _, ok := chainAddressesSelf[chainCode]; !ok {
			chainAddressesSelf[chainCode] = make(map[string]bool)
		}
		for _, account := range accounts {
			chainAddressesSelf[chainCode][account.GetAddress()] = true
		}
	}
	// filter deps that receiver not our account
	var result []*domain.Deposit
	for _, dep := range deps {
		if _, ok := chainAddressesSelf[dep.GetChain()][dep.GetReceiver()]; !ok {
			continue
		}
		result = append(result, dep)
	}

	return result, nil
}

type AmountFilter struct {
	FilterSpecification
	chainSvc *chain.Service
}

func NewAmountFilter(chainSvc *chain.Service) *AmountFilter {
	return &AmountFilter{
		chainSvc: chainSvc,
	}
}

func (af *AmountFilter) Satisfied(deps []*domain.Deposit) ([]*domain.Deposit, error) {

	chainIdentities := make(map[string][]string)
	for _, dep := range deps {
		if _, ok := chainIdentities[dep.GetChain()]; !ok {
			chainIdentities[dep.GetChain()] = []string{}
		}
		chainIdentities[dep.GetChain()] = append(chainIdentities[dep.GetChain()], dep.GetIdentity())
	}
	// retrieve identity from repo
	assetsSelf, err := af.chainSvc.GetAssets()
	if err != nil {
		return nil, err
	}
	chainIdentitiesSelf := make(map[string]map[string]bool)
	for _, asset := range assetsSelf {
		if _, ok := chainIdentitiesSelf[asset.Chain]; !ok {
			chainIdentitiesSelf[asset.Chain] = make(map[string]bool)
		}
		chainIdentitiesSelf[asset.Chain][asset.Identity] = true
	}
	// filter deps that identity not our asset
	var result []*domain.Deposit
	for _, dep := range deps {
		if _, ok := chainIdentitiesSelf[dep.GetChain()][dep.GetIdentity()]; !ok {
			continue
		}
		result = append(result, dep)
	}

	return result, nil
}
