package repositories

import (
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func buildUserQuery() bob.BaseQuery[*dialect.SelectQuery] {
	query := psql.Select(
		sm.Columns(
			"u.id",
			"u.name",
			"u.phone_number",
			"u.nationality_id",
			"u.created_at",
		),
		sm.From("users").As("u"),
	)

	return query
}
