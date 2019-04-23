package valueobject

// NewUserNames ...
func NewUserNames(name, nameKana HumanName, alias string) UserNames {
	// FIXME: name に対する形式チェック！（桁チェック。）
	// FIXME: nameKana に対する形式チェック！（桁チェック。カタカナのみチェック。）
	// FIXME: alias に対する形式チェック！（桁チェック。半角英数字記号のみチェック。）
	return &userName{
		name:     name,
		nameKana: nameKana,
		alias:    alias,
	}
}

// UserNames ... ユーザー名
type UserNames interface {
	GetName() HumanName
	GetNameKana() HumanName
	GetAlias() string
	Equals(comparison UserNames) bool
}

type userName struct {
	// 姓名
	name HumanName

	// 姓名（カナ）
	nameKana HumanName

	// 別名
	alias string
}

// GetName ...
func (v *userName) GetName() HumanName {
	if v == nil {
		return nil
	}
	return v.name
}

// GetNameKana ...
func (v *userName) GetNameKana() HumanName {
	if v == nil {
		return nil
	}
	return v.nameKana
}

// GetAlias ...
func (v *userName) GetAlias() string {
	if v == nil {
		return ""
	}
	return v.alias
}

// Equals ...
func (v *userName) Equals(comparison UserNames) bool {
	if v == nil || comparison == nil {
		return false
	}
	return v.GetName().Equals(comparison.GetName()) &&
		v.GetNameKana().Equals(comparison.GetNameKana()) &&
		v.GetAlias() == comparison.GetAlias()
}
