package repository

import (
	"context"
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NoticeCommandRepository ... 「お知らせ」データへのCRUDのうち、Commandに該当するCUDを担う。
type NoticeCommandRepository interface {
	/*
	 * 新規登録
	 */
	// Create ... 引数で渡された「お知らせ」データ１件を作成する。
	Create(context.Context, aggregate.Notice) error.ApplicationError

	// CreateBatch ... 引数で渡された「お知らせ」データ複数件を作成する。
	CreateBatch(context.Context, []aggregate.Notice) error.ApplicationError

	/*
	 * 更新
	 */
	// UpdateByUniqueID ... 引数で渡された「お知らせ」データに含まれるIDを条件に１件を更新する。
	UpdateByUniqueID(context.Context, aggregate.Notice) error.ApplicationError

	// UpdateByCondition ... 引数で渡された「お知らせ」データ更新条件に合致する複数の「お知らせ」データを更新する。
	UpdateByCondition(context.Context, aggregate.NoticeCommandCondition, aggregate.Notice) error.ApplicationError

	// UpdateBatch ... 引数で渡された「お知らせ」データ複数件を更新する。
	UpdateBatch(context.Context, []aggregate.Notice) error.ApplicationError

	/*
	 * 削除
	 */
	// DeleteByUniqueID ... 引数で渡されたユニークIDで特定される「お知らせ」データ１件を削除する。
	DeleteByUniqueID(context.Context, vo.UniqueID) error.ApplicationError

	// DeleteByCondition ... 引数で渡された「お知らせ」データ削除条件に合致する複数の「お知らせ」データを削除する。
	DeleteByCondition(context.Context, aggregate.NoticeCommandCondition) error.ApplicationError

	// DeleteBatch ... 引数で渡されたユニークIDで特定される「お知らせ」データ複数件を削除する。
	DeleteBatch(context.Context, []vo.UniqueID) error.ApplicationError
}

// NoticeQueryRepository ... 「お知らせ」データへのCRUDのうち、Queryに該当するRを担う。
type NoticeQueryRepository interface {
	// GetByUniqueID ... 引数で渡されたユニークIDで特定される「お知らせ」データ１件を取得する。（取得対象が存在しない場合は nil を返却する。）
	GetByUniqueID(context.Context, vo.UniqueID) (aggregate.Notice, error.ApplicationError)

	// GetByCondition ... 引数で渡された「お知らせ」データ取得条件に合致する複数の「お知らせ」データを返却する。（取得条件が nil の場合は全ての「お知らせ」データを返却する。）
	GetByCondition(context.Context, aggregate.NoticeQueryCondition) ([]aggregate.Notice, error.ApplicationError)
}
