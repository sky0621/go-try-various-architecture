package factory

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/entity"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
	"time"
)

// NewKnowledgeFactory ...
func NewKnowledgeFactory(id string) *KnowledgeFactory {
	return &KnowledgeFactory{uid: vo.NewUniqueIDByParam(id)}
}

// KnowledgeFactory ... 「お知らせ」集約のあらゆる生成方法を担う
type KnowledgeFactory struct {
	uid vo.UniqueID
}

// CreateKnowledge ...
func (f *KnowledgeFactory) CreateKnowledge(title, detail string, publishType int, from, to *time.Time) (aggregate.Knowledge, error.ApplicationError) {
	attribute, err := entity.NewKnowledgeAttribute(f.uid, title, detail)
	if err != nil {
		return nil, err
	}
	publishControl := vo.NewPublishControl(vo.NewPublishType(publishType), vo.NewPublishTerm(from, to))

	return aggregate.NewKnowledge(attribute, publishControl), nil
}
