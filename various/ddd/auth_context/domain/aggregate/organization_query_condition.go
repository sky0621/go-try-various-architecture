package aggregate

// NewOrganizationQueryCondition ...
func NewOrganizationQueryCondition(organizations []Organization) OrganizationQueryCondition {
	return &organizationQueryCondition{organizations: organizations}
}

// OrganizationQueryCondition ... 条件に該当する「組織」データを決定するために利用
type OrganizationQueryCondition interface {

	// FIXME 複数の Organization を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type organizationQueryCondition struct {
	organizations []Organization
}
