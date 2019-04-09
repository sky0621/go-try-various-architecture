package valueobject

const (
	// GeneralPublic ... 一般公開
	GeneralPublic = iota + 1

	// InHousePublish ... 組織内公開
	InHousePublish

	// Private ... 非公開
	Private
)

// NewPublishType ...
func NewPublishType(val int) PublishType {
	// 事前条件
	assertion := func(v int) bool {
		if v == GeneralPublic {
			return true
		}
		if v == InHousePublish {
			return true
		}
		if v == Private {
			return true
		}
		return false
	}
	if !assertion(val) {
		return nil
	}

	return &publishType{val: val}
}

// NewNoticeSeverityInfo ... 一般公開設定を生成する
func NewGeneralPublic() PublishType {
	return &publishType{val: GeneralPublic}
}

// NewInHousePublish ... 組織内公開設定を生成する
func NewInHousePublish() PublishType {
	return &publishType{val: InHousePublish}
}

// NewPrivate ... 非公開設定を生成する
func NewPrivate() PublishType {
	return &publishType{val: Private}
}

// PublishType ... 特定情報の公開設定
type PublishType interface {
	GetVal() int
	Equals(comparison PublishType) bool
	IsGeneralPublic() bool
	IsInHousePublish() bool
	IsPrivate() bool
}

type publishType struct {
	val int
}

// GetVal ...
func (pt *publishType) GetVal() int {
	return pt.val
}

// Equals ...
func (pt *publishType) Equals(comparison PublishType) bool {
	if pt == nil || comparison == nil {
		return false
	}
	return pt.GetVal() == comparison.GetVal()
}

// IsGeneralPublic ... 一般公開設定か否かを返す
func (pt *publishType) IsGeneralPublic() bool {
	return pt.val == GeneralPublic
}

// IsInHousePublish ... 組織内公開設定か否かを返す
func (pt *publishType) IsInHousePublish() bool {
	return pt.val == InHousePublish
}

// IsPrivate ... 非公開設定か否かを返す
func (pt *publishType) IsPrivate() bool {
	return pt.val == Private
}

// IsGeneralPublic ... 一般公開設定か否かを返す
func IsGeneralPublic(val int) bool {
	return val == GeneralPublic
}

// IsInHousePublish ... 組織内公開設定か否かを返す
func IsInHousePublish(val int) bool {
	return val == InHousePublish
}

// IsPrivate ... 非公開設定か否かを返す
func IsPrivate(val int) bool {
	return val == Private
}
