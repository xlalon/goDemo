package service

import (
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/database/mysql"
)

type WalletService struct {
	Service
}

func NewWalletService() *WalletService {
	return &WalletService{Service{DomainRegistry: domain.DomainRegistry}}
}

func (s *WalletService) NewAccount(chain, label string) (*wallet.AccountDTO, error) {
	chainRpc, _ := s.DomainRegistry.OnChainSvc.GetChainApi(onchain.Code(chain))
	acctNew, err := chainRpc.NewAccount(onchain.Label(label))
	if err != nil || acctNew == nil {
		return nil, err
	}
	acct := wallet.AccountFactory(&wallet.AccountDTO{
		Id:       mysql.NextID(),
		Chain:    string(acctNew.Chain),
		Address:  acctNew.Address,
		Label:    string(acctNew.Label),
		Memo:     acctNew.Memo,
		Status:   string(wallet.AccountStatusValid),
		Sequence: 0,
		Balances: nil,
	})
	if err1 := s.DomainRegistry.WalletRepository.Save(acct); err1 != nil {
		return nil, err1
	}
	account, err := s.DomainRegistry.WalletRepository.GetAccountByChainAddress(chain, acctNew.Address)
	if err != nil {
		return nil, err
	}

	return account.ToAccountDTO(), nil
}

func (s *WalletService) GetAccount(chain, address string) (*wallet.AccountDTO, error) {
	acctDM, err := s.DomainRegistry.WalletRepository.GetAccountByChainAddress(chain, address)
	if err != nil {
		return nil, err
	}
	chainRpc, _ := s.DomainRegistry.OnChainSvc.GetChainApi(onchain.Code(chain))
	acctOX, err := chainRpc.GetAccount(acctDM.Address())
	if err != nil {
		return nil, err
	}
	var balances []*wallet.Balance
	for _, balanceInfo := range acctOX.Balance {
		balances = append(balances, &wallet.Balance{
			Identity: balanceInfo.Identity,
			Amount:   balanceInfo.Amount,
		})
	}
	return &wallet.AccountDTO{
		Id:       acctDM.Id(),
		Chain:    acctDM.Chain(),
		Address:  acctDM.Address(),
		Label:    acctDM.Label(),
		Memo:     "",
		Status:   acctDM.Status(),
		Sequence: acctOX.Sequence,
		Balances: balances,
	}, nil

}

func (s *WalletService) GetAccountBalance(chain, address string) ([]*wallet.Balance, error) {
	acct, err := s.GetAccount(chain, address)
	if err != nil {
		return nil, err
	}
	return acct.Balances, nil
}

func (s *WalletService) GetAccountsByChain(chain string) ([]*wallet.AccountDTO, error) {
	return s.accountsToDTOs(s.DomainRegistry.WalletRepository.GetAccountsByChain(chain))
}

func (s *WalletService) GetAccountsByChainAddresses(chain string, addresses []string) ([]*wallet.AccountDTO, error) {
	return s.accountsToDTOs(s.DomainRegistry.WalletRepository.GetAccountsByChainAddresses(chain, addresses))
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
	for _, account := range accounts {
		accountsDTO = append(accountsDTO, account.ToAccountDTO())
	}
	return accountsDTO, nil
}
