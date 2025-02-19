package repositories

import (
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func buildFinancialAccountQuery() bob.BaseQuery[*dialect.SelectQuery] {
	query := psql.Select(
		sm.Columns(
			"fa.id",
			"fa.user_id",
			"fa.balance",
			"fa.bank_account_number",
			"fa.created_at",
		),
		sm.From("financial_accounts").As("fa"),
	)

	return query
}
