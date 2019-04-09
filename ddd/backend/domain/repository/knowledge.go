package repository

import (
	"context"
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// KnowledgeCommandRepository ... 「ナレッジ」データへのCRUDのうち、Commandに該当するCUDを担う。
type KnowledgeCommandRepository interface {
	/*
	 * 新規登録
	 */
	// Create ... 引数で渡された「ナレッジ」データ１件を作成する。
	Create(context.Context, aggregate.Knowledge) error.ApplicationError

	// CreateBatch ... 引数で渡された「ナレッジ」データ複数件を作成する。
	CreateBatch(context.Context, []aggregate.Knowledge) error.ApplicationError

	/*
	 * 更新
	 */
	// UpdateByUniqueID ... 引数で渡された「ナレッジ」データに含まれるIDを条件に１件を更新する。
	UpdateByUniqueID(context.Context, aggregate.Knowledge) error.ApplicationError

	// UpdateByCondition ... 引数で渡された「ナレッジ」データ更新条件に合致する複数の「ナレッジ」データを更新する。
	UpdateByCondition(context.Context, aggregate.KnowledgeCommandCondition, aggregate.Knowledge) error.ApplicationError

	// UpdateBatch ... 引数で渡された「ナレッジ」データ複数件を更新する。
	UpdateBatch(context.Context, []aggregate.Knowledge) error.ApplicationError

	/*
	 * 削除
	 */
	// DeleteByUniqueID ... 引数で渡されたユニークIDで特定される「ナレッジ」データ１件を削除する。
	DeleteByUniqueID(context.Context, vo.UniqueID) error.ApplicationError

	// DeleteByCondition ... 引数で渡された「ナレッジ」データ削除条件に合致する複数の「ナレッジ」データを削除する。
	DeleteByCondition(context.Context, aggregate.KnowledgeCommandCondition) error.ApplicationError

	// DeleteBatch ... 引数で渡されたユニークIDで特定される「ナレッジ」データ複数件を削除する。
	DeleteBatch(context.Context, []vo.UniqueID) error.ApplicationError
}

// KnowledgeQueryRepository ... 「ナレッジ」データへのCRUDのうち、Queryに該当するRを担う。
type KnowledgeQueryRepository interface {
	// GetByUniqueID ... 引数で渡されたユニークIDで特定される「ナレッジ」データ１件を取得する。（取得対象が存在しない場合は nil を返却する。）
	GetByUniqueID(context.Context, vo.UniqueID) (aggregate.Knowledge, error.ApplicationError)

	// GetByCondition ... 引数で渡された「ナレッジ」データ取得条件に合致する複数の「ナレッジ」データを返却する。（取得条件が nil の場合は全ての「ナレッジ」データを返却する。）
	GetByCondition(context.Context, aggregate.KnowledgeQueryCondition) ([]aggregate.Knowledge, error.ApplicationError)
}
