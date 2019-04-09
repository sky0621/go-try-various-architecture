package aggregate

// NewUserQueryCondition ...
func NewUserQueryCondition(user []User) UserQueryCondition {
	return &userQueryCondition{user: user}
}

// UserQueryCondition ... 条件に該当する「ユーザー」データを決定するために利用
type UserQueryCondition interface {

	// FIXME 複数の User を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type userQueryCondition struct {
	user []User
}
