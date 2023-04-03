package chainasset

import (
	"time"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/ecode"
)

func (d *Dao) SaveAsset(asset *chainasset.Asset) error {
	if asset == nil {
		return ecode.AssetInvalid
	}
	var createdAt time.Time
	assetDb, err := d.getAssetById(asset.Id())
	if err == nil && assetDb != nil {
		createdAt = assetDb.CreatedAt
	}
	return d.db.Model(&Asset{Model: mysql.Model{ID: asset.Id()}}).Save(&Asset{
		Model: mysql.Model{
			ID:        asset.Id(),
			CreatedAt: createdAt,
		},
		ChainCode: string(asset.Chain()),
		Code:      string(asset.Code()),
		Name:      asset.Name(),
		Identity:  asset.Identity(),
		Precision: asset.Precession(),
		Status:    string(asset.Status()),
	}).Error
	//if asset.Setting() != nil {
	//	err = d.SaveAssetSetting(asset.Chain(), asset.Code(), asset.Setting())
	//}
}

func (d *Dao) SaveAssets(assets []*chainasset.Asset) error {
	if len(assets) == 0 {
		return nil
	}
	// update assets
	chainAssets := make(map[chainasset.ChainCode][]*chainasset.Asset)
	for _, asset := range assets {
		chainAssets[asset.Chain()] = append(chainAssets[asset.Chain()], asset)
	}
	for chainCode, assetsDM := range chainAssets {
		if len(assetsDM) == 0 {
			continue
		}
		assetsIdDM := make(map[int64]*chainasset.Asset)
		for _, assetDM := range assetsDM {
			id := assetDM.Id()
			if id == 0 {
				id = mysql.NextID()
			}
			assetsIdDM[id] = assetDM
		}
		assetsDb, err := d.getAssetsByChain(string(chainCode))
		if err != nil {
			return err
		}
		assetsDBIdCT := make(map[int64]time.Time)
		for _, assetDb := range assetsDb {
			if _, ok := assetsIdDM[assetDb.ID]; !ok {
				d.db.Delete(&Asset{}, assetDb.ID)
				delete(assetsIdDM, assetDb.ID)
			} else {
				assetsDBIdCT[assetDb.ID] = assetDb.CreatedAt
			}
		}
		for assetId, _asset := range assetsIdDM {
			if err = d.db.Model(&Asset{Model: mysql.Model{ID: assetId}}).Save(&Asset{
				Model: mysql.Model{
					ID:        assetId,
					CreatedAt: assetsDBIdCT[assetId],
				},
				ChainCode: string(_asset.Chain()),
				Code:      string(_asset.Code()),
				Name:      _asset.Name(),
				Identity:  _asset.Identity(),
				Precision: _asset.Precession(),
				Status:    string(_asset.Status()),
			}).Error; err != nil {
				return err
			}
			if _asset.Setting() != nil {
				if err = d.SaveAssetSetting(_asset.Chain(), _asset.Code(), _asset.Setting()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (d *Dao) GetAssets() ([]*chainasset.Asset, error) {
	return d.assetsDbToDomain(d.getAssets())
}

func (d *Dao) GetAssetsByCode(assetCode chainasset.AssetCode) ([]*chainasset.Asset, error) {
	return d.assetsDbToDomain(d.getAssetsByCode(string(assetCode)))
}

func (d *Dao) GetChainAssets(chainCode chainasset.ChainCode) ([]*chainasset.Asset, error) {
	return d.assetsDbToDomain(d.getAssetsByChain(string(chainCode)))
}

func (d *Dao) GetAssetByCode(chainCode chainasset.ChainCode, assetCode chainasset.AssetCode) (*chainasset.Asset, error) {
	return d.assetDbToDomain(d.getAssetByCode(string(chainCode), string(assetCode)))
}

func (d *Dao) GetAssetByIdentity(chainCode chainasset.ChainCode, identity string) (*chainasset.Asset, error) {
	return d.assetDbToDomain(d.getAssetByIdentity(string(chainCode), identity))
}

func (d *Dao) getAssetById(assetId int64) (*Asset, error) {
	assetDB := &Asset{}
	if err := d.db.First(assetDB, "id = ?", assetId).Error; err != nil {
		return nil, err
	}
	return assetDB, nil
}

func (d *Dao) getAssets() ([]Asset, error) {
	var assetsDB []Asset
	if err := d.db.Find(&assetsDB).Error; err != nil {
		return nil, err
	}
	return assetsDB, nil
}

func (d *Dao) getAssetsByCode(assetCode string) ([]Asset, error) {
	var assetsDB []Asset
	if err := d.db.Find(&assetsDB, "code = ?", assetCode).Error; err != nil {
		return nil, err
	}
	return assetsDB, nil
}

func (d *Dao) getAssetsByChain(chainCode string) ([]Asset, error) {
	var assetsDB []Asset
	if err := d.db.Find(&assetsDB, "chain_code = ?", chainCode).Error; err != nil {
		return nil, err
	}
	return assetsDB, nil
}

func (d *Dao) getAssetByCode(chainCode, assetCode string) (*Asset, error) {
	assetDB := &Asset{}
	if err := d.db.First(assetDB, "chain_code = ? AND  code= ?", chainCode, assetCode).Error; err != nil {
		return nil, err
	}
	return assetDB, nil
}

func (d *Dao) getAssetByIdentity(chainCode, identity string) (*Asset, error) {
	assetDB := &Asset{}
	if err := d.db.First(assetDB, "chain_code = ? AND identity = ?", chainCode, identity).Error; err != nil {
		return nil, err
	}
	return assetDB, nil
}

func (d *Dao) assetsDbToDomain(assetsDB []Asset, err error) ([]*chainasset.Asset, error) {
	if err != nil || assetsDB == nil || len(assetsDB) == 0 {
		return nil, err
	}
	var assetsDM []*chainasset.Asset
	for _, assetDB := range assetsDB {
		if assetDM, err := d.assetDbToDomain(&assetDB, nil); err == nil && assetDM != nil {
			assetsDM = append(assetsDM, assetDM)
		}
	}
	return assetsDM, nil
}

func (d *Dao) assetDbToDomain(a *Asset, err error) (*chainasset.Asset, error) {
	if err != nil || a == nil {
		return nil, err
	}
	return chainasset.AssetFactory(
		&chainasset.AssetDTO{
			Id:         a.ID,
			Code:       chainasset.AssetCode(a.Code),
			Name:       a.Name,
			Chain:      chainasset.ChainCode(a.ChainCode),
			Identity:   a.Identity,
			Precession: a.Precision,
			Status:     a.Status,
		},
	), nil
}
