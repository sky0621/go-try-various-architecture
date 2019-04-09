package usecase

import (
	"context"
	"go-ddd/backend/domain/error"
	"go-ddd/backend/domain/repository"
	"go-ddd/backend/usecase/input"
)

// NewNotice ...
func NewNotice(command repository.NoticeCommandRepository, query repository.NoticeQueryRepository) NoticeUsecase {
	return &noticeUsecase{
		command: command,
		query:   query,
	}
}

// NoticeUsecase ... 「お知らせ」に対するユースケース定義
type NoticeUsecase interface {
	AddNotice(context.Context, *input.Notice) error.ApplicationError
}

type noticeUsecase struct {
	command repository.NoticeCommandRepository
	query   repository.NoticeQueryRepository
}

// AddNotice ...
func (u *noticeUsecase) AddNotice(ctx context.Context, in *input.Notice) error.ApplicationError {
	n, err := in.ConvertToAggregate()
	if err != nil {
		return err
	}
	return u.command.Create(ctx, n)
}
