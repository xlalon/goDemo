package chainasset

import (
	"strings"

	"github.com/xlalon/golee/internal/domain/model"
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AssetCode string

func (ac AssetCode) Normalize() AssetCode {
	return AssetCode(strings.ToUpper(string(ac)))
}

type AssetStatus string

const (
	AssetStatusOnline  AssetStatus = "ONLINE"
	AssetStatusOffline             = "OFFLINE"
)

// Asset Entity
type Asset struct {
	model.IdentifiedDomainObject

	code AssetCode
	name string

	chain ChainCode

	identity   string
	precession int64

	status AssetStatus

	setting *AssetSetting
}

func NewAsset(assetDTO *AssetDTO) *Asset {
	asset := &Asset{}
	if err := asset.SetId(assetDTO.Id); err != nil {
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
	if err := asset.setStatus(AssetStatus(assetDTO.Status)); err != nil {
		return nil
	}
	if assetDTO.Setting != nil {
		if _, err := asset.ApplySetting(
			assetDTO.Setting.MinDepositAmount,
			assetDTO.Setting.WithdrawFee,
			assetDTO.Setting.ToHotThreshold,
		); err != nil {
			return nil
		}
	}

	return asset
}

func (a *Asset) Code() AssetCode {
	return a.code
}

func (a *Asset) setCode(code AssetCode) error {
	if a.code != "" {
		return ecode.AssetCodeChange
	}
	if code == "" {
		return ecode.AssetCodeInvalid
	}
	a.code = code.Normalize()
	return nil
}

func (a *Asset) Name() string {
	return a.name
}

func (a *Asset) setName(name string) error {
	if name == "" {
		return ecode.AssetNameInvalid
	}
	a.name = name
	return nil
}

func (a *Asset) Chain() ChainCode {
	return a.chain
}

func (a *Asset) setChain(chain ChainCode) error {
	if a.Chain() != "" {
		return ecode.ChainCodeChange
	}
	if chain == "" {
		return ecode.ChainCodeInvalid
	}
	a.chain = chain
	return nil
}

func (a *Asset) Identity() string {
	return a.identity
}

func (a *Asset) setIdentity(identity string) error {
	if a.Identity() != "" {
		return ecode.AssetIdentityChange
	}
	if identity == "" {
		return ecode.AssetIdentityInvalid
	}
	a.identity = identity
	return nil
}

func (a *Asset) Precession() int64 {
	return a.precession
}

func (a *Asset) setPrecession(precession int64) error {
	if precession < 0 {
		return ecode.AssetPrecessionInvalid
	}
	a.precession = precession
	return nil
}

func (a *Asset) Status() AssetStatus {
	return a.status
}

func (a *Asset) setStatus(status AssetStatus) error {
	if status != AssetStatusOffline && status != AssetStatusOnline {
		return ecode.AssetStatusInvalid
	}
	a.status = status
	return nil
}

func (a *Asset) Setting() *AssetSetting {
	return a.setting
}

func (a *Asset) ApplySetting(minDepositAmount, withdrawFee, toHotThreshold decimal.Decimal) (*AssetSetting, error) {
	setting := NewAssetSetting(
		minDepositAmount,
		withdrawFee,
		toHotThreshold,
	)
	if setting == nil {
		return nil, ecode.AssetSettingInvalid
	}
	a.setting = setting
	return setting, nil
}

func (a *Asset) DustAmount() decimal.Decimal {
	if a.Setting() != nil {
		return a.Setting().MinDepositAmount().Div(decimal.NewFromInt(100))
	}
	return decimal.NewFromInt(0)
}

func (a *Asset) CalculateAmount(amountRaw decimal.Decimal) decimal.Decimal {
	return amountRaw.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(a.Precession())))
}

func (a *Asset) CalculateAmountRaw(amount decimal.Decimal) decimal.Decimal {
	return amount.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(a.Precession())))
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

type AssetDTO struct {
	Id         int64     `json:"id"`
	Code       AssetCode `json:"code"`
	Name       string    `json:"name"`
	Chain      ChainCode `json:"chain"`
	Identity   string    `json:"identity"`
	Precession int64     `json:"precession"`
	Status     string    `json:"status"`

	Setting *AssetSettingDTO `json:"-"`
}
