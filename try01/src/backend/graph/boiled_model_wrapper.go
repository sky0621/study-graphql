package graph

import "github.com/sky0621/study-graphql/try01/src/backend/boiled"

type CustomerWithRowNum struct {
	RowNum          int64 `boil:"row_num"`
	boiled.Customer `boil:",bind"`
}
