package chainasset

import (
	"time"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/pkg/database/mysql"
)

func (d *Dao) SaveAssetSetting(chainCode chainasset.ChainCode, assetCode chainasset.AssetCode, settings *chainasset.AssetSetting) error {
	var createdAt time.Time
	id := mysql.NextID()
	if assetSetting, err := d.getAssetSetting(string(chainCode), string(assetCode)); err == nil && assetSetting != nil {
		id = assetSetting.ID
		createdAt = assetSetting.CreatedAt
	}
	return d.db.Save(&AssetSetting{
		Model: mysql.Model{
			ID:        id,
			CreatedAt: createdAt,
		},
		ChainCode:        string(chainCode),
		AssetCode:        string(assetCode),
		MinDepositAmount: settings.MinDepositAmount(),
		WithdrawFee:      settings.WithdrawFee(),
		ToHotThreshold:   settings.ToHotThreshold(),
	}).Error
}

func (d *Dao) GetAssetSetting(chainCode chainasset.ChainCode, assetCode chainasset.AssetCode) (*chainasset.AssetSetting, error) {
	return d.assetSettingDBToDM(d.getAssetSetting(string(chainCode), string(assetCode)))
}

func (d *Dao) GetAssetSettings() ([]*chainasset.AssetSetting, error) {
	return d.assetSettingsDBToDM(d.getAssetSettings())
}

func (d *Dao) GetAssetSettingsByChain(chainCode chainasset.ChainCode) ([]*chainasset.AssetSetting, error) {
	return d.assetSettingsDBToDM(d.getAssetSettingsByChain(string(chainCode)))
}

func (d *Dao) getAssetSettings() ([]AssetSetting, error) {
	var assetsSettingDB []AssetSetting
	if err := d.db.Find(&assetsSettingDB).Error; err != nil {
		return nil, err
	}
	return assetsSettingDB, nil
}

func (d *Dao) getAssetSetting(chainCode, assetCode string) (*AssetSetting, error) {
	assetSettingDB := &AssetSetting{}
	if err := d.db.First(assetSettingDB, "chain_code = ? AND asset_code = ?", chainCode, assetCode).Error; err != nil {
		return nil, err
	}
	return assetSettingDB, nil
}

func (d *Dao) getAssetSettingsByChain(chainCode string) ([]AssetSetting, error) {
	var assetsSettingDB []AssetSetting
	if err := d.db.Find(&assetsSettingDB, "chain_code = ?", chainCode).Error; err != nil {
		return nil, err
	}
	return assetsSettingDB, nil
}

func (d *Dao) assetSettingsDBToDM(settings []AssetSetting, err error) ([]*chainasset.AssetSetting, error) {
	if err != nil || settings == nil || len(settings) == 0 {
		return nil, err
	}
	var settingsDM []*chainasset.AssetSetting
	for _, setting := range settings {
		if settingDM, _ := d.assetSettingDBToDM(&setting, nil); settingDM != nil {
			settingsDM = append(settingsDM, settingDM)
		}
	}
	return settingsDM, nil
}

func (d *Dao) assetSettingDBToDM(setting *AssetSetting, err error) (*chainasset.AssetSetting, error) {
	if err != nil || setting == nil {
		return nil, err
	}
	return chainasset.NewAssetSetting(setting.MinDepositAmount, setting.WithdrawFee, setting.ToHotThreshold), nil
}
