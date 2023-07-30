package service

import (
	"context"
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/pkg/database/mysql"
)

type WalletService struct {
	Service
}

func NewWalletService() *WalletService {
	return &WalletService{Service{DomainRegistry: domain.DomainRegistry}}
}

func (s *WalletService) NewAccount(ctx context.Context, chain, label string) (*wallet.AccountDTO, error) {
	chainRpc, _ := s.DomainRegistry.OnChainSvc.GetChainApi(xchain.Chain(chain))
	acctNew, err := chainRpc.NewAccount(ctx, xchain.WalletLabel(label))
	if err != nil || acctNew == nil {
		return nil, err
	}
	acct := wallet.NewAccount(&wallet.AccountDTO{
		Id:       mysql.NextID(),
		Chain:    chain,
		Address:  string(acctNew.Address),
		Memo:     string(acctNew.Memo),
		Status:   string(wallet.AccountStatusValid),
		Balances: nil,
	})
	if err1 := s.DomainRegistry.AccountRepository.Save(acct); err1 != nil {
		return nil, err1
	}
	acctDM, err := s.DomainRegistry.AccountRepository.GetAccountByChainAddress(chain, string(acctNew.Address))
	if err != nil {
		return nil, err
	}

	return acctDM.ToAccountDTO(), nil
}

func (s *WalletService) GetAccountBalance(ctx context.Context, chainCode, address, assetCode string) (interface{}, error) {
	return s.DomainRegistry.AccountSvc.GetAccountBalance(ctx, chainCode, address, assetCode)
}

func (s *WalletService) GetAccountsByChain(chain string) ([]*wallet.AccountDTO, error) {
	return s.accountsToDTOs(s.DomainRegistry.AccountRepository.GetAccountsByChain(chain))
}

func (s *WalletService) GetAccountsByChainAddresses(chain string, addresses []string) ([]*wallet.AccountDTO, error) {
	return s.accountsToDTOs(s.DomainRegistry.AccountRepository.GetAccountsByChainAddresses(chain, addresses))
}

func (s *WalletService) accountToDTO(account *wallet.Account, err error) (*wallet.AccountDTO, error) {
	if err != nil {
		return nil, err
	}
	return account.ToAccountDTO(), nil
}

func (s *WalletService) accountsToDTOs(accounts []*wallet.Account, err error) ([]*wallet.AccountDTO, error) {
	if err != nil {
		return nil, err
	}
	accountsDTO := make([]*wallet.AccountDTO, 0, len(accounts))
	for _, acct := range accounts {
		accountsDTO = append(accountsDTO, acct.ToAccountDTO())
	}
	return accountsDTO, nil
}
