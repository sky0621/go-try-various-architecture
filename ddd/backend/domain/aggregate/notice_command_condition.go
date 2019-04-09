package aggregate

// NewNoticeCommandCondition ...
func NewNoticeCommandCondition(notices []Notice) NoticeCommandCondition {
	return &noticeCommandCondition{notices: notices}
}

// NoticeCommandCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeCommandCondition interface {

	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type noticeCommandCondition struct {
	notices []Notice
}
