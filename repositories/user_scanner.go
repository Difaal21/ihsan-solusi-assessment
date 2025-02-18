package repositories

import (
	"database/sql"
	"difaal21/ihsan-solusi-assessment/database/postgresql"
	"difaal21/ihsan-solusi-assessment/entities"
)

func scanUser(row *sql.Row) (user entities.Users, err error) {
	var phoneNumber sql.NullString
	var phoneNumberVerifiedAt sql.NullTime
	var mfaSecretKey sql.NullString
	var mfaRecoveryCode sql.NullString

	err = row.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.Email, &user.IsEmailVerified, &user.EmailVerifiedAt, &phoneNumber, &user.IsPhoneNumberVerified, &phoneNumberVerifiedAt, &user.Password, &user.IsMFAEnabled, &mfaSecretKey, &mfaRecoveryCode)

	err = postgresql.WrapError(err)
	if err != nil {
		return
	}

	if phoneNumber.Valid {
		user.PhoneNumber = phoneNumber.String
	}

	if mfaSecretKey.Valid {
		user.MFASecretKey = mfaSecretKey.String
	}

	if mfaRecoveryCode.Valid {
		user.MFARecoveryCode = mfaRecoveryCode.String
	}

	return
}
