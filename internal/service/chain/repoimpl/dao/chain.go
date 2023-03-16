package dao

import (
	"time"

	"github.com/xlalon/golee/internal/service/chain/domain"
	"github.com/xlalon/golee/internal/service/chain/repoimpl/model"
	"github.com/xlalon/golee/pkg/database/mysql"
)

// Save any way better?
func (d *Dao) Save(chainDM *domain.Chain) error {
	var createAt = time.Time{}
	chainDB, err := d.getChainById(chainDM.GetId())
	if err == nil && chainDB != nil {
		createAt = chainDB.CreatedAt
	}
	d.db.Model(&model.Chain{Model: mysql.Model{ID: chainDM.GetId()}}).Save(
		&model.Chain{
			Model: mysql.Model{
				ID:        chainDM.GetId(),
				CreatedAt: createAt,
			},
			Code:   chainDM.GetCode(),
			Name:   chainDM.GetName(),
			Status: string(chainDM.GetStatus()),
		})
	// update assets
	assetIdDMUpdated := make(map[int64]*domain.Asset)
	for _, assetDMUpdated := range chainDM.GetAssets() {
		assetIdDMUpdated[assetDMUpdated.GetId()] = assetDMUpdated
	}
	var assetsDb []model.Asset
	assetsDb, err = d.getAssetsByChain(chainDM.GetCode())
	if err != nil {
		return err
	}
	assetDBCI := make(map[int64]time.Time)
	for _, assetDb := range assetsDb {
		if _, ok := assetIdDMUpdated[assetDb.ID]; !ok {
			d.db.Delete(&model.Asset{}, assetDb.ID)
			delete(assetIdDMUpdated, assetDb.ID)
		} else {
			assetDBCI[assetDb.ID] = assetDb.CreatedAt
		}
	}
	for assetId, assetUpdated := range assetIdDMUpdated {
		d.db.Model(&model.Asset{Model: mysql.Model{ID: assetId}}).Save(&model.Asset{
			Model: mysql.Model{
				ID:        assetId,
				CreatedAt: assetDBCI[assetId],
			},
			ChainCode: assetUpdated.GetChain(),
			Code:      assetUpdated.GetCode(),
			Name:      assetUpdated.GetName(),
			Identity:  assetUpdated.GetIdentity(),
			Precision: assetUpdated.GetPrecession(),
			Status:    string(assetUpdated.GetStatus()),
		})
	}

	return nil
}

func (d *Dao) GetChains() ([]*domain.Chain, error) {
	var chainsDB []model.Chain
	if err := d.db.Find(&chainsDB).Error; err != nil {
		return nil, err
	}
	var chainsDM []*domain.Chain
	for _, c := range chainsDB {
		assetsDM, err := d.GetAssetsByChain(c.Code)
		if err != nil {
			return nil, err
		}
		chainsDM = append(chainsDM, d.chainDbToDomain(&c, assetsDM))
	}
	return chainsDM, nil
}

func (d *Dao) GetChainByCode(chainCode string) (*domain.Chain, error) {
	chainDB := &model.Chain{}
	if err := d.db.Last(chainDB, "code = ?", chainCode).Error; err != nil {
		return nil, err
	}
	assets, err := d.GetAssetsByChain(chainDB.Code)
	if err != nil {
		return nil, err
	}
	return d.chainDbToDomain(chainDB, assets), nil
}

func (d *Dao) GetChainByAsset(assetCode string) (*domain.Chain, error) {
	assetDM, err := d.GetAssetByCode(assetCode)
	if err != nil {
		return nil, err
	}
	return d.GetChainByCode(assetDM.GetChain())
}

func (d *Dao) getChainById(chainId int64) (*model.Chain, error) {
	chainDB := &model.Chain{}
	if err := d.db.Last(chainDB, "id = ?", chainId).Error; err != nil {
		return nil, err
	}
	return chainDB, nil
}

func (d *Dao) chainDbToDomain(chain *model.Chain, assets []*domain.Asset) *domain.Chain {
	return domain.ChainFactory(
		&domain.ChainDTO{
			Id:     chain.ID,
			Code:   chain.Code,
			Name:   chain.Name,
			Status: chain.Status,
			Assets: nil,
		},
		assets,
	)
}
