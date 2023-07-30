package chainasset

import (
	"time"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/ecode"
)

// SaveChain any way better?
func (d *Dao) SaveChain(chain *chainasset.Chain) error {
	if chain == nil {
		return ecode.ChainInvalid
	}
	var createdAt time.Time
	chainDB, err := d.getChainById(chain.Id())
	if err == nil && chainDB != nil {
		createdAt = chainDB.CreatedAt
	}
	return d.db.Model(&Chain{Model: mysql.Model{ID: chain.Id()}}).Save(
		&Chain{
			Model: mysql.Model{
				ID:        chain.Id(),
				CreatedAt: createdAt,
			},
			Code:   string(chain.Code()),
			Name:   chain.Name(),
			Status: string(chain.Status()),
		}).Error
	//if chain.Assets() != nil && len(chain.Assets()) > 0 {
	//	err = d.SaveAssets(chain.Assets())
	//}
}

func (d *Dao) GetChains() ([]*chainasset.Chain, error) {
	return d.chainsDbToDomain(d.getChains())
}

func (d *Dao) GetChainByCode(chainCode chainasset.ChainCode) (*chainasset.Chain, error) {
	chainDB, err := d.getChainByCode(string(chainCode))
	if err != nil {
		return nil, err
	}
	assets, err := d.GetChainAssets(chainasset.ChainCode(chainDB.Code))
	if err != nil {
		return nil, err
	}
	return d.chainDbToDomain(chainDB, assets), nil
}

func (d *Dao) GetChainByCodes(chainCodes []chainasset.ChainCode) ([]*chainasset.Chain, error) {
	chainCodesStr := make([]string, 0, len(chainCodes))
	for _, chainCode := range chainCodes {
		chainCodesStr = append(chainCodesStr, string(chainCode))
	}
	return d.chainsDbToDomain(d.getChainByCodes(chainCodesStr))
}

func (d *Dao) GetAssetChains(assetCode chainasset.AssetCode) ([]*chainasset.Chain, error) {
	assetsDB, err := d.getAssetsByCode(string(assetCode))
	if err != nil {
		return nil, err
	}
	var chainCodes []chainasset.ChainCode
	for _, assetDB := range assetsDB {
		chainCodes = append(chainCodes, chainasset.ChainCode(assetDB.ChainCode))
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

func (d *Dao) chainsDbToDomain(chainsDB []Chain, err error) ([]*chainasset.Chain, error) {
	if err != nil || chainsDB == nil || len(chainsDB) == 0 {
		return nil, err
	}
	var chainsDM []*chainasset.Chain
	for _, chainDB := range chainsDB {
		assets, err1 := d.GetChainAssets(chainasset.ChainCode(chainDB.Code))
		if err1 != nil {
			return nil, err1
		}
		if chainDM := d.chainDbToDomain(&chainDB, assets); chainDM != nil {
			chainsDM = append(chainsDM, chainDM)
		}
	}
	return chainsDM, nil
}

func (d *Dao) chainDbToDomain(c *Chain, assets []*chainasset.Asset) *chainasset.Chain {
	if c == nil {
		return nil
	}
	return chainasset.NewChain(
		&chainasset.ChainDTO{
			Id:     c.ID,
			Code:   chainasset.ChainCode(c.Code),
			Name:   c.Name,
			Status: c.Status,
		},
		assets,
	)
}
