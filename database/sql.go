package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/otaxhu/go-htmx-project/config"
)

func GetSqlConnection(dbCfg config.Database) (*sql.DB, error) {
	return sql.Open(dbCfg.Driver, dbCfg.Url)
}
