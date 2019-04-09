package valueobject

import "go-ddd/backend/domain/service"

// NewUniqueID ...
func NewUniqueID() UniqueID {
	return &uniqueID{val: service.CreateUniqueID()}
}

// NewUniqueIDByParam ...
func NewUniqueIDByParam(val string) UniqueID {
	return &uniqueID{val: val}
}

// UniqueID ... 各エンティティをユニークに識別するIDとして利用
type UniqueID interface {
	GetVal() string
	Equals(comparison UniqueID) bool
}

type uniqueID struct {
	val string
}

// GetVal ...
func (id *uniqueID) GetVal() string {
	return id.val
}

// Equals ...
func (id *uniqueID) Equals(comparison UniqueID) bool {
	if id == nil || comparison == nil {
		return false
	}
	return id.GetVal() == comparison.GetVal()
}
