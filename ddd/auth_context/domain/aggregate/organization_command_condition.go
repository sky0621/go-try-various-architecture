package aggregate

// NewOrganizationCommandCondition ...
func NewOrganizationCommandCondition(organizations []Organization) OrganizationCommandCondition {
	return &organizationCommandCondition{organizations: organizations}
}

// OrganizationCommandCondition ... 条件に該当する「組織」データを決定するために利用
type OrganizationCommandCondition interface {

	// FIXME 複数の Organization を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type organizationCommandCondition struct {
	organizations []Organization
}
