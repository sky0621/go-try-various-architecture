package valueobject

// NewHumanName ...
func NewHumanName(lastName, firstName string) HumanName {
	return &humanName{lastName: lastName, firstName: firstName}
}

// HumanName ... 名前
type HumanName interface {
	GetLastName() string
	GetFirstName() string
	Equals(comparison HumanName) bool
}

type humanName struct {
	// 姓
	lastName string

	// 名
	firstName string
}

// GetLastName ...
func (v *humanName) GetLastName() string {
	if v == nil {
		return ""
	}
	return v.lastName
}

// GetFirstName ...
func (v *humanName) GetFirstName() string {
	if v == nil {
		return ""
	}
	return v.firstName
}

// Equals ...
func (v *humanName) Equals(comparison HumanName) bool {
	if v == nil || comparison == nil {
		return false
	}
	return (v.GetLastName() == comparison.GetLastName()) && (v.GetFirstName() == comparison.GetFirstName())
}
