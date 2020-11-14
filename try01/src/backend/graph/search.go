package graph

import (
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type searchParam struct {
	orderKey       string
	orderDirection string
	tableName      string
	baseCondition  string
	rowNumFrom     int64
	rowNumTo       int64
}

// TODO: とりあえず雑に作った。複数テーブルへの対応等、どこまで汎用性を持たせるかは要件次第。
func buildSearchQueryMod(p searchParam) qm.QueryMod {
	if p.baseCondition == "" {
		p.baseCondition = "true"
	}
	q := `
		SELECT row_num, * FROM (
			SELECT ROW_NUMBER() OVER (ORDER BY %s %s) AS row_num, *
			FROM %s
			WHERE %s
		) AS tmp
		WHERE row_num BETWEEN %d AND %d
	`
	sql := fmt.Sprintf(q,
		p.orderKey, p.orderDirection,
		p.tableName,
		p.baseCondition,
		p.rowNumFrom, p.rowNumTo,
	)
	return qm.SQL(sql)
}
