package chainasset

import (
	"github.com/xlalon/golee/internal/domain/model"
	"github.com/xlalon/golee/pkg/ecode"
	"strings"
)

type ChainCode string

func (cc ChainCode) Normalize() ChainCode {
	return ChainCode(strings.ToUpper(string(cc)))
}

type ChainStatus string

const (
	ChainStatusOnline  ChainStatus = "ONLINE"
	ChainStatusOffline             = "OFFLINE"
)

// Chain Aggregation
type Chain struct {
	model.IdentifiedDomainObject

	code ChainCode
	name string

	status ChainStatus

	assets []*Asset
}

func ChainFactory(chainDTO *ChainDTO, assets []*Asset) *Chain {
	chain := &Chain{}
	if err := chain.SetId(chainDTO.Id); err != nil {
		return nil
	}
	if err := chain.setCode(chainDTO.Code); err != nil {
		return nil
	}
	if err := chain.setName(chainDTO.Name); err != nil {
		return nil
	}
	if err := chain.setStatus(ChainStatus(chainDTO.Status)); err != nil {
		return nil
	}
	if err := chain.setAssets(assets); err != nil {
		return nil
	}
	return chain
}

func (c *Chain) Code() ChainCode {
	return c.code
}

func (c *Chain) setCode(code ChainCode) error {
	if c.Code() != "" {
		return ecode.ChainCodeChange
	}
	if code == "" {
		return ecode.ChainCodeInvalid
	}
	c.code = code.Normalize()
	return nil
}

func (c *Chain) Name() string {
	return c.name
}

func (c *Chain) setName(name string) error {
	if name == "" {
		return ecode.ChainNameInvalid
	}
	c.name = name
	return nil
}

func (c *Chain) Status() ChainStatus {
	return c.status
}

func (c *Chain) setStatus(status ChainStatus) error {
	if status != ChainStatusOffline && status != ChainStatusOnline {
		return ecode.ChainStatusInvalid
	}
	c.status = status
	return nil
}

func (c *Chain) Assets() []*Asset {
	return c.assets
}

func (c *Chain) setAssets(assets []*Asset) error {
	if assets == nil || len(assets) == 0 {
		return nil
	}
	for _, asset := range assets {
		if asset.Chain() != c.Code() {
			return ecode.AssetInvalid
		}
	}
	c.assets = assets
	return nil
}

func (c *Chain) RegisterAsset(assetCode AssetCode, assetName, identity string, precession int64) (*Asset, error) {
	for _, a := range c.Assets() {
		if a.Code() == assetCode {
			return nil, ecode.AssetExist
		}
	}
	asset := AssetFactory(&AssetDTO{
		Id:         c.NextId(),
		Code:       assetCode,
		Name:       assetName,
		Chain:      c.Code(),
		Identity:   identity,
		Precession: precession,
		Status:     AssetStatusOffline,
	})
	c.assets = append(c.assets, asset)

	return asset, nil
}

func (c *Chain) IsOnline() bool {
	return c.Status() == ChainStatusOnline
}

func (c *Chain) Offline() {
	if c.Status() != ChainStatusOffline {
		_ = c.setStatus(ChainStatusOffline)
	}
}

func (c *Chain) Online() {
	if c.Status() != ChainStatusOnline {
		_ = c.setStatus(ChainStatusOnline)
	}
}

func (c *Chain) ToChainDTO() *ChainDTO {
	var assetsDTO []*AssetDTO
	for _, assetDM := range c.Assets() {
		assetsDTO = append(assetsDTO, assetDM.ToAssetDTO())
	}
	return &ChainDTO{
		Id:     c.Id(),
		Code:   c.Code(),
		Name:   c.Name(),
		Status: string(c.Status()),
		Assets: assetsDTO,
	}
}

type ChainDTO struct {
	Id     int64       `json:"id"`
	Code   ChainCode   `json:"code"`
	Name   string      `json:"name"`
	Status string      `json:"status"`
	Assets []*AssetDTO `json:"assets"`
}
