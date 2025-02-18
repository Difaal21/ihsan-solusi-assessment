package postgresql

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func (p *PostgreSQLImpl) Connect(maxOpenConns int, maxIdleConns int) (db *sql.DB, err error) {
	db, err = sql.Open(p.DriverName, p.DataSourceName)
	if err != nil {
		return
	}

	db.SetMaxIdleConns(maxOpenConns)
	db.SetMaxOpenConns(maxIdleConns)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	if err = db.Ping(); err != nil {
		return
	}

	p.logger.Info("PostgreSQL connected")
	return
}

func (p *PostgreSQLImpl) Close(db *sql.DB) (err error) {
	if err = db.Close(); err != nil {
		p.logger.Info("Failed to close MySQL connection: ", err)
		return
	}
	p.logger.Info("MySQL connection closed")
	return
}
