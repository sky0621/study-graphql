package backend

import (
	"context"

	"github.com/sky0621/study-graphql/backend/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}
func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	panic("not implemented")
}
