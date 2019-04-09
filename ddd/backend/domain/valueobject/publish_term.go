package valueobject

import "time"

// NewPublishTerm ...
func NewPublishTerm(from, to *time.Time) PublishTerm {
	if from == nil || to == nil {
		return nil
	}
	if from.UnixNano() > to.UnixNano() {
		return nil
	}
	return &publishTerm{from: from, to: to}
}

// NewPublishTermFrom ...
func NewPublishTermFrom(from *time.Time) PublishTerm {
	if from == nil {
		return nil
	}
	return &publishTerm{from: from}
}

// NewPublishTermTo ...
func NewPublishTermTo(to *time.Time) PublishTerm {
	if to == nil {
		return nil
	}
	return &publishTerm{to: to}
}

// PublishTerm ... 公開期間
type PublishTerm interface {
	PublishFrom() *time.Time
	PublishTo() *time.Time
}

type publishTerm struct {
	// 公開開始日時（ nil の場合は即時開始）
	from *time.Time

	// 公開終了日時（ nil の場合は終了）
	to *time.Time
}

// PublishFrom ... 公開開始日時を返却
func (v *publishTerm) PublishFrom() *time.Time {
	if v == nil {
		return nil
	}
	return v.from
}

// PublishTo ... 公開終了日時を返却
func (v *publishTerm) PublishTo() *time.Time {
	if v == nil {
		return nil
	}
	return v.to
}
