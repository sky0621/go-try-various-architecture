package driver

import (
	"database/sql"

	cloudsqlproxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/go-sql-driver/mysql"
)

// NewCloudSQLProxyConn ...
func NewCloudSQLProxyConn(address, dbName, user, passwd string) (*sql.DB, error) {
	return cloudsqlproxy.DialCfg(&mysql.Config{
		Addr:      address,
		DBName:    dbName,
		User:      user,
		Passwd:    passwd,
		Net:       "cloudsql",
		ParseTime: true,
		TLSConfig: "",
	})
}
