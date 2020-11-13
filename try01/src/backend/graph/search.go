package graph

import (
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type searchParam struct {
	orderKey       string
	orderDirection orderDirection
	tableName      string
	baseCondition  string
	compareSymbol  compareSymbol
	decodedCursor  int64
	limit          int64
}

type orderDirection string

const (
	orderDirectionAsc  orderDirection = "ASC"
	orderDirectionDesc orderDirection = "DESC"
)

type compareSymbol string

const (
	compareSymbolGt compareSymbol = ">"
	compareSymbolGe compareSymbol = ">="
	compareSymbolLt compareSymbol = "<"
	compareSymbolLe compareSymbol = "<="
	compareSymbolEq compareSymbol = "="
)

// TODO: とりあえず雑に作ったが、少なくとも JOIN には対応したい。どこまで汎用性を持たせるかは要件次第。
func buildSearchQueryMod(p searchParam) qm.QueryMod {
	q := `
		SELECT row_num, * FROM (
			SELECT ROW_NUMBER() OVER (ORDER BY %s %s) AS row_num, *
			FROM %s
		) AS tmp
		WHERE %s
		AND row_num %s %d
		LIMIT %d
	`
	sql := fmt.Sprintf(q,
		p.orderKey, p.orderDirection,
		p.tableName,
		p.baseCondition, p.compareSymbol, p.decodedCursor,
		p.limit,
	)
	return qm.SQL(sql)
}
