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
	err := database.NewTodoDao(r.DB).InsertOne(ctx, &database.Todo{
		ID:   id,
		Text: input.Text,
		Done: false,
		User: database.User{
			ID: input.UserID,
		},
	})
	if err != nil {
		return "", err
	}
	return id, nil
}
func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (string, error) {
	log.Printf("[mutationResolver.CreateUser] input: %#v", input)
	id := util.CreateUniqueID()
	err := database.NewUserDao(r.DB).InsertOne(ctx, &database.User{
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
	todos, err := database.NewTodoDao(r.DB).FindAll(ctx)
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
	todo, err := database.NewTodoDao(r.DB).FindOne(ctx, id)
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
func (r *queryResolver) TodoConnection(ctx context.Context,
	filterWord *models.TextFilterCondition,
	pageCondition *models.PageCondition,
	edgeOrder *models.EdgeOrder) (*models.TodoConnection, error) {

	log.Println("[queryResolver.TodoConnection]")

	dao := database.NewTodoDao(r.DB)

	/*
	 * 検索条件に合致する全件数を取得
	 */
	totalCount, err := dao.CountByTextFilter(ctx, filterWord)
	if err != nil {
		return nil, err
	}
	if totalCount == 0 {
		return models.EmptyTodoConnection(), nil
	}

	// 検索結果全件数と１ページあたりの表示件数から、今回の検索による総ページ数を算出
	totalPage := pageCondition.TotalPage(totalCount)

	// ページ情報を計算・収集しておく
	pageInfo := &models.PageInfo{
		HasNextPage:     (totalPage - int64(pageCondition.MoveToPageNo())) >= 1, // 遷移後も、まだ先のページがあるか
		HasPreviousPage: pageCondition.MoveToPageNo() > 1,                       // 遷移後も、まだ前のページがあるか
	}

	/*
	 * 検索条件、ページング条件、ソート条件に合致する結果を取得
	 */
	todos, err := dao.FindByCondition(ctx, filterWord, pageCondition, getOrder(edgeOrder))
	if err != nil {
		return nil, err
	}
	if todos == nil || len(todos) == 0 {
		return models.EmptyTodoConnection(), nil
	}

	var edges []*models.TodoEdge
	for idx, todo := range todos {
		// 当該レコードをユニークに特定するためのカーソルを計算
		cursor := util.CreateCursor("todo", todo.ID)

		// 検索結果をEdge形式に変換（カーソルの値も格納）
		edges = append(edges, &models.TodoEdge{
			Cursor: cursor,
			Node: &models.Todo{
				ID:        todo.ID,
				Text:      todo.Text,
				Done:      todo.Done,
				CreatedAt: todo.CreatedAt.Unix(),
				User: &models.User{
					ID:   todo.User.ID,
					Name: todo.User.Name,
				},
			},
		})

		if idx == 0 {
			// 今回表示するページの１件目のレコードを特定するカーソルをセット
			// （「前ページ」遷移時に「このカーソルよりも前のレコード」という検索条件に用いる）
			pageInfo.StartCursor = cursor
		}
		if idx == len(todos)-1 {
			// 今回表示するページの最後のレコードを特定するカーソルをセット
			// （「次ページ」遷移時に「このカーソルよりも後のレコード」という検索条件に用いる）
			pageInfo.EndCursor = cursor
		}
	}

	return &models.TodoConnection{
		PageInfo:   pageInfo,
		Edges:      edges,
		TotalCount: totalCount,
	}, nil
}

func getOrder(edgeOrder *models.EdgeOrder) *models.EdgeOrder {
	// 並べ替えは未指定時にデフォルト値を与えておかないとページングとの組み合わせが成り立たないので、ここでセット
	var order *models.EdgeOrder
	if edgeOrder == nil {
		createdAt := models.TodoOrderKeyCreatedAt
		order = &models.EdgeOrder{
			Key:       &models.OrderKey{TodoOrderKey: &createdAt},
			Direction: models.OrderDirectionDesc,
		}
	} else {
		order = edgeOrder
	}
	return order
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	log.Println("[queryResolver.Users]")
	users, err := database.NewUserDao(r.DB).FindAll(ctx)
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
	user, err := database.NewUserDao(r.DB).FindOne(ctx, id)
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
	todos, err := database.NewTodoDao(r.DB).FindByUserID(ctx, obj.ID)
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
