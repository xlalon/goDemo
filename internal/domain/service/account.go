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

func (acct *AccountService) GetAccountBalance(ctx context.Context, chainCode, address, assetCode string) (*wallet.BalanceDTO, error) {
	balance := wallet.NewAssetValue(assetCode, decimal.Zero())
	chainRpc, _ := acct.onChainSvc.GetChainApi(xchain.Chain(chainCode))
	asset, err := acct.chainRepository.GetAssetByCode(chainasset.ChainCode(chainCode), chainasset.AssetCode(assetCode))
	if err != nil {
		return balance.ToAssetValueDTO(), err
	}
	coinValue, err := chainRpc.GetAccountBalance(ctx, xchain.Address(address), xchain.Identity(asset.Identity()))
	if err != nil {
		return balance.ToAssetValueDTO(), err
	}
	if string(coinValue.Identity) == asset.Identity() {
		balance = wallet.NewAssetValue(assetCode, asset.CalculateAmount(coinValue.Amount))
	}
	return balance.ToAssetValueDTO(), nil
}
