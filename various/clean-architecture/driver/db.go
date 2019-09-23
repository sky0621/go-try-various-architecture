package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	DataSourceFormat = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

// RDBという他のレイヤーから直接呼ぶのはNG（ドライバーという詳細に依存してしまう）なのでDI戦略が必要
func NewDBConnection(user, password, instance, db string) (*gorm.DB, error) {
	dataSource := fmt.Sprintf(DataSourceFormat, user, password, instance, db)
	return gorm.Open("mysql", dataSource)
}
