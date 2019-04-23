package entity

import (
	"errors"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NewOrganizationAttribute ...
func NewOrganizationAttribute(id vo.OrganizationID, organizationNames vo.OrganizationNames) (OrganizationAttribute, error.ApplicationError) {
	if id.GetVal() == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("id is required"))
	}

	return &organizationAttribute{
		id:                id,
		organizationNames: organizationNames,
	}, nil
}

// OrganizationAttribute ... 「組織」データ定義
type OrganizationAttribute interface {
	GetID() vo.OrganizationID
	GetOrganizationNames() vo.OrganizationNames
}

type organizationAttribute struct {
	// ユニークに特定するID
	id vo.OrganizationID

	// 組織名
	organizationNames vo.OrganizationNames
}

// GetID ...
func (e *organizationAttribute) GetID() vo.OrganizationID {
	if e == nil {
		return nil
	}
	return e.id
}

// GetTitle ...
func (e *organizationAttribute) GetOrganizationNames() vo.OrganizationNames {
	if e == nil {
		return nil
	}
	return e.organizationNames
}
