package aggregate

// NewUserCommandCondition ...
func NewUserCommandCondition(user []User) UserCommandCondition {
	return &userCommandCondition{user: user}
}

// UserCommandCondition ... 条件に該当する「ユーザー」データを決定するために利用
type UserCommandCondition interface {

	// FIXME 複数の User を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type userCommandCondition struct {
	user []User
}
