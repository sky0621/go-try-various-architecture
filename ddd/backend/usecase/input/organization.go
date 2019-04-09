package input

import (
	"go-ddd/backend/domain/aggregate"
	"go-ddd/backend/domain/error"
	"go-ddd/backend/domain/factory"
)

// Organization ... ユースケースとしての「組織」情報を定義
type Organization struct {
	// ID
	ID string

	// 論理名
	LogicalName string

	// 論理名（カナ）
	LogicalKanaName string

	// 物理名
	PhysicalName string

	// 親組織のID
	ParentID string
}

// ConvertToAggregate ...
func (i *Organization) ConvertToAggregate() (aggregate.Organization, error.ApplicationError) {
	f := factory.NewOrganizationFactory(i.ID)
	n, err := f.CreateOrganization(i.LogicalName, i.LogicalKanaName, i.PhysicalName, i.ParentID)
	if err != nil {
		return nil, err
	}
	return n, nil
}
