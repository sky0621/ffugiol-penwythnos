package model

// 作成者
type WorkHolder struct {
	// UUID
	ID string `json:"id"`
	// 姓
	FirstName string `json:"firstName"`
	// 名
	LastName string `json:"lastName"`
	// ニックネーム
	Nickname *string `json:"nickname"`
	// 所属組織群
	// 所持作品群
}

func (WorkHolder) IsNode() {}

// 作成者検索条件
type WorkHolderCondition struct {
	// UUID
	ID *string `json:"id"`
	// 姓
	FirstName *string `json:"firstName"`
	// 名
	LastName *string `json:"lastName"`
	// ニックネーム
	Nickname *string `json:"nickname"`
	// 所属組織ID
	OrganizationID *string `json:"organizationId"`
}

// 作成者入力情報
type WorkHolderInput struct {
	// UUID
	ID *string `json:"id"`
	// 姓
	FirstName string `json:"firstName"`
	// 名
	LastName string `json:"lastName"`
	// ニックネーム
	Nickname *string `json:"nickname"`
	// 所属組織ID群（※必ずしも所属する必要はない）
	OrganizationIds []*string `json:"organizationIds"`
}
