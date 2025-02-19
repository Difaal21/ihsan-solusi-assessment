package repositories

import (
	"database/sql"
	"difaal21/ihsan-solusi-assessment/database/postgresql"
	"difaal21/ihsan-solusi-assessment/entities"
)

func scanUser(row *sql.Row) (*entities.Users, error) {
	var user = &entities.Users{}

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.NationalityID, &user.CreatedAt, &user.Balance)
	err = postgresql.WrapError(err)
	if err != nil {
		return nil, err
	}

	return user, nil
}
