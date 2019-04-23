package aggregate

// NewKnowledgeCommandCondition ...
func NewKnowledgeCommandCondition(knowledges []Knowledge) KnowledgeCommandCondition {
	return &knowledgeCommandCondition{knowledges: knowledges}
}

// KnowledgeCommandCondition ... 条件に該当する「ナレッジ」データを決定するために利用
type KnowledgeCommandCondition interface {

	// FIXME 複数の Knowledge を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

type knowledgeCommandCondition struct {
	knowledges []Knowledge
}
