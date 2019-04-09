package aggregate

// NewKnowledgeQueryCondition ...
func NewKnowledgeQueryCondition(knowledges []Knowledge) KnowledgeQueryCondition {
	return &knowledgeQueryCondition{knowledges: knowledges}
}

// KnowledgeQueryCondition ... 条件に該当する「ナレッジ」データを決定するために利用
type KnowledgeQueryCondition interface {

	// FIXME 複数の Knowledge を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type knowledgeQueryCondition struct {
	knowledges []Knowledge
}
