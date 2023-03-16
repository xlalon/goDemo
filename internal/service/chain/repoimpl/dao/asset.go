package dao

import (
	"github.com/xlalon/golee/internal/service/chain/domain"
	"github.com/xlalon/golee/internal/service/chain/repoimpl/model"
)

func (d *Dao) GetAssets() ([]*domain.Asset, error) {
	var assetsDB []model.Asset
	if err := d.db.Find(&assetsDB).Error; err != nil {
		return nil, err
	}
	var assetsDM []*domain.Asset
	for _, assetDB := range assetsDB {
		assetsDM = append(assetsDM, d.assetDbToDomain(&assetDB))
	}
	return assetsDM, nil
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
	assetsDB, err := d.getAssetsByChain(chainCode)
	if err != nil {
		return nil, err
	}
	var assetsDM []*domain.Asset
	for _, assetDM := range assetsDB {
		assetsDM = append(assetsDM, d.assetDbToDomain(&assetDM))
	}
	return assetsDM, nil
}

func (d *Dao) getAssetsByChain(chainCode string) ([]model.Asset, error) {
	var assetsDB []model.Asset
	if err := d.db.Find(&assetsDB, "chain_code = ?", chainCode).Error; err != nil {
		return nil, err
	}
	return assetsDB, nil
}

func (d *Dao) getAssetById(assetId int64) (*model.Asset, error) {
	assetDB := &model.Asset{}
	if err := d.db.Last(assetDB, "id = ?", assetId).Error; err != nil {
		return nil, err
	}
	return assetDB, nil
}

func (d *Dao) assetDbToDomain(a *model.Asset) *domain.Asset {
	return domain.AssetFactory(
		&domain.AssetDTO{
			Id:         a.ID,
			Code:       a.Code,
			Name:       a.Name,
			Chain:      a.ChainCode,
			Identity:   a.Identity,
			Precession: a.Precision,
			Status:     a.Status,
		},
	)
}
