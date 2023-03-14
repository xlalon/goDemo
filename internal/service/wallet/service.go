package wallet

import (
	"fmt"

	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/service/wallet/conf"
	"github.com/xlalon/golee/internal/service/wallet/domain"
	"github.com/xlalon/golee/internal/service/wallet/repository"
	"github.com/xlalon/golee/internal/service/wallet/repository/dao"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Service struct {
	repo       repository.WalletRepository
	onchainSvc *onchain.Service
}

func NewService(conf *conf.Config) *Service {
	return &Service{
		repo:       dao.New(conf),
		onchainSvc: onchain.NewService(),
	}
}

func (s *Service) NewAccount(chain, label string) (*domain.Account, error) {
	chainRpc, _ := s.onchainSvc.GetChainApi(onchain.Code(chain))
	acctNew, err := chainRpc.NewAccount(onchain.Label(label))
	if err != nil || acctNew == nil {
		return nil, err
	}
	acct := &domain.Account{
		Chain:    string(acctNew.Chain),
		Address:  acctNew.Address,
		Label:    string(acctNew.Label),
		Memo:     acctNew.Memo,
		Status:   "VALID",
		Sequence: 0,
		Balances: nil,
	}
	if err1 := s.repo.NewAccount(acct); err1 != nil {
		return nil, err1
	}

	return acct, nil
}

func (s *Service) GetAccount(chain, address string) (*domain.Account, error) {
	acctDB, err := s.repo.GetAccountByChainAddress(chain, address)
	if err != nil {
		return nil, err
	}
	chainRpc, _ := s.onchainSvc.GetChainApi(onchain.Code(chain))
	acctOX, err := chainRpc.GetAccount(acctDB.Address)
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
	return &domain.Account{
		Chain:    chain,
		Address:  address,
		Label:    acctDB.Label,
		Memo:     "",
		Status:   acctDB.Status,
		Sequence: acctOX.Sequence,
		Balances: balances,
	}, nil

}

func (s *Service) GetAccountBalance(chain, address string) ([]*domain.Balance, error) {
	acct, err := s.GetAccount(chain, address)
	if err != nil {
		return nil, err
	}
	return acct.GetSBalances(), nil
}

func (s *Service) GetAccountsByChain(chain string) ([]*domain.Account, error) {
	return s.repo.GetAccountsByChain(chain)
}

func (s *Service) GetAccountsByChainAddresses(chain string, addresses []string) ([]*domain.Account, error) {
	return s.repo.GetAccountsByChainAddresses(chain, addresses)
}

func (s *Service) Transfer(chain, sender, receiver, identity string, amount decimal.Decimal, fee map[string]interface{}, extra interface{}) (*domain.Receipt, error) {
	fromAccount, _ := s.GetAccount(chain, sender)
	receiverAccount, _ := s.GetAccount(chain, receiver)
	feeNM := &domain.Fee{}
	for item, value := range fee {
		switch item {
		case "identity":
			feeNM.Identity = value.(string)
		case "amount":
			feeNM.Amount = value.(decimal.Decimal)
		case "gas":
			feeNM.Gas = value.(int64)
		case "gas_price":
			feeNM.GasPrice = value.(decimal.Decimal)
		}
	}
	if fromAccount.Chain != receiverAccount.Chain {
		return nil, fmt.Errorf("invalid chain param")
	}

	senderAccount := &onchain.Account{
		Chain:    onchain.Code(fromAccount.Chain),
		Address:  fromAccount.Address,
		Label:    onchain.Label(fromAccount.Label),
		Sequence: fromAccount.Sequence,
		Extra:    extra,
	}
	receiverAccount1 := &onchain.Account{
		Chain:   onchain.Code(receiverAccount.Chain),
		Address: receiverAccount.Address,
		Label:   onchain.Label(receiverAccount.Label),
		Memo:    receiverAccount.Memo,
	}

	onchainFee := &onchain.Fee{
		Identity: feeNM.Identity,
		Amount:   feeNM.Amount,
	}

	reqData := &onchain.TransferDTO{
		Sender:   senderAccount,
		Receiver: receiverAccount1,
		Identity: identity,
		Amount:   amount,
		Fee:      onchainFee,
		Extra:    extra,
	}
	chainRpc, _ := s.onchainSvc.GetChainApi(onchain.Code(chain))
	resp, err := chainRpc.Transfer(reqData)
	if err != nil {
		return nil, err
	}
	return &domain.Receipt{
		TxHash: resp.TxHash,
		Status: domain.TxnStatus(resp.Status),
	}, nil
}
