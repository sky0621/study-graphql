package backend

import (
	"context"
	"log"

	"github.com/sky0621/study-graphql/backend/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

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

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	log.Printf("[mutationResolver.CreateTodo] input: %#v", input)
	return &models.Todo{
		ID:   "todo001",
		Text: "部屋の掃除",
		Done: false,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*models.User, error) {
	log.Printf("[mutationResolver.CreateUser] input: %#v", input)
	return &models.User{
		ID:   "user001",
		Name: "たろー",
	}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	log.Println("[queryResolver.Todos]")
	return []*models.Todo{
		{
			ID:   "todo001",
			Text: "部屋の掃除",
			Done: false,
		},
		{
			ID:   "todo002",
			Text: "買い物",
			Done: true,
		},
	}, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	log.Printf("[queryResolver.Todo] id: %s", id)
	return &models.Todo{
		ID:   "todo001",
		Text: "ヘヤノソウジ",
		Done: false,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	log.Println("[queryResolver.Users]")
	return []*models.User{
		{
			ID:   "user001",
			Name: "タロー",
		},
		{
			ID:   "user002",
			Name: "ジロー",
		},
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	log.Printf("[queryResolver.User] id: %s", id)
	return &models.User{
		ID:   "user001",
		Name: "タロー",
	}, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	log.Printf("[todoResolver.User] id: %#v", obj)
	return &models.User{
		ID:   "user101",
		Name: "タロー101",
	}, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	log.Println("[userResolver.Todos]")
	return []*models.Todo{
		{
			ID:   "todo101",
			Text: "部屋の掃除101",
			Done: false,
		},
		{
			ID:   "todo102",
			Text: "買い物102",
			Done: true,
		},
	}, nil
}
