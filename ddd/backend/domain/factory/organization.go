package factory

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/entity"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NewOrganizationFactory ...
func NewOrganizationFactory(id string) *OrganizationFactory {
	return &OrganizationFactory{uid: vo.NewUniqueIDByParam(id)}
}

// OrganizationFactory ... 「組織」集約のあらゆる生成方法を担う
type OrganizationFactory struct {
	uid vo.UniqueID
}

// CreateOrganization ...
func (f *OrganizationFactory) CreateOrganization(logicalName, logicalKanaName, physicalName, parentID string) (aggregate.Organization, error.ApplicationError) {
	attribute, err := entity.NewOrganizationAttribute(f.uid, vo.NewOrganizationName(logicalName, logicalKanaName, physicalName))
	if err != nil {
		return nil, err
	}

	return aggregate.NewOrganization(attribute, entity.NewOrganizationAttribute(vo.NewUniqueIDByParam(parentID), nil), nil), nil
}

// CreateOrganizationWithChildren ...
func (f *OrganizationFactory) CreateOrganizationWithChildren(organizationName vo.OrganizationName, parent entity.OrganizationAttribute, children []entity.OrganizationAttribute) (aggregate.Organization, error.ApplicationError) {
	attribute, err := entity.NewOrganizationAttribute(f.uid, organizationName)
	if err != nil {
		return nil, err
	}

	return aggregate.NewOrganization(attribute, parent, children), nil
}
