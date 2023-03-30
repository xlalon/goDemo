package chainasset

import (
	"github.com/xlalon/golee/pkg/ecode"
)

// Asset Entity
type Asset struct {
	id int64

	code       string
	name       string
	chain      string
	identity   string
	precession int64
	status     Status

	setting *AssetSetting
}

func AssetFactory(assetDTO *AssetDTO) *Asset {
	asset := &Asset{}
	if err := asset.setId(assetDTO.Id); err != nil {
		return nil
	}
	if err := asset.setName(assetDTO.Name); err != nil {
		return nil
	}
	if err := asset.setCode(assetDTO.Code); err != nil {
		return nil
	}
	if err := asset.setChain(assetDTO.Chain); err != nil {
		return nil
	}
	if err := asset.setIdentity(assetDTO.Identity); err != nil {
		return nil
	}
	if err := asset.setPrecession(assetDTO.Precession); err != nil {
		return nil
	}
	if err := asset.setStatus(Status(assetDTO.Status)); err != nil {
		return nil
	}
	if assetDTO.Setting != nil {
		if err := asset.SetSetting(NewAssetSetting(
			assetDTO.Setting.MinDepositAmount,
			assetDTO.Setting.WithdrawFee,
			assetDTO.Setting.ToHotThreshold,
		)); err != nil {
			return nil
		}
	}

	return asset
}

func (a *Asset) Id() int64 {
	return a.id
}

func (a *Asset) setId(id int64) error {
	if a.Id() != 0 {
		return ecode.ParameterChangeError
	}
	if id <= 0 {
		return ecode.ParameterInvalidError
	}
	a.id = id
	return nil
}

func (a *Asset) Code() string {
	return a.code
}

func (a *Asset) setCode(code string) error {
	if code == "" {
		return ecode.ParameterNullError
	}
	a.code = code
	return nil
}

func (a *Asset) Name() string {
	return a.name
}

func (a *Asset) setName(name string) error {
	if name == "" {
		return ecode.ParameterNullError
	}
	a.name = name
	return nil
}

func (a *Asset) Chain() string {
	return a.chain
}

func (a *Asset) setChain(chain string) error {
	if a.Chain() != "" {
		return ecode.ParameterChangeError
	}
	if chain == "" {
		return ecode.ParameterNullError
	}
	a.chain = chain
	return nil
}

func (a *Asset) Identity() string {
	return a.identity
}

func (a *Asset) setIdentity(identity string) error {
	if a.Identity() != "" {
		return ecode.ParameterChangeError
	}
	if identity == "" {
		return ecode.ParameterNullError
	}
	a.identity = identity
	return nil
}

func (a *Asset) Precession() int64 {
	return a.precession
}

func (a *Asset) setPrecession(precession int64) error {
	if precession < 0 {
		return ecode.ParameterInvalidError
	}
	a.precession = precession
	return nil
}

func (a *Asset) Status() Status {
	return a.status
}

func (a *Asset) setStatus(status Status) error {
	if status != AssetStatusOffline && status != AssetStatusOnline {
		return ecode.ParameterInvalidError
	}
	a.status = status
	return nil
}

func (a *Asset) Setting() *AssetSetting {
	return a.setting
}

func (a *Asset) SetSetting(setting *AssetSetting) error {
	a.setting = setting
	return nil
}

func (a *Asset) Offline() {
	if a.Status() != AssetStatusOffline {
		_ = a.setStatus(AssetStatusOffline)
	}
}

func (a *Asset) Online() {
	if a.Status() != AssetStatusOnline {
		_ = a.setStatus(AssetStatusOnline)
	}
}

func (a *Asset) IsOnline() bool {
	return a.Status() == AssetStatusOnline
}

func (a *Asset) ToAssetDTO() *AssetDTO {
	return &AssetDTO{
		Id:         a.Id(),
		Code:       a.Code(),
		Name:       a.Name(),
		Chain:      a.Chain(),
		Identity:   a.Identity(),
		Precession: a.Precession(),
		Status:     string(a.Status()),
	}
}
