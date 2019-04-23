package input

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/error"
	"go-ddd/backend/domain/factory"
	"time"
)

// Notice ... ユースケースとしての「お知らせ」情報を定義
type Notice struct {
	// ID
	ID string

	// タイトル
	Title string

	// 詳細
	Detail string

	// お知らせの重要度
	NoticeSeverity int

	// 特定情報の公開設定
	PublishType int

	// 公開開始日時
	From *time.Time

	// 公開終了日時
	To *time.Time
}

// ConvertToAggregate ...
func (i *Notice) ConvertToAggregate() (aggregate.Notice, error.ApplicationError) {
	f := factory.NewNoticeFactory(i.ID)
	n, err := f.CreateNotice(i.Title, i.Detail, i.NoticeSeverity, i.PublishType, i.From, i.To)
	if err != nil {
		return nil, err
	}
	return n, nil
}
