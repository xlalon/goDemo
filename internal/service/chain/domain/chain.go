package domain

// Chain Aggregation
type Chain struct {
	Id int64

	code   string
	name   string
	status Status

	assets []*Asset
}

func ChainFactory(chainDTO *ChainDTO, assetsDM []*Asset) *Chain {
	return &Chain{
		Id:     chainDTO.Id,
		code:   chainDTO.Code,
		name:   chainDTO.Name,
		status: Status(chainDTO.Status),
		assets: assetsDM,
	}
}

func (c *Chain) GetId() int64 {
	return c.Id
}

func (c *Chain) GetCode() string {
	return c.code
}

func (c *Chain) GetName() string {
	return c.name
}

func (c *Chain) GetStatus() Status {
	return c.status
}

func (c *Chain) SetStatus(status Status) {
	c.status = status
}

func (c *Chain) GetAssets() []*Asset {
	return c.assets
}

func (c *Chain) AddAsset(asset *Asset) {
	// todo, duplicated check
	c.assets = append(c.assets, asset)
}

func (c *Chain) RemoveAsset(asset *Asset) {
	var assets []*Asset
	for _, a := range c.GetAssets() {
		if a.GetCode() != asset.GetCode() {
			assets = append(assets, a)
		}
	}
	c.assets = assets
}

func (c *Chain) ToChainDTO() *ChainDTO {
	var assetsDTO []*AssetDTO
	for _, assetDM := range c.GetAssets() {
		assetsDTO = append(assetsDTO, assetDM.ToAssetDTO())
	}
	return &ChainDTO{
		Code:   c.GetCode(),
		Name:   c.GetName(),
		Status: string(c.GetStatus()),
		Assets: assetsDTO,
	}
}
