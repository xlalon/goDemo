package asset

type Repo interface {
	Save(a *Asset) error

	GetAssetById(id int64) (*Asset, error)
	GetAssetByCode(code Code) (*Asset, error)
	GetAssets() ([]*Asset, error)
}
