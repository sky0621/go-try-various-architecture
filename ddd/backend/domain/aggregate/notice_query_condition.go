package aggregate

// NewNoticeQueryCondition ...
func NewNoticeQueryCondition(notices []Notice) NoticeQueryCondition {
	return &noticeQueryCondition{notices: notices}
}

// NoticeQueryCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeQueryCondition interface {

	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type noticeQueryCondition struct {
	notices []Notice
}
