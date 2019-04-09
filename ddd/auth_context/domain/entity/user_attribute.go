package entity

import (
	"errors"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NewUserAttribute ...
func NewUserAttribute(id vo.UserID, userNames vo.UserNames) (UserAttribute, error.ApplicationError) {
	if id.GetVal() == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("id is required"))
	}

	return &userAttribute{
		id:        id,
		userNames: userNames,
	}, nil
}

// UserAttribute ... 「ユーザー」データ定義
type UserAttribute interface {
	GetID() vo.UserID
	GetUserNames() vo.UserNames
}

type userAttribute struct {
	// ユニークに特定するID
	id vo.UserID

	// ユーザー名
	userNames vo.UserNames
}

// GetID ...
func (e *userAttribute) GetID() vo.UserID {
	if e == nil {
		return nil
	}
	return e.id
}

// GetTitle ...
func (e *userAttribute) GetUserNames() vo.UserNames {
	if e == nil {
		return nil
	}
	return e.userNames
}
