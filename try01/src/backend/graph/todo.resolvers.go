package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
	"github.com/sky0621/study-graphql/try01/src/backend/sqlboiler"
)

func (r *queryResolver) TodoConnection(ctx context.Context,
	filterWord *model.TextFilterCondition,
	pageCondition *model.PageCondition,
	edgeOrder *model.EdgeOrder) (*model.TodoConnection, error) {

	var mods []qm.QueryMod

	/*
	 * 検索文字列フィルタ設定
	 */
	if filterWord != nil {
		// なんで SQL Boiler には Like メソッドが無いんだろう・・・。
		mods = append(mods, qm.Where(sqlboiler.TodoColumns.Task+" LIKE ?", filterWord.MatchString()))
		// 上記、フィルタ対象のカラムを１つにしているが、要件に応じて、複数カラムを OR でつないでもいい。
	}

	/*
	 * １ページに表示する分の絞り込み条件追加前に件数を取得
	 */
	totalCount, err := sqlboiler.Todos(mods...).Count(ctx, r.DB)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	/*
	 * 並べ替え設定
	 */
	if edgeOrder.ExistsOrder() {
		key := sqlboiler.TodoColumns.ID
		// `todo` テーブルの何のカラムを並べ替えのキーにしているか
		if edgeOrder.Key.TodoOrderKey != nil {
			switch *edgeOrder.Key.TodoOrderKey {
			case model.TodoOrderKeyID:
				// MEMO: デフォルトキーにしてるので特に処理なし
			case model.TodoOrderKeyTask:
				key = sqlboiler.TodoColumns.Task
			case model.TodoOrderKeyUserName:
				// FIXME: 現状、 `user` テーブルの情報取得は dataloader 経由にしているため使用不可。
				// `todo` : `user` = N:1 を想定するなら、 `user` テーブルの情報を dataloader 経由でなく
				// inner join で取得する方式に変更すれば使用可能になるか。
			}
		}
		mods = append(mods, qm.OrderBy(fmt.Sprintf("%s %s", key, edgeOrder.Direction.String())))
	}

	/*
	 * ページング設定
	 */
	if pageCondition.ExistsPaging() {
		/*
		 * ページング指定無しの初期ページビュー
		 */
		if pageCondition.IsInitialPageView() {
			/*
			 * 表示件数指定がある場合
			 */
			if pageCondition.HasInitialLimit() {
				if edgeOrder.ExistsOrder() {
					switch edgeOrder.Direction {
					case model.OrderDirectionAsc:
					case model.OrderDirectionDesc:
					}
				}
			}
		}
	}

	mods = append(mods, qm.Limit(3))

	/*
	 * 検索実行
	 */
	todos, err := sqlboiler.Todos(mods...).All(ctx, r.DB)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	/*
	 * 結果の詰め替え
	 */
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
	if len(edges) == 0 {
		return emptyTodoConnection(), nil
	}

	// 検索結果全件数と１ページあたりの表示件数から、今回の検索による総ページ数を算出
	totalPage := pageCondition.TotalPage(totalCount)

	todoConn := &model.TodoConnection{
		PageInfo: &model.PageInfo{
			HasNextPage:     (totalPage - int64(pageCondition.MoveToPageNo())) >= 1, // 遷移後も、まだ先のページがあるか
			HasPreviousPage: pageCondition.MoveToPageNo() > 1,                       // 遷移後も、まだ前のページがあるか
			StartCursor:     edges[0].Cursor,
			EndCursor:       edges[len(edges)-1].Cursor,
		},
		Edges:      edges,
		TotalCount: int(totalCount),
	}
	return todoConn, nil
}

func emptyTodoConnection() *model.TodoConnection {
	return &model.TodoConnection{
		PageInfo: &model.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
	}
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
