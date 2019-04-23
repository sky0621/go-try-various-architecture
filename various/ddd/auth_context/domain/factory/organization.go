package factory

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/entity"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NewOrganizationFactory ...
func NewOrganizationFactory(id string) *OrganizationFactory {
	return &OrganizationFactory{uid: vo.NewOrganizationIDByParam(id)}
}

// OrganizationFactory ... 「組織」集約のあらゆる生成方法を担う
type OrganizationFactory struct {
	uid vo.OrganizationID
}

// CreateOrganization ...
func (f *OrganizationFactory) CreateOrganization(logicalName, logicalKanaName, alias, parentID string) (aggregate.Organization, error.ApplicationError) {
	attribute, err := entity.NewOrganizationAttribute(f.uid, vo.NewOrganizationNames(logicalName, logicalKanaName, alias))
	if err != nil {
		return nil, err
	}

	return aggregate.NewOrganization(attribute, entity.NewOrganizationAttribute(vo.NewOrganizationIDByParam(parentID), nil), nil), nil
}

// CreateOrganizationWithChildren ...
func (f *OrganizationFactory) CreateOrganizationWithChildren(organizationName vo.OrganizationNames, parent entity.OrganizationAttribute, children []entity.OrganizationAttribute) (aggregate.Organization, error.ApplicationError) {
	attribute, err := entity.NewOrganizationAttribute(f.uid, organizationName)
	if err != nil {
		return nil, err
	}

	return aggregate.NewOrganization(attribute, parent, children), nil
}
