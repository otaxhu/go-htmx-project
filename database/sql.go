package database

import (
	"database/sql"
	"fmt"

	"github.com/otaxhu/go-htmx-project/settings"
)

func GetSqlConnection(dbSettings settings.Database) (*sql.DB, error) {
	return sql.Open(dbSettings.Driver, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		dbSettings.User,
		dbSettings.Password,
		dbSettings.Host,
		dbSettings.Port,
		dbSettings.Name,
	))
}
