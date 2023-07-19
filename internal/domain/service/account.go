package service

import (
	"context"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AccountService struct {
	accountRepository wallet.AccountRepository
	chainRepository   chainasset.ChainRepository
	onChainSvc        *xchain.Service
}

func NewAccountService(accountRepository wallet.AccountRepository, chainRepository chainasset.ChainRepository, onChainSvc *xchain.Service) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
		chainRepository:   chainRepository,
		onChainSvc:        onChainSvc,
	}
}

func (acct *AccountService) GetAccountBalance(ctx context.Context, chainCode, address, assetCode string) (xchain.Coin, error) {

	chainRpc, _ := acct.onChainSvc.GetChainApi(xchain.Chain(chainCode))
	asset, err := acct.chainRepository.GetAssetByCode(chainasset.ChainCode(chainCode), chainasset.AssetCode(assetCode))
	if err != nil {
		return xchain.Coin{Identity: xchain.Identity(asset.Identity()), Amount: decimal.Zero()}, err
	}
	return chainRpc.GetAccountBalance(ctx, xchain.Address(address), xchain.Identity(asset.Identity()))
}
