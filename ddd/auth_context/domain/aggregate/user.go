package aggregate

import (
	"go-ddd/backend/domain/entity"
)

// NewUser ...
func NewUser(attribute entity.UserAttribute) User {
	return &user{
		attribute: attribute,
	}
}

// User ... 「ユーザー」集約情報
type User interface {
	GetAttribute() entity.UserAttribute
}

type user struct {
	// 「ユーザー」属性
	attribute entity.UserAttribute
}

// GetAttribute ...
func (a *user) GetAttribute() entity.UserAttribute {
	if a == nil {
		return nil
	}
	return a.attribute
}
