package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
)

func (r *customerResolver) Todos(ctx context.Context, obj *model.Customer) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CustomerConnection(ctx context.Context, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder, filterWord *model.TextFilterCondition) (*model.CustomerConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

type customerResolver struct{ *Resolver }
