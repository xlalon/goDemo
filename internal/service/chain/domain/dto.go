package domain

type ChainDTO struct {
	Id     int64       `json:"id"`
	Code   string      `json:"code"`
	Name   string      `json:"name"`
	Status string      `json:"status"`
	Assets []*AssetDTO `json:"assets"`
}

type AssetDTO struct {
	Id         int64  `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Chain      string `json:"chain"`
	Identity   string `json:"identity"`
	Precession int64  `json:"precession"`
	Status     string `json:"status"`
}
