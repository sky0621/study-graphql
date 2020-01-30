package database

import (
	"context"
	"time"

	"github.com/sky0621/study-graphql/src/backend/util"

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
	 * 文字列フィルタ"有り"
	 * ページング無し
	 * 並べ替え無し
	 */
	if filterWord.ExistsFilter() && pageCondition.NoPaging() && edgeOrder.NoSort() {
		// デフォルトは部分一致
		matchStr := filterWord.MatchString()

		todo := TableName(&Todo{})
		user := TableName(&User{})

		var results []*Todo
		res := d.db.
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

	/*
	 * 文字列フィルタ無し
	 * ページング"有り"
	 * 並べ替え無し
	 */
	if filterWord.NoFilter() && pageCondition.ExistsPaging() && edgeOrder.NoSort() {
		// 前ページ遷移指示
		if pageCondition.Backward != nil {
			// このカーソルをデコードした結果から todo のPKを取得してPK検索。その結果より前のレコードを検索対象とする
			cursor := pageCondition.Backward.Before
			if cursor == nil {
				return nil, nil
			}
			_, todoID, err := util.DecodeCursor(*cursor)
			if err != nil {
				return nil, err
			}

			// 比較対象カーソルに該当するレコードを取得
			var target *Todo
			if err := d.db.Where(&Todo{ID: todoID}).First(&target).Error; err != nil {
				return nil, err
			}

			// 並べ替え指示なしの場合は、デフォルトで「作成日時」の”降順”を指定

			todo := TableName(&Todo{})
			user := TableName(&User{})

			var results []*Todo
			res := d.db.
				Table(todo).
				Joins(InnerJoin(user)+On("%s.id = %s.user_id", user, todo)).
				Where("todo.created_at > ?", target.CreatedAt).
				Order("created_at DESC").
				Limit(pageCondition.Backward.Last).
				Find(&results)
			if res.Error != nil {
				return nil, res.Error
			}

			return results, nil
		}
		// 次ページ遷移指示
		if pageCondition.Forward != nil {
			// このカーソルより後のレコードを検索対象とする
			//cursor := pageCondition.Forward.After
		}
	}

	/*
	 * 文字列フィルタ無し
	 * ページング無し
	 * 並べ替え"有り"
	 */
	if filterWord.NoFilter() && pageCondition.NoPaging() && edgeOrder.ExistsSort() {
		// FIXME:
	}

	// FIXME: その他のパターン

	return nil, nil
}
