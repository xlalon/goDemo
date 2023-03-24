package chainasset

import (
	"time"

	"github.com/xlalon/golee/internal/domain/chainasset/model"
	"github.com/xlalon/golee/pkg/database/mysql"
)

// SaveChain any way better?
func (d *Dao) SaveChain(chainDM *model.Chain) error {
	var createdAt time.Time
	chainDB, err := d.getChainById(chainDM.Id())
	if err == nil && chainDB != nil {
		createdAt = chainDB.CreatedAt
	}
	d.db.Model(&Chain{Model: mysql.Model{ID: chainDM.Id()}}).Save(
		&Chain{
			Model: mysql.Model{
				ID:        chainDM.Id(),
				CreatedAt: createdAt,
			},
			Code:   chainDM.Code(),
			Name:   chainDM.Name(),
			Status: string(chainDM.Status()),
		})
	if chainDM.Assets() != nil && len(chainDM.Assets()) > 0 {
		err = d.SaveAssets(chainDM.Assets())
	}
	return err
}

func (d *Dao) GetChains() ([]*model.Chain, error) {
	return d.chainsDbToDomain(d.getChains())
}

func (d *Dao) GetChainByCode(chainCode string) (*model.Chain, error) {
	chainDB, err := d.getChainByCode(chainCode)
	if err != nil {
		return nil, err
	}
	assets, err := d.GetChainAssets(chainDB.Code)
	if err != nil {
		return nil, err
	}
	return d.chainDbToDomain(chainDB, assets), nil
}

func (d *Dao) GetChainByCodes(chainCodes []string) ([]*model.Chain, error) {
	return d.chainsDbToDomain(d.getChainByCodes(chainCodes))
}

func (d *Dao) GetAssetChains(assetCode string) ([]*model.Chain, error) {
	assetsDB, err := d.getAssetsByCode(assetCode)
	if err != nil {
		return nil, err
	}
	var chainCodes []string
	for _, assetDB := range assetsDB {
		chainCodes = append(chainCodes, assetDB.ChainCode)
	}
	return d.GetChainByCodes(chainCodes)
}

func (d *Dao) getChainById(chainId int64) (*Chain, error) {
	chainDB := &Chain{}
	if err := d.db.Last(chainDB, "id = ?", chainId).Error; err != nil {
		return nil, err
	}
	return chainDB, nil
}

func (d *Dao) getChains() ([]Chain, error) {
	var chainsDB []Chain
	if err := d.db.Find(&chainsDB).Error; err != nil {
		return nil, err
	}
	return chainsDB, nil
}

func (d *Dao) getChainByCode(chainCode string) (*Chain, error) {
	chainDB := &Chain{}
	if err := d.db.Last(chainDB, "code = ?", chainCode).Error; err != nil {
		return nil, err
	}
	return chainDB, nil
}

func (d *Dao) getChainByCodes(chainCodes []string) ([]Chain, error) {
	var chainsDB []Chain
	if err := d.db.Find(&chainsDB, "code IN ?", chainCodes).Error; err != nil {
		return nil, err
	}
	return chainsDB, nil
}

func (d *Dao) chainsDbToDomain(chainsDB []Chain, err error) ([]*model.Chain, error) {
	if err != nil || chainsDB == nil || len(chainsDB) == 0 {
		return nil, err
	}
	var chainsDM []*model.Chain
	for _, chainDB := range chainsDB {
		assets, err1 := d.GetChainAssets(chainDB.Code)
		if err1 != nil {
			return nil, err1
		}
		if chainDM := d.chainDbToDomain(&chainDB, assets); chainDM != nil {
			chainsDM = append(chainsDM, chainDM)
		}
	}
	return chainsDM, nil
}

func (d *Dao) chainDbToDomain(c *Chain, assets []*model.Asset) *model.Chain {
	if c == nil {
		return nil
	}
	return model.ChainFactory(
		&model.ChainDTO{
			Id:     c.ID,
			Code:   c.Code,
			Name:   c.Name,
			Status: c.Status,
		},
		assets,
	)
}
