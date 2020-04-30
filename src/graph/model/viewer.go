package model

type Viewer struct {
	// UUID
	ID string `json:"id"`
	// 名前
	Name string `json:"name"`
	// ニックネーム
	Nickname *string `json:"nickname"`
}

func (Viewer) IsNode() {}
