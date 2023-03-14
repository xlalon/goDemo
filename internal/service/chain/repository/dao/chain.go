package dao

import (
	"github.com/xlalon/golee/internal/service/chain/domain"
	"github.com/xlalon/golee/internal/service/chain/repository/model"
)

func (d *Dao) NewChain(c *domain.ChainDTO) error {
	return d.db.Create(&model.Chain{
		Code:   c.Code,
		Name:   c.Name,
		Status: c.Status,
	}).Error
}

func (d *Dao) GetChains() ([]*domain.Chain, error) {
	var chainsDB []model.Chain
	err := d.db.Find(&chainsDB).Error
	if err != nil {
		return nil, err
	}
	var chains []*domain.Chain
	for _, c := range chainsDB {
		_chain, err1 := d.chainDbToDomain(&c)
		if err1 != nil {
			return nil, err1
		}
		chains = append(chains, _chain)
	}
	return chains, nil
}

func (d *Dao) GetChainByCode(chainCode string) (*domain.Chain, error) {
	chainDB := &model.Chain{}
	err := d.db.Last(chainDB, "code = ?", chainCode).Error
	if err != nil {
		return nil, err
	}
	return d.chainDbToDomain(chainDB)
}

func (d *Dao) GetChainByAsset(assetCode string) (*domain.Chain, error) {
	a, err := d.GetAssetByCode(assetCode)
	if err != nil {
		return nil, err
	}
	return d.GetChainByCode(a.Chain)
}

func (d *Dao) chainDbToDomain(a *model.Chain) (*domain.Chain, error) {
	assets, err := d.GetAssetsByChain(a.Code)
	if err != nil {
		return nil, err
	}
	return &domain.Chain{
		Code:   a.Code,
		Name:   a.Name,
		Status: a.Status,
		Assets: assets,
	}, nil
}
