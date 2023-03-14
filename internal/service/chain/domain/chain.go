package domain

type Chain struct {
	Code string

	Name   string
	Status string
	Assets []*Asset
}

func (c *Chain) GetName() string {
	return c.Name
}

func (c *Chain) GetStatus() string {
	return c.Status
}

func (c *Chain) SetStatus(status string) {
	c.Status = status
}

func (c *Chain) GetAssets() []*Asset {
	return c.Assets
}

func (c *Chain) AddAsset(asset *Asset) {
	c.Assets = append(c.Assets, asset)
}

func (c *Chain) RemoveAsset(asset *Asset) {
	var assets []*Asset
	for _, a := range c.Assets {
		if a.Code != asset.Code {
			assets = append(assets, a)
		}
	}
	c.Assets = assets
}

func (c *Chain) ToChainDTO() *ChainDTO {
	return &ChainDTO{
		Code:   c.Code,
		Name:   c.GetName(),
		Status: c.GetStatus(),
		Assets: c.GetAssets(),
	}
}
