package service

import (
	"context"

	"github.com/xlalon/golee/internal/domain/model/account"
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/math/decimal"
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

func (acct *AccountService) GetAccountDetail(ctx context.Context, chain, address string, withBalances bool) (*account.AccountDTO, error) {
	acctDM, err := acct.accountRepository.GetAccountByChainAddress(chain, address)
	if err != nil {
		return nil, err
	}
	chainRpc, _ := acct.onChainSvc.GetChainApi(onchain.Code(chain))
	acctDTO := &account.AccountDTO{
		Id:      acctDM.Id(),
		Chain:   acctDM.Chain(),
		Address: acctDM.Address(),
		Label:   acctDM.Label(),
		Status:  acctDM.Status(),
	}

	acctOX, err := chainRpc.GetAccount(ctx, address)
	if err == nil {
		acctDTO.Sequence = acctOX.Sequence
	}

	if !withBalances {
		return acctDTO, nil
	}

	assets, err := acct.chainRepository.GetChainAssets(chainasset.ChainCode(chain))
	if err == nil {
		var balances []*account.BalanceDTO
		for _, asset := range assets {
			balance, errBalance := chainRpc.GetBalance(ctx, &onchain.Account{Address: address}, asset.Identity())
			if errBalance == nil {
				balanceInfo := &account.Balance{
					Asset:      string(asset.Code()),
					Identity:   asset.Identity(),
					Precession: asset.Precession(),
					Amount:     asset.CalculateAmount(balance),
					AmountRaw:  balance,
				}
				balances = append(balances, balanceInfo.ToBalanceDTO())
			}
		}
		acctDTO.Balances = balances
	}

	return acctDTO, nil
}

func (acct *AccountService) GetAccountBalance(ctx context.Context, chainCode, address, assetCode string) (decimal.Decimal, error) {
	acctOX := &onchain.Account{
		Chain:   onchain.Code(chainCode),
		Address: address,
	}
	chainRpc, _ := acct.onChainSvc.GetChainApi(acctOX.Chain)
	asset, err := acct.chainRepository.GetAssetByCode(chainasset.ChainCode(chainCode), chainasset.AssetCode(assetCode))
	if err != nil {
		return decimal.Zero(), err
	}
	return chainRpc.GetBalance(ctx, acctOX, asset.Identity())
}
