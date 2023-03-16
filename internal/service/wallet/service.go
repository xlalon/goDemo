package wallet

import (
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/service/wallet/conf"
	"github.com/xlalon/golee/internal/service/wallet/domain"
	"github.com/xlalon/golee/internal/service/wallet/repoimpl/dao"
	"github.com/xlalon/golee/pkg/database/mysql"
)

type Service struct {
	walletRepo domain.WalletRepository
	onchainSvc *onchain.Service
}

func NewService(conf *conf.Config) *Service {
	return &Service{
		walletRepo: dao.New(conf),
		onchainSvc: onchain.NewService(),
	}
}

func (s *Service) NewAccount(chain, label string) (*domain.AccountDTO, error) {
	chainRpc, _ := s.onchainSvc.GetChainApi(onchain.Code(chain))
	acctNew, err := chainRpc.NewAccount(onchain.Label(label))
	if err != nil || acctNew == nil {
		return nil, err
	}
	acct := domain.AccountFactory(&domain.AccountDTO{
		Id:       mysql.NextID(),
		Chain:    string(acctNew.Chain),
		Address:  acctNew.Address,
		Label:    string(acctNew.Label),
		Memo:     acctNew.Memo,
		Status:   string(domain.AccountStatusValid),
		Sequence: 0,
		Balances: nil,
	})
	if err1 := s.walletRepo.Save(acct); err1 != nil {
		return nil, err1
	}
	account, err := s.walletRepo.GetAccountByChainAddress(chain, acctNew.Address)
	if err != nil {
		return nil, err
	}

	return account.ToAccountDTO(), nil
}

func (s *Service) GetAccount(chain, address string) (*domain.AccountDTO, error) {
	acctDM, err := s.walletRepo.GetAccountByChainAddress(chain, address)
	if err != nil {
		return nil, err
	}
	chainRpc, _ := s.onchainSvc.GetChainApi(onchain.Code(chain))
	acctOX, err := chainRpc.GetAccount(acctDM.GetAddress())
	if err != nil {
		return nil, err
	}
	var balances []*domain.Balance
	for _, balanceInfo := range acctOX.Balance {
		balances = append(balances, &domain.Balance{
			Identity: balanceInfo.Identity,
			Amount:   balanceInfo.Amount,
		})
	}
	return &domain.AccountDTO{
		Id:       acctDM.GetId(),
		Chain:    acctDM.GetChain(),
		Address:  acctDM.GetAddress(),
		Label:    acctDM.GetLabel(),
		Memo:     "",
		Status:   acctDM.GetStatus(),
		Sequence: acctOX.Sequence,
		Balances: balances,
	}, nil

}

func (s *Service) GetAccountBalance(chain, address string) ([]*domain.Balance, error) {
	acct, err := s.GetAccount(chain, address)
	if err != nil {
		return nil, err
	}
	return acct.Balances, nil
}

func (s *Service) GetAccountsByChain(chain string) ([]*domain.AccountDTO, error) {
	return s.accountsDMToDTO(s.walletRepo.GetAccountsByChain(chain))
}

func (s *Service) GetAccountsByChainAddresses(chain string, addresses []string) ([]*domain.AccountDTO, error) {
	return s.accountsDMToDTO(s.walletRepo.GetAccountsByChainAddresses(chain, addresses))
}

func (s *Service) accountDMToDTO(accountDM *domain.Account, err error) (*domain.AccountDTO, error) {
	if err != nil {
		return nil, err
	}
	return accountDM.ToAccountDTO(), nil
}

func (s *Service) accountsDMToDTO(accountsDM []*domain.Account, err error) ([]*domain.AccountDTO, error) {
	if err != nil {
		return nil, err
	}
	accountsDTO := make([]*domain.AccountDTO, len(accountsDM))
	for _, accountDM := range accountsDM {
		accountsDTO = append(accountsDTO, accountDM.ToAccountDTO())
	}
	return accountsDTO, nil
}
