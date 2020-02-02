package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sky0621/study-graphql/src/backend/util"

	"github.com/sky0621/study-graphql/src/backend/models"

	"github.com/jinzhu/gorm"
)

type TodoForInput struct {
	ID        string    `gorm:"column:id;primary_key"`
	Text      string    `gorm:"column:text"`
	Done      bool      `gorm:"column:done"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UserID    string    `gorm:column:user_id`
}

func (t *TodoForInput) TableName() string {
	return "todo"
}

// Tableを実装するマーカーインタフェース
func (t *TodoForInput) IsTable() bool {
	return true
}

type Todo struct {
	ID        string    `gorm:"column:id;primary_key"`
	Text      string    `gorm:"column:text"`
	Done      bool      `gorm:"column:done"`
	CreatedAt time.Time `gorm:"column:created_at"`
	User
}

func (t *Todo) Columns() string {
	tn := t.TableName()
	return fmt.Sprintf("%s.id, %s.text, %s.done, %s.created_at", tn, tn, tn, tn)
}

func (t *Todo) TableName() string {
	return "todo"
}

// Tableを実装するマーカーインタフェース
func (t *Todo) IsTable() bool {
	return true
}

type TodoDao interface {
	InsertOne(ctx context.Context, u *TodoForInput) error
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

func (d *todoDao) InsertOne(ctx context.Context, u *TodoForInput) error {
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

func (d *todoDao) FindByCondition(ctx context.Context, filterCondition *models.TextFilterCondition, pageCondition *models.PageCondition, edgeOrder *models.EdgeOrder) ([]*Todo, error) {
	// todoテーブルのテーブル名
	todo := TableName(&Todo{})
	// userテーブルのテーブル名
	user := TableName(&User{})

	// 条件によらず固定の部分
	base := d.db.Table(todo).
		Joins(InnerJoin(user) + On("%s.id = %s.user_id", user, todo)).
		Select((&Todo{}).Columns() + ", " + (&User{}).Columns())

	/*
	 * 文字列フィルタ条件が指定されていた場合
	 */
	if filterCondition.ExistsFilter() {
		// デフォルトは部分一致(例：「テスト」 -> 「%テスト%」)
		matchStr := filterCondition.MatchString()
		// todoテーブルのtextカラムかuserテーブルのnameカラムとLIKE検索
		base = base.Where(Col(todo, "text").Like(matchStr)).Or(Col(user, "name").Like(matchStr))
	}

	/*
	 * ページング指定無しの初期ページビュー
	 */
	if pageCondition.IsInitialPageView() {
		if pageCondition.HasInitialLimit() {
			base = base.Limit(*pageCondition.InitialLimit)
		}
	}

	/*
	 * 並べ替え条件が指定されていた場合
	 */
	if edgeOrder.ExistsOrder() {
		orderKey := edgeOrder.Key.TodoOrderKey.Val()
		if orderKey != "" {
			// TODO: テーブル名のエイリアスを付けていないので同じカラムを持つテーブルを複数JOIN時に困る
			base = base.Order(fmt.Sprintf("%s %s", orderKey, edgeOrder.Direction.String()))
		}
	}

	/*
	 * ページング条件が指定されていた場合
	 * （※並べ替えのキー項目と昇順・降順の指定がないとページング不可のため、if文の判定に追加）
	 */
	if pageCondition.ExistsPaging() && edgeOrder.ExistsOrder() {
		col, err := getColumnNameByOrderKey(*edgeOrder.Key.TodoOrderKey)
		if err != nil {
			return nil, err
		}
		/*
		 * どの項目で並べ替えをしているかによって、ページ遷移のために比較対象レコードのどのカラムと比較するかが決まる。
		 * また、比較対象レコードのカラムと比較する時、当該カラムの昇順で並んでいるか降順で並んでいるかによって「 > 」にするか「 < 」にするかが変わる。
		 *
		 * 【説明】
		 * 　1 〜 17 までの数値(カラム名は「col」とする)が１ページ５件”昇順”で並んでいて、現在２ページ目を表示していたとする。
		 */
		switch edgeOrder.Direction {
		/*
		 *　★ 7, 8, 9, 10, 11 の昇順で並んでいる場合
		 */
		case models.OrderDirectionAsc:
			/*
			 * 次ページに遷移する場合、12, 13, 14, 15, 16 を取得する条件にする必要がある。
			 * pageCondition.Forward.Afterが、今表示している一覧の"最終行"を示すカーソルなので、そこから「 11 」という数値が取得できる。
			 * 結果、「col > 11」を５件取得する条件を追加すればいい。
			 */
			if pageCondition.Forward != nil {
				// 「このレコードよりも後のレコードを取得」という条件に使うための比較対象レコードを取得
				target, err := d.getCompareTarget(pageCondition.Forward.After)
				if err != nil {
					return nil, err
				}
				targetValue := getTargetValueByOrderKey(*edgeOrder.Key.TodoOrderKey, target)
				if targetValue == nil {
					return nil, errors.New("no target value")
				}
				base = base.Where(col.GreaterThan(targetValue)).Order(col_ASC(edgeOrder)).Limit(pageCondition.Forward.First)
			}

			/*
			 * 前ページに遷移する場合、2, 3, 4, 5, 6 を取得する条件にする必要がある。
			 * pageCondition.Backward.Beforeが、今表示している一覧の"１行目"を示すカーソルなので、そこから「 7 」という数値が取得できる。
			 * 結果、「num < 7」を５件取得する条件を追加すればいい。
			 * ※ただし、並べ替え順を”昇順”のままにすると、小さいものから５件取得する条件である都合上、意図に反して 1, 2, 3, 4, 5 が取得される。
			 *  そのため、いったん”降順”で並べ替えて取得した後、再度、”昇順”で並べ替え直す必要がある。
			 */
			if pageCondition.Backward != nil {
				// 「このレコードよりも前のレコードを取得」という条件に使うための比較対象レコードを取得
				target, err := d.getCompareTarget(pageCondition.Backward.Before)
				if err != nil {
					return nil, err
				}
				targetValue := getTargetValueByOrderKey(*edgeOrder.Key.TodoOrderKey, target)
				if targetValue == nil {
					return nil, errors.New("no target value")
				}
				subQuery := base.New()
				base = base.Where(subQuery.Where(col.LessThan(targetValue)).Order(col_DESC(edgeOrder)).Limit(pageCondition.Backward.Last).SubQuery()).
					Order(col_ASC(edgeOrder))
			}
		/*
		 *　★ 10, 9, 8, 7, 6 の降順で並んでいる場合
		 */
		case models.OrderDirectionDesc:
			/*
			 * 次ページに遷移する場合、15, 14, 13, 12, 11 を取得する条件にする必要がある。
			 * pageCondition.Forward.Afterが、今表示している一覧の"１行目"を示すカーソルなので、そこから「 10 」という数値が取得できる。
			 * 結果、「num > 10」を５件取得する条件を追加すればいい。
			 * ※ただし、並べ替え順を”降順”のままにすると、大きいものから５件取得する条件である都合上、意図に反して 16, 15, 14, 13, 12 が取得される。
			 *  そのため、いったん”昇順”で並べ替えて取得した後、再度、”降順”で並べ替え直す必要がある。
			 */
			if pageCondition.Forward != nil {
				// 「このレコードよりも後のレコードを取得」という条件に使うための比較対象レコードを取得
				target, err := d.getCompareTarget(pageCondition.Forward.After)
				if err != nil {
					return nil, err
				}
				targetValue := getTargetValueByOrderKey(*edgeOrder.Key.TodoOrderKey, target)
				if targetValue == nil {
					return nil, errors.New("no target value")
				}
				base = d.db.Where(base.Where(col.GreaterThan(targetValue)).Order(col_ASC(edgeOrder)).Limit(pageCondition.Forward.First).QueryExpr()).
					Order(col_DESC(edgeOrder))
			}
			/*
			 * 前ページに遷移する場合、5, 4, 3, 2, 1 を取得する条件にする必要がある。
			 * pageCondition.Backward.Beforeが、今表示している一覧の"最終行"を示すカーソルなので、そこから「 6 」という数値が取得できる。
			 * 結果、「num < 6」を５件取得する条件を追加すればいい。
			 */
			if pageCondition.Backward != nil {
				// 「このレコードよりも前のレコードを取得」という条件に使うための比較対象レコードを取得
				target, err := d.getCompareTarget(pageCondition.Backward.Before)
				if err != nil {
					return nil, err
				}
				targetValue := getTargetValueByOrderKey(*edgeOrder.Key.TodoOrderKey, target)
				if targetValue == nil {
					return nil, errors.New("no target value")
				}
				base = base.Where(col.GreaterThan(targetValue)).Order(col_DESC(edgeOrder)).Limit(pageCondition.Backward.Last)
			}
		}
	}

	var results []*Todo
	if err := base.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (d *todoDao) getCompareTarget(cursor *string) (*Todo, error) {
	if cursor == nil {
		return nil, errors.New("cursor is nil")
	}
	_, todoID, err := util.DecodeCursor(*cursor)
	if err != nil {
		return nil, err
	}

	// 比較対象カーソルに該当するレコードを取得
	var target Todo
	if err := d.db.Where(&Todo{ID: todoID}).First(&target).Error; err != nil {
		return nil, err
	}
	return &target, nil
}

func getColumnNameByOrderKey(todoOrderKey models.TodoOrderKey) (*c, error) {
	// todoテーブルのテーブル名
	todo := TableName(&Todo{})
	// userテーブルのテーブル名
	user := TableName(&User{})

	switch todoOrderKey {
	case models.TodoOrderKeyText:
		return Col(todo, "text"), nil
	case models.TodoOrderKeyDone:
		return Col(todo, "done"), nil
	case models.TodoOrderKeyCreatedAt:
		return Col(todo, "created_at"), nil
	case models.TodoOrderKeyUserName:
		return Col(user, "name"), nil
	default:
		return nil, errors.New("not target orderKey")
	}
}

func getTargetValueByOrderKey(todoOrderKey models.TodoOrderKey, todo *Todo) interface{} {
	switch todoOrderKey {
	case models.TodoOrderKeyText:
		return todo.Text
	case models.TodoOrderKeyDone:
		return todo.Done
	case models.TodoOrderKeyCreatedAt:
		return todo.CreatedAt
	case models.TodoOrderKeyUserName:
		return todo.Name
	default:
		return nil
	}
}

func col_ASC(o *models.EdgeOrder) string {
	return fmt.Sprintf("%s %s", o.Key.TodoOrderKey.Val(), models.OrderDirectionAsc.String())
}

func col_DESC(o *models.EdgeOrder) string {
	return fmt.Sprintf("%s %s", o.Key.TodoOrderKey.Val(), models.OrderDirectionDesc.String())
}
