package model

// 作品
type Work struct {
	// UUID
	ID string `json:"id"`
	// 作品名
	Name string `json:"name"`
	// 価格（無料は0円）
	Price int `json:"price"`
	// 作成者群（不明な場合もある）
}

func (Work) IsNode() {}

// 作品検索条件
type WorkCondition struct {
	// UUID
	ID *string `json:"id"`
	// 作品名
	Name *string `json:"name"`
	// 価格（無料は0円）
	Price *int `json:"price"`
	// 作成者ID
	WorkHolderID *string `json:"workHolderId"`
}

// 作品入力情報
type WorkInput struct {
	// UUID
	ID *string `json:"id"`
	// 作品名
	Name string `json:"name"`
	// 価格（無料は0円）
	Price int `json:"price"`
	// 作成者ID群
	ItemHolderIds []*string `json:"itemHolderIds"`
}
