package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sky0621/study-graphql/try01/src/backend/boiled"

	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
)

func (r *queryResolver) CustomerConnection(ctx context.Context, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder, filterWord *model.TextFilterCondition) (*model.CustomerConnection, error) {
	records, err := boiled.Customers().All(ctx, r.DB)
	if err != nil {
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

func (r *customerResolver) Todos(ctx context.Context, obj *model.Customer) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

type customerResolver struct{ *Resolver }
