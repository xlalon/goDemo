package model

import (
	"github.com/xlalon/golee/pkg/ecode"
)

// Chain Aggregation
type Chain struct {
	id int64

	code   string
	name   string
	status Status

	assets []*Asset
}

func ChainFactory(chainDTO *ChainDTO, assetsDM []*Asset) *Chain {
	chain := &Chain{}
	if err := chain.setId(chainDTO.Id); err != nil {
		return nil
	}
	if err := chain.setCode(chainDTO.Code); err != nil {
		return nil
	}
	if err := chain.setName(chainDTO.Name); err != nil {
		return nil
	}
	if err := chain.setStatus(Status(chainDTO.Status)); err != nil {
		return nil
	}
	if err := chain.setAssets(assetsDM); err != nil {
		return nil
	}
	return chain
}

func (c *Chain) Id() int64 {
	return c.id
}

func (c *Chain) setId(id int64) error {
	if c.Id() != 0 {
		return ecode.ParameterChangeError
	}
	if id <= 0 {
		return ecode.ParameterInvalidError
	}
	c.id = id
	return nil
}

func (c *Chain) Code() string {
	return c.code
}

func (c *Chain) setCode(code string) error {
	if c.Code() != "" {
		return ecode.ParameterNullError
	}
	c.code = code
	return nil
}

func (c *Chain) Name() string {
	return c.name
}

func (c *Chain) setName(name string) error {
	if name == "" {
		return ecode.ParameterNullError
	}
	c.name = name
	return nil
}

func (c *Chain) Status() Status {
	return c.status
}

func (c *Chain) setStatus(status Status) error {
	if status != ChainStatusOffline && status != ChainStatusOnline {
		return ecode.ParameterInvalidError
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
			return ecode.ParameterInvalidError
		}
	}
	c.assets = assets
	return nil
}

func (c *Chain) AddAsset(asset *Asset) error {
	if asset.Chain() != c.Code() {
		return ecode.AssetInvalid
	}
	// duplicate check
	for _, a := range c.Assets() {
		if a.Code() == asset.Code() {
			return ecode.AssetExist
		}
	}
	c.assets = append(c.assets, asset)
	return nil
}

func (c *Chain) RemoveAsset(asset *Asset) {
	var assets []*Asset
	for _, a := range c.Assets() {
		if a.Code() != asset.Code() {
			assets = append(assets, a)
		}
	}
	c.assets = assets
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
