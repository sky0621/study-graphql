package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
	"github.com/sky0621/study-graphql/try01/src/backend/sqlboiler"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	orderKey := sqlboiler.TodoColumns.ID
	if edgeOrder.ExistsOrder() {
		// `todo` テーブルの何のカラムを並べ替えのキーにしているか
		if edgeOrder.Key.TodoOrderKey != nil {
			switch *edgeOrder.Key.TodoOrderKey {
			case model.TodoOrderKeyID:
				// MEMO: デフォルトキーにしてるので特に処理なし
			case model.TodoOrderKeyTask:
				orderKey = sqlboiler.TodoColumns.Task
			case model.TodoOrderKeyUserName:
				// FIXME: 現状、 `user` テーブルの情報取得は dataloader 経由にしているため使用不可。
				// `todo` : `user` = N:1 を想定するなら、 `user` テーブルの情報を dataloader 経由でなく
				// inner join で取得する方式に変更すれば使用可能になるか。
			}
		}
		mods = append(mods, qm.OrderBy(fmt.Sprintf("%s %s", orderKey, edgeOrder.Direction.String())))
	}

	/*
	 * ページング設定
	 */
	if pageCondition.IsInitialPageView() {
		/*
		 * ページング指定無しの初期ページビュー
		 */
		// 表示件数が指定されている場合
		if pageCondition.HasInitialLimit() {
			mods = append(mods, qm.Limit(*pageCondition.InitialLimit))
		}
	} else {
		/*
		 * 前ページへの遷移指示
		 */
		if pageCondition.Backward != nil {
			// 表示件数が指定されている場合
			if pageCondition.Backward.Last > 0 {
				mods = append(mods, qm.Limit(pageCondition.Backward.Last))
			}
			// 1, 2, 3, [[4], 5, 6], 7, 8, 9

			// 9, 8, 7, [[6], 5, 4], 3, 2, 1

		}
		/*
		 * 次ページへの遷移指示
		 */
		if pageCondition.Forward != nil {
			// 表示件数が指定されている場合
			if pageCondition.Forward.First > 0 {
				mods = append(mods, qm.Limit(pageCondition.Forward.First))
			}

		}
	}

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
			// 単なるページングだけでなく、指定のキーで昇順・降順の並べ替えをする要件がある場合、
			// カーソルに含める要素を「その時、並べ替えのキーに指定されている要素」にすると、
			// 次回ページング時、カーソルをデコードした要素よりも前（ないし後）という条件指定が可能。
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

// フィルタの結果、０件だった場合の空リターン用
func emptyTodoConnection() *model.TodoConnection {
	return &model.TodoConnection{
		PageInfo: &model.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
	}
}

func selectCursorKey(orderKey string) {
	// FIXME:
}

// todo : user = N : 1 を想定するならあえて dataloader でなくてもよいか？
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj == nil {
		return nil, nil
	}
	return For(ctx).UsersByIDs.Load(obj.UserID)
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
