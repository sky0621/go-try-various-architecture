package repository

import (
	"context"
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// OrganizationCommandRepository ... 「組織」データへのCRUDのうち、Commandに該当するCUDを担う。
type OrganizationCommandRepository interface {
	/*
	 * 新規登録
	 */
	// Create ... 引数で渡された「組織」データ１件を作成する。
	Create(context.Context, aggregate.Organization) error.ApplicationError

	// CreateBatch ... 引数で渡された「組織」データ複数件を作成する。
	CreateBatch(context.Context, []aggregate.Organization) error.ApplicationError

	/*
	 * 更新
	 */
	// UpdateByUniqueID ... 引数で渡された「組織」データに含まれるIDを条件に１件を更新する。
	UpdateByUniqueID(context.Context, aggregate.Organization) error.ApplicationError

	// UpdateByCondition ... 引数で渡された「組織」データ更新条件に合致する複数の「組織」データを更新する。
	UpdateByCondition(context.Context, aggregate.OrganizationCommandCondition, aggregate.Organization) error.ApplicationError

	// UpdateBatch ... 引数で渡された「組織」データ複数件を更新する。
	UpdateBatch(context.Context, []aggregate.Organization) error.ApplicationError

	/*
	 * 削除
	 */
	// DeleteByUniqueID ... 引数で渡されたユニークIDで特定される「組織」データ１件を削除する。
	DeleteByUniqueID(context.Context, vo.OrganizationID) error.ApplicationError

	// DeleteByCondition ... 引数で渡された「組織」データ削除条件に合致する複数の「組織」データを削除する。
	DeleteByCondition(context.Context, aggregate.OrganizationCommandCondition) error.ApplicationError

	// DeleteBatch ... 引数で渡されたユニークIDで特定される「組織」データ複数件を削除する。
	DeleteBatch(context.Context, []vo.OrganizationID) error.ApplicationError
}

// OrganizationQueryRepository ... 「組織」データへのCRUDのうち、Queryに該当するRを担う。
type OrganizationQueryRepository interface {
	// GetByUniqueID ... 引数で渡されたユニークIDで特定される「組織」データ１件を取得する。（取得対象が存在しない場合は nil を返却する。）
	GetByUniqueID(context.Context, vo.OrganizationID) (aggregate.Organization, error.ApplicationError)

	// GetByCondition ... 引数で渡された「組織」データ取得条件に合致する複数の「組織」データを返却する。（取得条件が nil の場合は全ての「組織」データを返却する。）
	GetByCondition(context.Context, aggregate.OrganizationQueryCondition) ([]aggregate.Organization, error.ApplicationError)
}
