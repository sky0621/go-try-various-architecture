package valueobject

// NewOrganizationNames ...
func NewOrganizationNames(logicalName, logicalKanaName, alias string) OrganizationNames {
	// FIXME: logicalName に対する形式チェック！（桁チェック。）
	// FIXME: logicalKanaName に対する形式チェック！（桁チェック。カタカナのみチェック。）
	// FIXME: physicalName に対する形式チェック！（桁チェック。半角英数字記号のみチェック。）
	return &organizationNames{
		logicalName:     logicalName,
		logicalKanaName: logicalKanaName,
		alias:           alias,
	}
}

// OrganizationNames ... 組織名
type OrganizationNames interface {
	GetLogicalName() string
	GetLogicalKanaName() string
	GetAlias() string
	Equals(comparison OrganizationNames) bool
}

type organizationNames struct {
	// 論理名
	logicalName string

	// 論理名（カナ）
	logicalKanaName string

	// 別名
	alias string
}

// GetLogicalName ...
func (v *organizationNames) GetLogicalName() string {
	return v.logicalName
}

// GetLogicalKanaName ...
func (v *organizationNames) GetLogicalKanaName() string {
	return v.logicalKanaName
}

// GetAlias ...
func (v *organizationNames) GetAlias() string {
	if v == nil {
		return ""
	}
	return v.alias
}

// Equals ...
func (v *organizationNames) Equals(comparison OrganizationNames) bool {
	if v == nil || comparison == nil {
		return false
	}
	return (v.GetLogicalName() == comparison.GetLogicalName()) &&
		(v.GetLogicalKanaName() == comparison.GetLogicalKanaName()) &&
		(v.GetAlias() == comparison.GetAlias())
}
