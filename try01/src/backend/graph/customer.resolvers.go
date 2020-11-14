package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/sky0621/study-graphql/try01/src/backend/boiled"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *customerResolver) Todos(ctx context.Context, obj *model.Customer) ([]*model.Todo, error) {
	// FIXME: dataloader で実装
	return []*model.Todo{}, nil
}

func (r *queryResolver) CustomerConnection(ctx context.Context, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder, filterWord *model.TextFilterCondition) (*model.CustomerConnection, error) {
	/*
	 * SQL構築に必要な各種要素の保持用
	 */
	params := searchParam{
		// 情報取得先のテーブル名
		tableName: boiled.TableNames.Customer,

		// 検索文字列フィルタ未指定時のデフォルト（※とりあえずSQL文を固定化しているのでダミー条件を指定）
		baseCondition: "true",

		// 並び順のデフォルトはIDの降順
		orderKey:       boiled.CustomerColumns.ID,
		orderDirection: model.OrderDirectionDesc.String(),

		// 表示件数指定無しの場合でもパフォーマンス観点からMax件数は指定
		limit: 1000,
	}

	/*
	 * ページング後の次ページ、前ページの存在有無判定のために必要な
	 * 検索文字列フィルタ適用後の結果件数保持用
	 */
	var totalCount int64 = 0
	{
		var err error
		if filterWord == nil {
			totalCount, err = boiled.Customers().Count(ctx, r.DB)
		} else {
			predicate := filterWord.MatchString()
			totalCount, err = boiled.Customers(qm.Where(boiled.CustomerColumns.Name+" LIKE ?", predicate)).Count(ctx, r.DB)
			/*
			 * 検索文字列フィルタをSQLに適用
			 * TODO: 複数カラムにフィルタを適用したい場合など、ここで AND でつなぐか buildSearchQueryMod() を拡張するか検討が必要
			 */
			params.baseCondition = fmt.Sprintf("%s LIKE '%s'", boiled.CustomerColumns.Name, predicate)
		}
		if err != nil {
			log.Print(err)
			return nil, err
		}
	}

	if edgeOrder.CustomerOrderKeyExists() {
		params.orderKey = edgeOrder.Key.CustomerOrderKey.String()
		params.orderDirection = edgeOrder.Direction.String()
	}

	params.compareSymbol = compareSymbolGt
	params.decodedCursor = 0

	/*
	 * 検索実行
	 */
	var records []*CustomerWithRowNum
	if err := boiled.Customers(buildSearchQueryMod(params)).Bind(ctx, r.DB, &records); err != nil {
		log.Print(err)
		return nil, err
	}

	/*
	 * Relay返却形式
	 */
	result := &model.CustomerConnection{
		TotalCount: totalCount,
	}

	/*
	 * 検索結果をEdgeスライス形式に変換
	 */
	var edges []*model.CustomerEdge
	for _, record := range records {
		edges = append(edges, &model.CustomerEdge{
			Node: &model.Customer{
				ID:   strconv.Itoa(int(record.ID)),
				Name: record.Name,
				Age:  record.Age,
			},
			Cursor: createCursor("customer", record.RowNum),
		})
	}
	result.Edges = edges

	// 検索結果全件数と１ページあたりの表示件数から、今回の検索による総ページ数を算出
	totalPage := pageCondition.TotalPage(totalCount)

	/*
	 * クライアント側での画面表示及び次回ページングに必要な情報
	 */
	pageInfo := &model.PageInfo{
		HasNextPage:     (totalPage - int64(pageCondition.MoveToPageNo())) >= 1, // 遷移後も、まだ先のページがあるか
		HasPreviousPage: pageCondition.MoveToPageNo() > 1,                       // 遷移後も、まだ前のページがあるか
	}
	if len(edges) > 0 {
		pageInfo.StartCursor = edges[0].Cursor
		pageInfo.EndCursor = edges[len(edges)-1].Cursor
	}
	result.PageInfo = pageInfo

	return result, nil
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

type customerResolver struct{ *Resolver }
