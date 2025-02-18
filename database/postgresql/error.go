package postgresql

import (
	"database/sql"
	"difaal21/ihsan-solusi-assessment/exceptions"
)

func WrapError(err error) error {
	if err == sql.ErrNoRows {
		return exceptions.ErrNotFound
	}
	return err
}
