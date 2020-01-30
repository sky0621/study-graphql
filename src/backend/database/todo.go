package database

import (
	"context"
	"time"

	"github.com/sky0621/study-graphql/src/backend/models"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID        string    `gorm:"column:id;primary_key"`
	Text      string    `gorm:"column:text"`
	Done      bool      `gorm:"column:done"`
	CreatedAt time.Time `gorm:"column:created_at"`
	// MEMO: 入れ子構造に対してSELECT結果をよしなにマッピングしてくれるかと思ったけどダメだった。
	//User      User      `gorm:"column:user"`
	UserID   string `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
}

func (t *Todo) TableName() string {
	return "todo"
}

// Tableを実装するマーカーインタフェース
func (t *Todo) IsTable() bool {
	return true
}

type TodoDao interface {
	InsertOne(ctx context.Context, u *Todo) error
	FindAll(ctx context.Context) ([]*Todo, error)
	FindByUserID(ctx context.Context, userID string) ([]*Todo, error)
	FindOne(ctx context.Context, id string) (*Todo, error)
	CountByTextFilter(ctx context.Context, filterWord *models.TextFilterCondition) (int64, error)
	FindByCondition(ctx context.Context, filterWord *models.TextFilterCondition, pageCondition *models.PageCondition, edgeOrder *models.EdgeOrder) ([]*Todo, error)
}

type todoDao struct {
	db *gorm.DB
}

func NewTodoDao(db *gorm.DB) TodoDao {
	return &todoDao{db: db}
}

func (d *todoDao) InsertOne(ctx context.Context, u *Todo) error {
	res := d.db.Create(u)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
func (d *todoDao) FindAll(ctx context.Context) ([]*Todo, error) {
	var todos []*Todo
	res := d.db.Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *todoDao) FindOne(ctx context.Context, id string) (*Todo, error) {
	var todos []*Todo
	res := d.db.Where("id = ?", id).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(todos) < 1 {
		return nil, nil
	}
	return todos[0], nil
}

func (d *todoDao) FindByUserID(ctx context.Context, userID string) ([]*Todo, error) {
	var todos []*Todo
	res := d.db.Where("user_id = ?", userID).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *todoDao) CountByTextFilter(ctx context.Context, filterWord *models.TextFilterCondition) (int64, error) {
	// 絞り込み無しのパターン
	if filterWord == nil || filterWord.FilterWord == "" {
		var cnt int64
		if err := d.db.Model(&Todo{}).Count(&cnt).Error; err != nil {
			return 0, err
		}
		return cnt, nil
	}

	// デフォルトは部分一致
	matchStr := "%" + filterWord.FilterWord + "%"
	if filterWord.MatchingPattern != nil && *filterWord.MatchingPattern == models.MatchingPatternExactMatch {
		matchStr = filterWord.FilterWord
	}

	todo := TableName(&Todo{})
	user := TableName(&User{})

	var cnt int64

	// MEMO: ある程度複雑になったら頑張らずに db.Row() で生SQLを書く方が保守性は高いかもしれない。（メソッド使っても生SQL部分は存在するし）
	res := d.db.
		Table(todo).
		Joins(InnerJoin(user) + On("%s.id = %s.user_id", user, todo)).
		Where(Col(todo, "text").Like(matchStr)).
		Or(Col(user, "name").Like(matchStr)).
		Count(&cnt)
	if res.Error != nil {
		return 0, res.Error
	}

	return cnt, nil
}

func (d *todoDao) FindByCondition(ctx context.Context, filterWord *models.TextFilterCondition, pageCondition *models.PageCondition, edgeOrder *models.EdgeOrder) ([]*Todo, error) {
	/*
	 * 文字列フィルタ条件の有無、ページング条件の有無、並べ替え条件の有無の組み合わせによってSQLが変わる。
	 */
	// 組み合わせパターン別にSQL実行
	/*
	 * 文字列フィルタ無し
	 * ページング無し
	 * 並べ替え無し
	 */
	if filterWord.NoFilter() && pageCondition.NoPaging() && edgeOrder.NoSort() {
		var results []*Todo
		if err := d.db.Model(&Todo{}).Find(&results).Error; err != nil {
			return nil, err
		}
		return results, nil
	}

	/*
	 * 文字列フィルタ無し
	 * ページング有り
	 * 並べ替え無し
	 */
	if filterWord.NoFilter() && pageCondition.ExistsPaging() && edgeOrder.NoSort() {
		// FIXME:
	}

	// --------------------------------------
	// 上記のパターン別実装が終われば、以下は不要。
	// --------------------------------------

	var results []*Todo
	// 絞り込み無しのパターン
	if filterWord == nil || filterWord.FilterWord == "" {
		if err := d.db.Model(&Todo{}).Find(&results).Error; err != nil {
			return nil, err
		}
		return results, nil
	}

	// デフォルトは部分一致
	matchStr := "%" + filterWord.FilterWord + "%"
	if filterWord.MatchingPattern != nil && *filterWord.MatchingPattern == models.MatchingPatternExactMatch {
		matchStr = filterWord.FilterWord
	}

	todo := TableName(&Todo{})
	user := TableName(&User{})

	// MEMO: ある程度複雑になったら頑張らずに db.Row() で生SQLを書く方が保守性は高いかもしれない。（メソッド使っても生SQL部分は存在するし）
	res := d.db.
		Select("todo.id, todo.text, todo.done, todo.created_at, user.id AS user_id, user.name AS user_name").
		Table(todo).
		Joins(InnerJoin(user) + On("%s.id = %s.user_id", user, todo)).
		Where(Col(todo, "text").Like(matchStr)).
		Or(Col(user, "name").Like(matchStr)).
		Find(&results)
	if res.Error != nil {
		return nil, res.Error
	}

	return results, nil
}
