package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
)

func (r *queryResolver) TodoConnection(ctx context.Context, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder, filterWord *model.TextFilterCondition) (*model.TodoConnection, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) Customer(ctx context.Context, obj *model.Todo) (*model.Customer, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
