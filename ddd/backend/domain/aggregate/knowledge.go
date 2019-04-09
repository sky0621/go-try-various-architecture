package aggregate

import (
	"go-ddd/backend/domain/entity"
	vo "go-ddd/backend/domain/valueobject"
)

// NewKnowledge ...
func NewKnowledge(attribute entity.KnowledgeAttribute, publishControl vo.PublishControl) Knowledge {
	return &knowledge{
		attribute:      attribute,
		publishControl: publishControl,
	}
}

// Knowledge ... 「ナレッジ」集約情報
type Knowledge interface {
	GetAttribute() entity.KnowledgeAttribute
	GetPublishControl() vo.PublishControl
}

type knowledge struct {
	// 「ナレッジ」属性
	attribute entity.KnowledgeAttribute

	// 公開設定
	publishControl vo.PublishControl
}

// GetAttribute ...
func (a *knowledge) GetAttribute() entity.KnowledgeAttribute {
	if a == nil {
		return nil
	}
	return a.attribute
}

// GetPublishControl ...
func (a *knowledge) GetPublishControl() vo.PublishControl {
	if a == nil {
		return nil
	}
	return a.publishControl
}
