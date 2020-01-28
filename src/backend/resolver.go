package backend

import (
	"context"
	"errors"
	"log"

	"github.com/sky0621/study-graphql/src/backend/database"
	"github.com/sky0621/study-graphql/src/backend/util"

	"github.com/jinzhu/gorm"
	"github.com/sky0621/study-graphql/src/backend/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB *gorm.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Noop(ctx context.Context, input *NoopInput) (*NoopPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (string, error) {
	log.Printf("[mutationResolver.CreateTodo] input: %#v", input)
	id := util.CreateUniqueID()
	err := database.NewTodoDao(r.DB).InsertOne(&database.Todo{
		ID:     id,
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}
func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (string, error) {
	log.Printf("[mutationResolver.CreateUser] input: %#v", input)
	id := util.CreateUniqueID()
	err := database.NewUserDao(r.DB).InsertOne(&database.User{
		ID:   id,
		Name: input.Name,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Node(ctx context.Context, id string) (Node, error) {
	panic("not implemented")
}
func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	log.Println("[queryResolver.Todos]")
	todos, err := database.NewTodoDao(r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*models.Todo
	for _, todo := range todos {
		results = append(results, &models.Todo{
			ID:        todo.ID,
			Text:      todo.Text,
			Done:      todo.Done,
			CreatedAt: todo.CreatedAt.Unix(),
		})
	}
	return results, nil
}
func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	log.Printf("[queryResolver.Todo] id: %s", id)
	todo, err := database.NewTodoDao(r.DB).FindOne(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("not found")
	}
	return &models.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}
func (r *queryResolver) TodoConnection(ctx context.Context, filterWord *models.TextFilterCondition, pageCondition *models.PageCondition, edgeOrder *models.EdgeOrder) (*models.TodoConnection, error) {
	log.Println("[queryResolver.TodoConnection]")
	dao := database.NewTodoDao(r.DB)
	/*
	 * 検索条件に合致する全件数を取得
	 */
	totalCount, err := dao.CountByTextFilter(filterWord)
	if err != nil {
		return nil, err
	}

	edges := []*models.TodoEdge{}

	pageInfo := &models.PageInfo{}

	return &models.TodoConnection{
		PageInfo:   pageInfo,
		Edges:      edges,
		TotalCount: totalCount,
	}, nil
}
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	log.Println("[queryResolver.Users]")
	users, err := database.NewUserDao(r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*models.User
	for _, user := range users {
		results = append(results, &models.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return results, nil
}
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	log.Printf("[queryResolver.User] id: %s", id)
	user, err := database.NewUserDao(r.DB).FindOne(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("not found")
	}
	return &models.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	log.Println("[userResolver.Todos]")
	todos, err := database.NewTodoDao(r.DB).FindByUserID(obj.ID)
	if err != nil {
		return nil, err
	}
	var results []*models.Todo
	for _, todo := range todos {
		results = append(results, &models.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
}
