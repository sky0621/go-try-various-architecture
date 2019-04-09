package input

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/error"
	"go-ddd/backend/domain/factory"
	"time"
)

// Knowledge ... ユースケースとしての「ナレッジ」情報を定義
type Knowledge struct {
	// ID
	ID string

	// タイトル
	Title string

	// 詳細
	Detail string

	// 特定情報の公開設定
	PublishType int

	// 公開開始日時
	From *time.Time

	// 公開終了日時
	To *time.Time
}

// ConvertToAggregate ...
func (i *Knowledge) ConvertToAggregate() (aggregate.Knowledge, error.ApplicationError) {
	f := factory.NewKnowledgeFactory(i.ID)
	n, err := f.CreateKnowledge(i.Title, i.Detail, i.PublishType, i.From, i.To)
	if err != nil {
		return nil, err
	}
	return n, nil
}
