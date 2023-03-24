package chainasset

import (
	"time"

	"github.com/xlalon/golee/internal/domain/chainasset/model"
	"github.com/xlalon/golee/pkg/database/mysql"
)

func (d *Dao) SaveAssetSetting(chainCode, assetCode string, settings *model.AssetSetting) error {
	var createdAt time.Time
	id := mysql.NextID()
	if assetSetting, err := d.getAssetSetting(chainCode, assetCode); err == nil && assetSetting != nil {
		id = assetSetting.ID
		createdAt = assetSetting.CreatedAt
	}
	return d.db.Save(&AssetSetting{
		Model: mysql.Model{
			ID:        id,
			CreatedAt: createdAt,
		},
		ChainCode:        chainCode,
		AssetCode:        assetCode,
		MinDepositAmount: settings.MinDepositAmount(),
		WithdrawFee:      settings.WithdrawFee(),
		ToHotThreshold:   settings.ToHotThreshold(),
	}).Error
}

func (d *Dao) GetAssetSetting(chainCode, assetCode string) (*model.AssetSetting, error) {
	return d.assetSettingDBToDM(d.getAssetSetting(chainCode, assetCode))
}

func (d *Dao) GetAssetSettings() ([]*model.AssetSetting, error) {
	return d.assetSettingsDBToDM(d.getAssetSettings())
}

func (d *Dao) GetAssetSettingsByChain(chainCode string) ([]*model.AssetSetting, error) {
	return d.assetSettingsDBToDM(d.getAssetSettingsByChain(chainCode))
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

func (d *Dao) assetSettingsDBToDM(settings []AssetSetting, err error) ([]*model.AssetSetting, error) {
	if err != nil || settings == nil || len(settings) == 0 {
		return nil, err
	}
	var settingsDM []*model.AssetSetting
	for _, setting := range settings {
		if settingDM, _ := d.assetSettingDBToDM(&setting, nil); settingDM != nil {
			settingsDM = append(settingsDM, settingDM)
		}
	}
	return settingsDM, nil
}

func (d *Dao) assetSettingDBToDM(setting *AssetSetting, err error) (*model.AssetSetting, error) {
	if err != nil || setting == nil {
		return nil, err
	}
	return model.NewAssetSetting(setting.MinDepositAmount, setting.WithdrawFee, setting.ToHotThreshold), nil
}