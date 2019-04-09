package valueobject

import "strings"

// NewOrganizationLogicalName ...
func NewOrganizationLogicalName(logicalName string) OrganizationName {
	return NewOrganizationName(logicalName, "", "")
}

// NewOrganizationName ...
func NewOrganizationName(logicalName, logicalKanaName, physicalName string) OrganizationName {
	// FIXME: logicalName に対する形式チェック！（桁チェック。）
	// FIXME: logicalKanaName に対する形式チェック！（桁チェック。カタカナのみチェック。）
	// FIXME: physicalName に対する形式チェック！（桁チェック。半角英数字記号のみチェック。）
	return &organizationName{
		logicalName:     logicalName,
		logicalKanaName: logicalKanaName,
		physicalName:    physicalName,
	}
}

// OrganizationName ... 組織名
type OrganizationName interface {
	GetOrganizationLogicalName() string
	GetOrganizationLogicalKanaName() string
	GetOrganizationPhysicalName() string
	Equals(comparison OrganizationName) bool
}

type organizationName struct {
	// 論理名
	logicalName string

	// 論理名（カナ）
	logicalKanaName string

	// 物理名
	physicalName string
}

// GetOrganizationLogicalName ...
func (v *organizationName) GetOrganizationLogicalName() string {
	return v.logicalName
}

// GetOrganizationLogicalKanaName ...
func (v *organizationName) GetOrganizationLogicalKanaName() string {
	return v.logicalKanaName
}

// GetOrganizationPhysicalName ...
func (v *organizationName) GetOrganizationPhysicalName() string {
	return v.physicalName
}

// Equals ...
func (v *organizationName) Equals(comparison OrganizationName) bool {
	if v == nil || comparison == nil {
		return false
	}
	return (strings.Compare(v.GetOrganizationLogicalName(), comparison.GetOrganizationLogicalName()) != -1) &&
		(strings.Compare(v.GetOrganizationLogicalKanaName(), comparison.GetOrganizationLogicalKanaName()) != -1) &&
		(strings.Compare(v.GetOrganizationPhysicalName(), comparison.GetOrganizationPhysicalName()) != -1)
}
