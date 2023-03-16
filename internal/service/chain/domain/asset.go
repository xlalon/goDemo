package domain

// Asset Entity
type Asset struct {
	Id int64

	code       string
	name       string
	chain      string
	identity   string
	precession int64
	status     Status
}

func AssetFactory(assetDTO *AssetDTO) *Asset {
	return &Asset{
		Id:         assetDTO.Id,
		code:       assetDTO.Code,
		name:       assetDTO.Name,
		chain:      assetDTO.Chain,
		identity:   assetDTO.Identity,
		precession: assetDTO.Precession,
		status:     Status(assetDTO.Status),
	}
}

func (a *Asset) GetId() int64 {
	return a.Id
}

func (a *Asset) GetCode() string {
	return a.code
}

func (a *Asset) GetName() string {
	return a.name
}

func (a *Asset) GetChain() string {
	return a.chain
}

func (a *Asset) GetIdentity() string {
	return a.identity
}

func (a *Asset) GetPrecession() int64 {
	return a.precession
}

func (a *Asset) GetStatus() Status {
	return a.status
}

func (a *Asset) SetStatus(status Status) {
	a.status = status
}

func (a *Asset) ToAssetDTO() *AssetDTO {
	return &AssetDTO{
		Code:       a.GetCode(),
		Name:       a.GetName(),
		Chain:      a.GetChain(),
		Identity:   a.GetIdentity(),
		Precession: a.GetPrecession(),
		Status:     string(a.GetStatus()),
	}
}
