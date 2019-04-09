package factory

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/entity"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
	"time"
)

// NewNoticeFactory ...
func NewNoticeFactory(id string) *NoticeFactory {
	return &NoticeFactory{uid: vo.NewUniqueIDByParam(id)}
}

// NoticeFactory ... 「お知らせ」集約のあらゆる生成方法を担う
type NoticeFactory struct {
	uid vo.UniqueID
}

// CreateNotice ... domain外からの値の受け渡しのために、プリミティブないしgo標準パッケージの型を使う
func (f *NoticeFactory) CreateNotice(title, detail string, noticeSeverity, publishType int, from, to *time.Time) (aggregate.Notice, error.ApplicationError) {
	attribute, err := entity.NewNoticeNoticeAttribute(f.uid, title, detail)
	if err != nil {
		return nil, err
	}
	severity := vo.NewNoticeSeverity(noticeSeverity)
	publishControl := vo.NewPublishControl(vo.NewPublishType(publishType), vo.NewPublishTerm(from, to))

	return aggregate.NewNotice(attribute, severity, publishControl), nil
}
