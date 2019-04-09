package valueobject

import "go-ddd/backend/domain/service"

// NewUserID ...
func NewUserID() UserID {
	return &userID{val: service.CreateUniqueID()}
}

// NewUserIDByParam ...
func NewUserIDByParam(val string) UserID {
	return &userID{val: val}
}

// UserID ... 「ユーザー」エンティティをユニークに識別するIDとして利用
type UserID interface {
	GetVal() string
	Equals(comparison UserID) bool
}

type userID struct {
	val string
}

// GetVal ...
func (v *userID) GetVal() string {
	if v == nil {
		return ""
	}
	return v.val
}

// Equals ...
func (v *userID) Equals(comparison UserID) bool {
	if v == nil || comparison == nil {
		return false
	}
	return v.GetVal() == comparison.GetVal()
}
