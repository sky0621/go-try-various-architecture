package valueobject

import "go-ddd/backend/domain/service"

// NewOrganizationID ...
func NewOrganizationID() OrganizationID {
	return &organizationID{val: service.CreateUniqueID()}
}

// NewOrganizationIDByParam ...
func NewOrganizationIDByParam(val string) OrganizationID {
	return &organizationID{val: val}
}

// OrganizationID ... 「組織」をユニークに識別するIDとして利用
type OrganizationID interface {
	GetVal() string
	Equals(comparison OrganizationID) bool
}

type organizationID struct {
	val string
}

// GetVal ...
func (v *organizationID) GetVal() string {
	return v.val
}

// Equals ...
func (v *organizationID) Equals(comparison OrganizationID) bool {
	if v == nil || comparison == nil {
		return false
	}
	return v.GetVal() == comparison.GetVal()
}
