package model

type Movie struct {
	// UUID
	ID string `json:"id"`
	// 名称
	Name string `json:"name"`
	// 動画URL
	MovieURL string `json:"movieUrl"`
	// 秒数
	Scale int `json:"scale"`
}

func (Movie) IsNode() {}
