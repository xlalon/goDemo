package service

import (
	"context"
	"github.com/xlalon/golee/internal/domain/model/account"
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/onchain"
)

type AccountService struct {
	accountRepository account.AccountRepository
	chainRepository   chainasset.ChainRepository
	onChainSvc        *onchain.Service
}

func NewAccountService(accountRepository account.AccountRepository, chainRepository chainasset.ChainRepository, onChainSvc *onchain.Service) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
		chainRepository:   chainRepository,
		onChainSvc:        onChainSvc,
	}
}

func (acct *AccountService) GetAccount(ctx context.Context, chain, address string) (*account.AccountDTO, error) {
	acctDM, err := acct.accountRepository.GetAccountByChainAddress(chain, address)
	if err != nil {
		return nil, err
	}
	chainRpc, _ := acct.onChainSvc.GetChainApi(onchain.Code(chain))
	acctOX, err := chainRpc.GetAccount(ctx, acctDM.Address())
	if err != nil {
		return nil, err
	}
	var balances []*account.BalanceDTO
	for _, balanceInfo := range acctOX.Balance {
		asset, errAsset := acct.chainRepository.GetAssetByIdentity(chainasset.ChainCode(chain), balanceInfo.Identity)
		if errAsset != nil {
			continue
		}
		balance := &account.Balance{
			Asset:      string(asset.Code()),
			Identity:   balanceInfo.Identity,
			Precession: asset.Precession(),
			Amount:     asset.CalculateAmount(balanceInfo.Amount),
			AmountRaw:  balanceInfo.Amount,
		}
		balances = append(balances, balance.ToBalanceDTO())
	}

	return &account.AccountDTO{
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
