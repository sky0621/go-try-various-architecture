package entity

import (
	"errors"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NewOrganizationAttribute ...
func NewOrganizationAttribute(uid vo.UniqueID, organizationName vo.OrganizationName) (OrganizationAttribute, error.ApplicationError) {
	if uid.GetVal() == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("id is required"))
	}

	return &organizationAttribute{
		id:               uid,
		organizationName: organizationName,
	}, nil
}

// Organization ... 「ナレッジ」データ定義
type OrganizationAttribute interface {
	GetID() vo.UniqueID
	GetOrganizationName() vo.OrganizationName
}

// OrganizationAttribute ... 「組織」データ定義
type organizationAttribute struct {
	// ユニークに特定するID
	id vo.UniqueID

	// 組織名称
	organizationName vo.OrganizationName
}

// GetID ...
func (e *organizationAttribute) GetID() vo.UniqueID {
	if e == nil {
		return nil
	}
	return e.id
}

// GetTitle ...
func (e *organizationAttribute) GetOrganizationName() vo.OrganizationName {
	if e == nil {
		return nil
	}
	return e.title
}
