package entity

import (
	"errors"
	"go-ddd/backend/domain/error"
	vo "go-ddd/backend/domain/valueobject"
)

// NewKnowledgeAttribute ...
func NewKnowledgeAttribute(uid vo.UniqueID, title, detail string) (KnowledgeAttribute, error.ApplicationError) {
	if uid.GetVal() == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("id is required"))
	}
	if title == "" {
		return nil, error.CreateValidationError(error.Required, errors.New("title is required"))
	}

	return &knowledgeAttribute{
		id:     uid,
		title:  title,
		detail: detail,
	}, nil
}

// Knowledge ... 「ナレッジ」データ定義
type KnowledgeAttribute interface {
	GetID() vo.UniqueID
	GetTitle() string
	GetDetail() string
}

type knowledgeAttribute struct {
	// ユニークに特定するID
	id vo.UniqueID

	// 概要を示すタイトル
	title string

	// 詳細
	detail string
}

// GetID ...
func (e *knowledgeAttribute) GetID() vo.UniqueID {
	if e == nil {
		return nil
	}
	return e.id
}

// GetTitle ...
func (e *knowledgeAttribute) GetTitle() string {
	if e == nil {
		return ""
	}
	return e.title
}

// GetDetail ...
func (e *knowledgeAttribute) GetDetail() string {
	if e == nil {
		return ""
	}
	return e.detail
}
