package service

import (
	"github.com/xlalon/golee/internal/service/wallet"
	"github.com/xlalon/golee/internal/service/wallet/domain"
)

type AccountService struct {
	walletSvc *wallet.Service
}

func NewAccountService(walletRepo domain.WalletRepository) *AccountService {
	return &AccountService{
		walletSvc: wallet.NewService(walletRepo),
	}
}

func (s *AccountService) NewAccountByChain(chainCode, label string) (interface{}, error) {
	return s.walletSvc.NewAccount(chainCode, label)
}

func (s *AccountService) GetAccountDetail(chainCode, address string) (interface{}, error) {
	return s.walletSvc.GetAccount(chainCode, address)
}

func (s *AccountService) GetAccountsByChain(chainCode string) (interface{}, error) {
	return s.walletSvc.GetAccountsByChain(chainCode)
}

func (s *AccountService) GetAccountBalance(chainCode, address, identity string) (interface{}, error) {
	balance, _ := s.walletSvc.GetAccountBalance(chainCode, address)
	if identity != "" {
		for _, b := range balance {
			if b.Identity == identity {
				return b, nil
			}
		}
		return nil, nil
	}
	return balance, nil
}
