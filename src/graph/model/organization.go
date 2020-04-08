package model

// 組織
type Organization struct {
	// UUID
	ID string `json:"id"`
	// 名称
	Name string `json:"Name"`
	// 上位組織
	// 下位組織群
}

func (Organization) IsNode() {}

// 組織検索条件
type OrganizationCondition struct {
	// UUID
	ID *string `json:"id"`
	// 名称
	Name string `json:"Name"`
}

// 組織入力情報
type OrganizationInput struct {
	// UUID
	ID *string `json:"id"`
	// 名称
	Name string `json:"Name"`
	// 上位組織ID
	UpperOrganizationID *string `json:"upperOrganizationId"`
	// 下位組織ID群
	LowerOrganizationsIds []*string `json:"lowerOrganizationsIds"`
}
