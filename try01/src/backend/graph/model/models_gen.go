// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// ページングを伴う結果返却用
type Connection interface {
	IsConnection()
}

// 検索結果一覧（※カーソル情報を含む）
type Edge interface {
	IsEdge()
}

type Node interface {
	IsNode()
}

// 前ページ遷移条件
type BackwardPagination struct {
	// 取得件数
	Last int `json:"last"`
	// 取得対象識別用カーソル（※前ページ遷移時にこのカーソルよりも前にあるレコードが取得対象）
	Before string `json:"before"`
}

// ユーザー
type Customer struct {
	// ID
	ID string `json:"id"`
	// 名前
	Name string `json:"name"`
	// 年齢
	Age int `json:"age"`
	// Todo
	Todos []*Todo `json:"todos"`
}

func (Customer) IsNode() {}

// ページングを伴う結果返却用
type CustomerConnection struct {
	// ページ情報
	PageInfo *PageInfo `json:"pageInfo"`
	// 検索結果一覧（※カーソル情報を含む）
	Edges []*CustomerEdge `json:"edges"`
	// 検索結果の全件数
	TotalCount int64 `json:"totalCount"`
}

func (CustomerConnection) IsConnection() {}

// 検索結果一覧（※カーソル情報を含む）
type CustomerEdge struct {
	Node   *Customer `json:"node"`
	Cursor string    `json:"cursor"`
}

func (CustomerEdge) IsEdge() {}

// 並び替え条件
type EdgeOrder struct {
	// 並べ替えキー項目
	Key *OrderKey `json:"key"`
	// ソート方向
	Direction OrderDirection `json:"direction"`
}

// 次ページ遷移条件
type ForwardPagination struct {
	// 取得件数
	First int `json:"first"`
	// 取得対象識別用カーソル（※次ページ遷移時にこのカーソルよりも後ろにあるレコードが取得対象）
	After string `json:"after"`
}

// 並べ替えのキー
//
// 【検討経緯】
// 汎用的な構造、かつ、タイプセーフにしたく、interface で定義の上、機能毎に input ないし enum で実装しようとした。
// しかし、input は interface を実装できない仕様だったので諦めた。
// enum に継承機能があればよかったが、それもなかった。
// union で TodoOrderKey や（増えたら）他の機能の並べ替えのキーも | でつなぐ方法も考えたが、
// union を input に要素として持たせることはできない仕様だったので、これも諦めた。
// とはいえ、並べ替えも共通の仕組みとして提供したく、結果として機能毎の enum フィールドを共通の input 内に列挙していく形にした。
type OrderKey struct {
	// ユーザー一覧の並べ替えキー
	CustomerOrderKey *CustomerOrderKey `json:"customerOrderKey"`
	// TODO一覧の並べ替えキー
	TodoOrderKey *TodoOrderKey `json:"todoOrderKey"`
}

// ページング条件
type PageCondition struct {
	// 前ページ遷移条件
	Backward *BackwardPagination `json:"backward"`
	// 次ページ遷移条件
	Forward *ForwardPagination `json:"forward"`
	// 現在ページ番号（今回のページング実行前の時点のもの）
	NowPageNo int `json:"nowPageNo"`
	// １ページ表示件数
	InitialLimit int `json:"initialLimit"`
}

// ページ情報
type PageInfo struct {
	// 次ページ有無
	HasNextPage bool `json:"hasNextPage"`
	// 前ページ有無
	HasPreviousPage bool `json:"hasPreviousPage"`
	// 当該ページの１レコード目
	StartCursor string `json:"startCursor"`
	// 当該ページの最終レコード
	EndCursor string `json:"endCursor"`
}

// 文字列フィルタ条件
type TextFilterCondition struct {
	// フィルタ文字列
	FilterWord string `json:"filterWord"`
	// マッチングパターン
	MatchingPattern MatchingPattern `json:"matchingPattern"`
}

// TODO
type Todo struct {
	// ID
	ID string `json:"id"`
	// タスク
	Task string `json:"task"`
	// ユーザー情報
	Customer *Customer `json:"customer"`
}

func (Todo) IsNode() {}

// ページングを伴う結果返却用
type TodoConnection struct {
	// ページ情報
	PageInfo *PageInfo `json:"pageInfo"`
	// 検索結果一覧（※カーソル情報を含む）
	Edges []*TodoEdge `json:"edges"`
	// 検索結果の全件数
	TotalCount int64 `json:"totalCount"`
}

func (TodoConnection) IsConnection() {}

// 検索結果一覧（※カーソル情報を含む）
type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor string `json:"cursor"`
}

func (TodoEdge) IsEdge() {}

// ユーザー並べ替えキー
type CustomerOrderKey string

const (
	// ID
	CustomerOrderKeyID CustomerOrderKey = "ID"
	// ユーザー名
	CustomerOrderKeyName CustomerOrderKey = "NAME"
)

var AllCustomerOrderKey = []CustomerOrderKey{
	CustomerOrderKeyID,
	CustomerOrderKeyName,
}

func (e CustomerOrderKey) IsValid() bool {
	switch e {
	case CustomerOrderKeyID, CustomerOrderKeyName:
		return true
	}
	return false
}

func (e CustomerOrderKey) String() string {
	return string(e)
}

func (e *CustomerOrderKey) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CustomerOrderKey(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CustomerOrderKey", str)
	}
	return nil
}

func (e CustomerOrderKey) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// マッチングパターン種別（※要件次第で「前方一致」や「後方一致」も追加）
type MatchingPattern string

const (
	// 部分一致
	MatchingPatternPartialMatch MatchingPattern = "PARTIAL_MATCH"
	// 完全一致
	MatchingPatternExactMatch MatchingPattern = "EXACT_MATCH"
)

var AllMatchingPattern = []MatchingPattern{
	MatchingPatternPartialMatch,
	MatchingPatternExactMatch,
}

func (e MatchingPattern) IsValid() bool {
	switch e {
	case MatchingPatternPartialMatch, MatchingPatternExactMatch:
		return true
	}
	return false
}

func (e MatchingPattern) String() string {
	return string(e)
}

func (e *MatchingPattern) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MatchingPattern(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MatchingPattern", str)
	}
	return nil
}

func (e MatchingPattern) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// 並べ替え方向
type OrderDirection string

const (
	// 昇順
	OrderDirectionAsc OrderDirection = "ASC"
	// 降順
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// TODO並べ替えキー
type TodoOrderKey string

const (
	// ID
	TodoOrderKeyID TodoOrderKey = "ID"
	// TODO
	TodoOrderKeyTask TodoOrderKey = "TASK"
	// ユーザー名
	TodoOrderKeyUserName TodoOrderKey = "USER_NAME"
)

var AllTodoOrderKey = []TodoOrderKey{
	TodoOrderKeyID,
	TodoOrderKeyTask,
	TodoOrderKeyUserName,
}

func (e TodoOrderKey) IsValid() bool {
	switch e {
	case TodoOrderKeyID, TodoOrderKeyTask, TodoOrderKeyUserName:
		return true
	}
	return false
}

func (e TodoOrderKey) String() string {
	return string(e)
}

func (e *TodoOrderKey) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoOrderKey(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoOrderKey", str)
	}
	return nil
}

func (e TodoOrderKey) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
