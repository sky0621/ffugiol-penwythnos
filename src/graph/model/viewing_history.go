package model

import "time"

type ViewingHistory struct {
	// UUID
	ID string `json:"id"`
	// 視聴者
	Viewer *Viewer `json:"viewer"`
	// 動画
	Movie *Movie `json:"movie"`
	// 視聴日時
	CreatedAt time.Time `json:"createdAt"`
}

func (ViewingHistory) IsNode() {}
