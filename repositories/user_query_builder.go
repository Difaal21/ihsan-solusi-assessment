package repositories

import (
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/mysql"
	"github.com/stephenafamo/bob/dialect/mysql/dialect"
	"github.com/stephenafamo/bob/dialect/mysql/sm"
)

func buildUserQuery() bob.BaseQuery[*dialect.SelectQuery] {
	query := mysql.Select(
		sm.Columns(
			"u.id",
			"u.name",
			"u.created_at",
			"ua.email",
			"ua.is_email_verified",
			"ua.email_verified_at",
			"ua.phone_number",
			"ua.is_phone_number_verified",
			"ua.phone_number_verified_at",
			"ua.password",
			"ua.is_mfa_enabled",
			"ua.mfa_secret_key",
			"ua.mfa_recovery_code",
		),
		sm.From("users").As("u"),
		sm.InnerJoin("users_account").As("ua").OnEQ(mysql.Raw("u.id"), mysql.Raw("ua.user_id")),
	)

	return query
}
