package service

import (
	"github.com/xlalon/golee/internal/domain/wallet/model"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/database/mysql"
)

type WalletService struct {
	domainRegistry *Registry
}

func NewWalletService() *WalletService {
	return &WalletService{DomainRegistry}
}

func (s *WalletService) NewAccount(chain, label string) (*model.AccountDTO, error) {
	chainRpc, _ := s.domainRegistry.onChainService.GetChainApi(onchain.Code(chain))
	acctNew, err := chainRpc.NewAccount(onchain.Label(label))
	if err != nil || acctNew == nil {
		return nil, err
	}
	acct := model.AccountFactory(&model.AccountDTO{
		Id:       mysql.NextID(),
		Chain:    string(acctNew.Chain),
		Address:  acctNew.Address,
		Label:    string(acctNew.Label),
		Memo:     acctNew.Memo,
		Status:   string(model.AccountStatusValid),
		Sequence: 0,
		Balances: nil,
	})
	if err1 := s.domainRegistry.walletRepository.Save(acct); err1 != nil {
		return nil, err1
	}
	account, err := s.domainRegistry.walletRepository.GetAccountByChainAddress(chain, acctNew.Address)
	if err != nil {
		return nil, err
	}

	return account.ToAccountDTO(), nil
}

func (s *WalletService) GetAccount(chain, address string) (*model.AccountDTO, error) {
	acctDM, err := s.domainRegistry.walletRepository.GetAccountByChainAddress(chain, address)
	if err != nil {
		return nil, err
	}
	chainRpc, _ := s.domainRegistry.onChainService.GetChainApi(onchain.Code(chain))
	acctOX, err := chainRpc.GetAccount(acctDM.Address())
	if err != nil {
		return nil, err
	}
	var balances []*model.Balance
	for _, balanceInfo := range acctOX.Balance {
		balances = append(balances, &model.Balance{
			Identity: balanceInfo.Identity,
			Amount:   balanceInfo.Amount,
		})
	}
	return &model.AccountDTO{
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

func (s *WalletService) GetAccountBalance(chain, address string) ([]*model.Balance, error) {
	acct, err := s.GetAccount(chain, address)
	if err != nil {
		return nil, err
	}
	return acct.Balances, nil
}

func (s *WalletService) GetAccountsByChain(chain string) ([]*model.AccountDTO, error) {
	return s.accountsToDTOs(s.domainRegistry.walletRepository.GetAccountsByChain(chain))
}

func (s *WalletService) GetAccountsByChainAddresses(chain string, addresses []string) ([]*model.AccountDTO, error) {
	return s.accountsToDTOs(s.domainRegistry.walletRepository.GetAccountsByChainAddresses(chain, addresses))
}

func (s *WalletService) accountToDTO(account *model.Account, err error) (*model.AccountDTO, error) {
	if err != nil {
		return nil, err
	}
	return account.ToAccountDTO(), nil
}

func (s *WalletService) accountsToDTOs(accounts []*model.Account, err error) ([]*model.AccountDTO, error) {
	if err != nil {
		return nil, err
	}
	accountsDTO := make([]*model.AccountDTO, 0, len(accounts))
	for _, account := range accounts {
		accountsDTO = append(accountsDTO, account.ToAccountDTO())
	}
	return accountsDTO, nil
}
