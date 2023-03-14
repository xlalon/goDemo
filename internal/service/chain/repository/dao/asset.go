package dao

import (
	"github.com/xlalon/golee/internal/service/chain/domain"
	"github.com/xlalon/golee/internal/service/chain/repository/model"
)

func (d *Dao) NewAsset(a *domain.AssetDTO) error {
	return d.db.Create(&model.Asset{
		ChainCode: a.Chain,
		Code:      a.Code,
		Name:      a.Name,
		Identity:  a.Identity,
		Precision: a.Precession,
		Status:    a.Status,
	}).Error
}

func (d *Dao) GetAssets() ([]*domain.Asset, error) {
	var assetsDB []model.Asset
	err := d.db.Find(&assetsDB).Error
	if err != nil {
		return nil, err
	}
	var assets []*domain.Asset
	for _, a := range assetsDB {
		assets = append(assets, d.assetDbToDomain(&a))
	}
	return assets, nil
}

func (d *Dao) GetAssetByCode(assetCode string) (*domain.Asset, error) {
	assetDB := &model.Asset{}
	if err := d.db.Last(assetDB, "code = ?", assetCode).Error; err != nil {
		return nil, err
	}
	return d.assetDbToDomain(assetDB), nil
}

func (d *Dao) GetAssetByIdentity(chainCode, identity string) (*domain.Asset, error) {
	assetDB := &model.Asset{}
	if err := d.db.Last(assetDB, "chain_code = ? AND identity = ?", chainCode, identity).Error; err != nil {
		return nil, err
	}
	return d.assetDbToDomain(assetDB), nil
}

func (d *Dao) GetAssetsByChain(chainCode string) ([]*domain.Asset, error) {
	var assetsDB []model.Asset
	if err := d.db.Find(&assetsDB, "chain_code = ?", chainCode).Error; err != nil {
		return nil, err
	}
	var assets []*domain.Asset
	for _, a := range assetsDB {
		assets = append(assets, d.assetDbToDomain(&a))
	}
	return assets, nil
}

func (d *Dao) assetDbToDomain(a *model.Asset) *domain.Asset {
	return &domain.Asset{
		Code:       a.Code,
		Name:       a.Name,
		Chain:      a.ChainCode,
		Identity:   a.Identity,
		Precession: a.Precision,
		Status:     a.Status,
	}
}
