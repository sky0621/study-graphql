package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/sky0621/study-graphql/try01/src/backend/boiled"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
)

func (r *customerResolver) Todos(ctx context.Context, obj *model.Customer) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CustomerConnection(ctx context.Context, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder, filterWord *model.TextFilterCondition) (*model.CustomerConnection, error) {
	var params searchParam

	/*
	 * 検索文字列フィルタ設定
	 */
	if filterWord != nil {
		switch filterWord.MatchingPattern {
		case model.MatchingPatternPartialMatch:

		case model.MatchingPatternExactMatch:

		default:
			return nil, errors.New("no match")
		}
		params.baseCondition = fmt.Sprintf("%s %s", boiled.CustomerColumns.Name)
	}

	params.tableName = boiled.TableNames.Customer
	params.baseCondition = "true"
	params.orderKey = "age"
	params.orderDirection = orderDirectionDesc
	params.compareSymbol = compareSymbolLt
	params.decodedCursor = 5
	params.limit = 5

	var records []*CustomerWithRowNum
	if err := boiled.Customers(buildSearchQueryMod(params)).Bind(ctx, r.DB, &records); err != nil {
		return nil, err
	}

	result := &model.CustomerConnection{}
	var edges []*model.CustomerEdge
	for _, record := range records {
		edges = append(edges, &model.CustomerEdge{
			Node: &model.Customer{
				ID:   strconv.Itoa(int(record.ID)),
				Name: record.Name,
				Age:  record.Age,
			},
		})
	}
	result.Edges = edges
	return result, nil
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

type customerResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type CustomerWithRowNum struct {
	RowNum          int64 `boil:"row_num"`
	boiled.Customer `boil:",bind"`
}
