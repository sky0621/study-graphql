package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
	"github.com/sky0621/study-graphql/try01/src/backend/sqlboiler"
)

func (r *queryResolver) TodoConnection(ctx context.Context,
	filterWord *model.TextFilterCondition,
	pageCondition *model.PageCondition,
	edgeOrder *model.EdgeOrder) (*model.TodoConnection, error) {

	if filterWord == nil && pageCondition == nil && edgeOrder == nil {
		todos, err := sqlboiler.Todos().All(ctx, r.DB)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		var edges []*model.TodoEdge
		for _, todo := range todos {
			edges = append(edges, &model.TodoEdge{
				Node: &model.Todo{
					ID:     todo.ID,
					Task:   todo.Task,
					UserID: todo.UserID,
				},
				Cursor: CreateCursor("todo", todo.ID),
			})
		}

		todoConn := &model.TodoConnection{
			Edges:      edges,
			TotalCount: len(edges),
		}
		return todoConn, nil
	}
	return nil, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj == nil {
		return nil, nil
	}
	return For(ctx).UsersByIDs.Load(obj.UserID)
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
