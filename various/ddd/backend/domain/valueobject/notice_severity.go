package valueobject

const (
	// NoticeSeverityInfo ... 通常レベルのお知らせ
	NoticeSeverityInfo = iota + 1

	// NoticeSeverityImportant ... 重要レベルのお知らせ
	NoticeSeverityImportant
)

// NewNoticeSeverity ...
func NewNoticeSeverity(val int) NoticeSeverity {
	// 事前条件
	assertion := func(v int) bool {
		if v == NoticeSeverityInfo {
			return true
		}
		if v == NoticeSeverityImportant {
			return true
		}
		return false
	}
	if !assertion(val) {
		return nil
	}

	return &noticeSeverity{val: val}
}

// NewNoticeSeverityInfo ... 通常レベルのお知らせを生成する
func NewNoticeSeverityInfo() NoticeSeverity {
	return &noticeSeverity{val: NoticeSeverityInfo}
}

// NewNoticeSeverityImportant ... 重要レベルのお知らせを生成する
func NewNoticeSeverityImportant() NoticeSeverity {
	return &noticeSeverity{val: NoticeSeverityImportant}
}

// NoticeSeverity ... お知らせの重要度
type NoticeSeverity interface {
	GetVal() int
	Equals(comparison NoticeSeverity) bool
	IsNoticeSeverityInfo() bool
	IsNoticeSeverityImportant() bool
}

type noticeSeverity struct {
	val int
}

// GetVal ...
func (v *noticeSeverity) GetVal() int {
	return v.val
}

// Equals ...
func (v *noticeSeverity) Equals(comparison NoticeSeverity) bool {
	if v == nil || comparison == nil {
		return false
	}
	return v.GetVal() == comparison.GetVal()
}

// IsNoticeSeverityInfo ... 通常レベルのお知らせか否かを返す
func (v *noticeSeverity) IsNoticeSeverityInfo() bool {
	if v == nil {
		return false
	}
	return v.val == NoticeSeverityInfo
}

// IsNoticeSeverityImportant ... 重要レベルのお知らせか否かを返す
func (v *noticeSeverity) IsNoticeSeverityImportant() bool {
	if v == nil {
		return false
	}
	return v.val == NoticeSeverityImportant
}

// IsNoticeSeverityInfo ... 通常レベルのお知らせか否かを返す
func IsNoticeSeverityInfo(val int) bool {
	return val == NoticeSeverityInfo
}

// IsNoticeSeverityImportant ... 重要レベルのお知らせか否かを返す
func IsNoticeSeverityImportant(val int) bool {
	return val == NoticeSeverityImportant
}
