package deposit

import (
	"fmt"

	"github.com/xlalon/golee/core/model/account"
	"github.com/xlalon/golee/core/model/asset"
	"github.com/xlalon/golee/core/model/chain"
	"github.com/xlalon/golee/core/model/deposit"
)

type Service struct {
	acctRepo  account.Repo
	assetRepo asset.Repo
	chainRepo chain.Repo
	depRepo   deposit.Repo
}

func NewService(acctRepo account.Repo, assetRepo asset.Repo, chainRepo chain.Repo, depRepo deposit.Repo) *Service {
	return &Service{
		acctRepo:  acctRepo,
		assetRepo: assetRepo,
		chainRepo: chainRepo,
		depRepo:   depRepo,
	}
}

func (s *Service) Income(chainCode chain.Code, minHeight, maxHeight int64) error {
	var c *chain.Chain
	var err error
	if c, err = s.chainRepo.GetChainByCode(chainCode); err != nil {
		return err
	}
	deps := s.ScanDeposits(c.Code(), minHeight, maxHeight)

	filters := []Filterer{NewAccountFilter(), NewAmountFilter()}

	deps = s.FilterDeposits(deps, filters)

	for _, dep := range deps {
		if err = s.depRepo.Save(dep); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func (s *Service) ScanDeposits(chainCode chain.Code, minHeight, maxHeight int64) []*deposit.Deposit {
	return nil
}

func (s *Service) FilterDeposits(deps []*deposit.Deposit, filters []Filterer) []*deposit.Deposit {
	var err error
	for _, filter := range filters {
		deps, err = filter.Satisfied(deps)
		if err != nil {
			return nil
		}
	}
	return deps
}
