package graph

import (
	"errors"

	"github.com/sky0621/study-graphql/try01/src/backend/boiled"
)

type CustomerWithRowNum struct {
	RowNum          int64 `boil:"row_num"`
	boiled.Customer `boil:",bind"`
}

func decodeCustomerCursor(cursor string) (int64, error) {
	modelName, key, err := decodeCursor(cursor)
	if err != nil {
		return 0, err
	}
	if modelName != "customer" {
		return 0, errors.New("not customer")
	}
	return key, nil
}
