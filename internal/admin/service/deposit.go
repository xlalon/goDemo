package service

import (
	"github.com/xlalon/golee/internal/admin/conf"
	"github.com/xlalon/golee/internal/service/chain"
	assetConf "github.com/xlalon/golee/internal/service/chain/conf"
	"github.com/xlalon/golee/internal/service/deposit"
	depositConf "github.com/xlalon/golee/internal/service/deposit/conf"
	"github.com/xlalon/golee/internal/service/wallet"
	walletConf "github.com/xlalon/golee/internal/service/wallet/conf"
)

type DepositService struct {
	chainSvc   *chain.Service
	depositSvc *deposit.Service
	walletSvc  *wallet.Service
}

func NewDepositService(conf *conf.Config) *DepositService {
	return &DepositService{
		chainSvc: chain.NewService(&assetConf.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		}),
		depositSvc: deposit.NewService(&depositConf.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		}),
		walletSvc: wallet.NewService(&walletConf.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		}),
	}
}

func (d *DepositService) GetDepositById(id int64) (interface{}, error) {
	return d.depositSvc.GetDeposit(id)
}

func (d *DepositService) GetDeposits() (interface{}, error) {
	return d.depositSvc.GetDeposits()
}
