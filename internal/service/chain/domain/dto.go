package domain

type ChainDTO struct {
	Code   string   `json:"code"`
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Assets []*Asset `json:"assets"`
}

type AssetDTO struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Chain      string `json:"chain"`
	Identity   string `json:"identity"`
	Precession int64  `json:"precession"`
	Status     string `json:"status"`
}
