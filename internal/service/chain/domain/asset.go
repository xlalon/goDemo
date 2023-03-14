package domain

type Asset struct {
	Code string

	Name       string
	Chain      string
	Identity   string
	Precession int64
	Status     string
}

func (a *Asset) GetName() string {
	return a.Name
}

func (a *Asset) GetChain() string {
	return a.Chain
}

func (a *Asset) GetIdentity() string {
	return a.Identity
}

func (a *Asset) GetPrecession() int64 {
	return a.Precession
}

func (a *Asset) GetStatus() string {
	return a.Status
}

func (a *Asset) SetStatus(status string) {
	a.Status = status
}

func (a *Asset) ToAssetDTO() *AssetDTO {
	return &AssetDTO{
		Code:       a.Code,
		Name:       a.GetName(),
		Chain:      a.GetChain(),
		Identity:   a.GetIdentity(),
		Precession: a.GetPrecession(),
		Status:     a.GetStatus(),
	}
}
