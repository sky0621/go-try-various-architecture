package aggregate

import (
	"go-ddd/backend/domain/entity"
	vo "go-ddd/backend/domain/valueobject"
)

// NewNotice ...
func NewNotice(attribute entity.NoticeAttribute, severity vo.NoticeSeverity, publishControl vo.PublishControl) Notice {
	return &notice{
		attribute:      attribute,
		severity:       severity,
		publishControl: publishControl,
	}
}

// Notice ... 「お知らせ」集約情報
type Notice interface {
	GetAttribute() entity.NoticeAttribute
	GetNoticeSeverity() vo.NoticeSeverity
	GetPublishControl() vo.PublishControl
}

type notice struct {
	// 「お知らせ」属性
	attribute entity.NoticeAttribute

	// 重要度
	severity vo.NoticeSeverity

	// 公開設定
	publishControl vo.PublishControl
}

// GetAttribute ...
func (a *notice) GetAttribute() entity.NoticeAttribute {
	if a == nil {
		return nil
	}
	return a.attribute
}

// GetNoticeSeverity ...
func (a *notice) GetNoticeSeverity() vo.NoticeSeverity {
	if a == nil {
		return nil
	}
	return a.severity
}

// GetPublishControl ...
func (a *notice) GetPublishControl() vo.PublishControl {
	if a == nil {
		return nil
	}
	return a.publishControl
}
