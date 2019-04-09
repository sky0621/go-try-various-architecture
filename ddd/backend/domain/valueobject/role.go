package valueobject

const (
	// SystemAdministrator ... システム管理者
	SystemAdministrator = iota + 1

	// OrganizationManager ... 組織管理者
	OrganizationManager

	// AuthenticationUser ... 認証ユーザー
	AuthenticationUser

	// GuestUser ... ゲストユーザー
	GuestUser
)

// NewRole ...
func NewRole(val int) Role {
	// 事前条件
	assertion := func(v int) bool {
		if v == SystemAdministrator {
			return true
		}
		if v == OrganizationManager {
			return true
		}
		if v == AuthenticationUser {
			return true
		}
		if v == GuestUser {
			return true
		}
		return false
	}
	if !assertion(val) {
		return nil
	}

	return &role{val: val}
}

// NewSystemAdministrator ... システム管理者ロールを生成する
func NewSystemAdministrator() Role {
	return &role{val: SystemAdministrator}
}

// NewOrganizationManager ... 組織管理者ロールを生成する
func NewOrganizationManager() Role {
	return &role{val: OrganizationManager}
}

// NewAuthenticationUser ... 認証ユーザーロールを生成する
func NewAuthenticationUser() Role {
	return &role{val: AuthenticationUser}
}

// NewGuestUser ... ゲストユーザーロールを生成する
func NewGuestUser() Role {
	return &role{val: GuestUser}
}

// Role ... ユーザーの役割
type Role interface {
	GetVal() int
	Equals(comparison Role) bool
	IsSystemAdministrator() bool
	IsOrganizationManager() bool
	IsAuthenticationUser() bool
	IsGuestUser() bool
}

type role struct {
	val int
}

// GetVal ...
func (r *role) GetVal() int {
	return r.val
}

// Equals ...
func (r *role) Equals(comparison Role) bool {
	if r == nil || comparison == nil {
		return false
	}
	return r.GetVal() == comparison.GetVal()
}

// IsSystemAdministrator ... システム管理者ロールか否かを返す
func (r *role) IsSystemAdministrator() bool {
	return r.val == SystemAdministrator
}

// IsOrganizationManager ... 組織管理者ロールか否かを返す
func (r *role) IsOrganizationManager() bool {
	return r.val == OrganizationManager
}

// IsAuthenticationUser ... 認証ユーザーロールか否かを返す
func (r *role) IsAuthenticationUser() bool {
	return r.val == AuthenticationUser
}

// IsGuestUser ... ゲストユーザーロールか否かを返す
func (r *role) IsGuestUser() bool {
	return r.val == GuestUser
}

// IsSystemAdministrator ... システム管理者ロールか否かを返す
func IsSystemAdministrator(val int) bool {
	return val == SystemAdministrator
}

// IsOrganizationManager ... 組織管理者ロールか否かを返す
func IsOrganizationManager(val int) bool {
	return val == OrganizationManager
}

// IsAuthenticationUser ... 認証ユーザーロールか否かを返す
func IsAuthenticationUser(val int) bool {
	return val == AuthenticationUser
}

// IsGuestUser ... ゲストユーザーロールか否かを返す
func IsGuestUser(val int) bool {
	return val == GuestUser
}
