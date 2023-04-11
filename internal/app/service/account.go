package service

import (
	"context"
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/domain/model/account"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/database/mysql"
)

type WalletService struct {
	Service
}

func NewWalletService() *WalletService {
	return &WalletService{Service{DomainRegistry: domain.DomainRegistry}}
}

func (s *WalletService) NewAccount(ctx context.Context, chain, label string) (*account.AccountDTO, error) {
	chainRpc, _ := s.DomainRegistry.OnChainSvc.GetChainApi(onchain.Code(chain))
	acctNew, err := chainRpc.NewAccount(ctx, onchain.Label(label))
	if err != nil || acctNew == nil {
		return nil, err
	}
	acct := account.AccountFactory(&account.AccountDTO{
		Id:       mysql.NextID(),
		Chain:    string(acctNew.Chain),
		Address:  acctNew.Address,
		Label:    string(acctNew.Label),
		Memo:     acctNew.Memo,
		Status:   string(account.AccountStatusValid),
		Sequence: 0,
		Balances: nil,
	})
	if err1 := s.DomainRegistry.AccountRepository.Save(acct); err1 != nil {
		return nil, err1
	}
	acctDM, err := s.DomainRegistry.AccountRepository.GetAccountByChainAddress(chain, acctNew.Address)
	if err != nil {
		return nil, err
	}

	return acctDM.ToAccountDTO(), nil
}

func (s *WalletService) GetAccount(ctx context.Context, chain, address string) (*account.AccountDTO, error) {
	return s.DomainRegistry.AccountSvc.GetAccount(ctx, chain, address)
}

func (s *WalletService) GetAccountBalance(ctx context.Context, chain, address string) ([]*account.BalanceDTO, error) {
	acct, err := s.GetAccount(ctx, chain, address)
	if err != nil {
		return nil, err
	}
	return acct.Balances, nil
}

func (s *WalletService) GetAccountsByChain(chain string) ([]*account.AccountDTO, error) {
	return s.accountsToDTOs(s.DomainRegistry.AccountRepository.GetAccountsByChain(chain))
}

func (s *WalletService) GetAccountsByChainAddresses(chain string, addresses []string) ([]*account.AccountDTO, error) {
	return s.accountsToDTOs(s.DomainRegistry.AccountRepository.GetAccountsByChainAddresses(chain, addresses))
}

func (s *WalletService) accountToDTO(account *account.Account, err error) (*account.AccountDTO, error) {
	if err != nil {
		return nil, err
	}
	return account.ToAccountDTO(), nil
}

func (s *WalletService) accountsToDTOs(accounts []*account.Account, err error) ([]*account.AccountDTO, error) {
	if err != nil {
		return nil, err
	}
	accountsDTO := make([]*account.AccountDTO, 0, len(accounts))
	for _, acct := range accounts {
		accountsDTO = append(accountsDTO, acct.ToAccountDTO())
	}
	return accountsDTO, nil
}
