package backend

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	return &Todo{
		ID:   "todo001",
		Text: "部屋の掃除",
		Done: false,
		User: &User{
			ID:   "user001",
			Name: "たろー",
		},
	},nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return []*Todo{
		&Todo{
			ID:   "todo001",
			Text: "部屋の掃除",
			Done: false,
			User: &User{
				ID:   "user001",
				Name: "たろー",
			},
		},
		&Todo{
			ID:   "todo002",
			Text: "買い物",
			Done: true,
			User: &User{
				ID:   "user001",
				Name: "たろー",
			},
		},
	},nil
}
