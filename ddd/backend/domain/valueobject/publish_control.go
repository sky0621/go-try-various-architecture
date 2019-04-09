package valueobject

// NewPublishControl ...
func NewPublishControl(publishType PublishType, publishTerm PublishTerm) PublishControl {
	return &publishControl{
		publishType: publishType,
		publishTerm: publishTerm,
	}
}

// PublishControl ... 特定情報の公開制御
type PublishControl interface {
	GetPublishType() PublishType
	GetPublishTerm() PublishTerm
}

type publishControl struct {
	// publishType ... 特定情報の公開設定
	publishType PublishType

	// publishTerm ... 公開期間
	publishTerm PublishTerm
}

// GetPublishType ...
func (v *publishControl) GetPublishType() PublishType {
	if v == nil {
		return nil
	}
	return v.publishType
}

// GetPublishTerm ...
func (v *publishControl) GetPublishTerm() PublishTerm {
	if v == nil {
		return nil
	}
	return v.publishTerm
}
