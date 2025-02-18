package postgresql

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type PostgreSQL interface {
	Connect(maxOpenConns int, maxIdleConns int) (db *sql.DB, err error)
	Close(db *sql.DB) (err error)
}

type PostgreSQLImpl struct {
	DriverName, DataSourceName string
	logger                     *logrus.Logger
}

func NewPostgreSQL(driver string, dataSource string, logger *logrus.Logger) PostgreSQL {
	return &PostgreSQLImpl{DriverName: driver, DataSourceName: dataSource, logger: logger}
}
