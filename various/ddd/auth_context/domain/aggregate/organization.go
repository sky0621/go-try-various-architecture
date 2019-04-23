package aggregate

import (
	"go-ddd/backend/domain/entity"
)

// NewOrganization ...
func NewOrganization(attribute, parent entity.OrganizationAttribute, children []entity.OrganizationAttribute) Organization {
	return &organization{
		attribute: attribute,
		parent:    parent,
		children:  children,
	}
}

// Organization ... 「組織」集約情報
type Organization interface {
	GetAttribute() entity.OrganizationAttribute
	GetParentAttribute() entity.OrganizationAttribute
	GetChildAttributes() []entity.OrganizationAttribute
}

type organization struct {
	// 「組織」属性
	attribute entity.OrganizationAttribute

	// 親「組織」属性
	parent entity.OrganizationAttribute

	// 子「組織」属性群
	children []entity.OrganizationAttribute
}

// GetAttribute ...
func (a *organization) GetAttribute() entity.OrganizationAttribute {
	if a == nil {
		return nil
	}
	return a.attribute
}

// GetParentAttribute ...
func (a *organization) GetParentAttribute() entity.OrganizationAttribute {
	if a == nil {
		return nil
	}
	return a.parent
}

// GetChildAttributes ...
func (a *organization) GetChildAttributes() []entity.OrganizationAttribute {
	if a == nil {
		return nil
	}
	return a.children
}
