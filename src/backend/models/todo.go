package models

import (
	"fmt"
	"io"
	"strconv"
)

type Todo struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"createdAt"`
	Done      bool   `json:"done"`
	// 本来 Todo の直接の属性でない User を持つのはGraphQLとしては悪手のはずだが、
	// 文字列フィルタの検索対象カラムに user.name を含めるには todo テーブルへのSELECT発行時に user テーブルもJOINするのが実装方法としてはシンプル。
	// todo テーブル検索後に１件ずつ user テーブルを検索していく方法で user.name をフィルタすると、
	// 全SQL発行後に残った件数を見て、１ページに必要な件数を満たすまで追加検索を行うことになり、処理が相当煩雑になる恐れがある。
	// ※dataloadenを使ったN+1問題解消にトライする際に、よりよい実装方法を再検討する。
	User *User `json:"user"`
}

func (t *Todo) IsNode() {}

// TODO並べ替えキー
type TodoOrderKey string

const (
	// TODO
	TodoOrderKeyText TodoOrderKey = "TEXT"
	// 済みフラグ
	TodoOrderKeyDone TodoOrderKey = "DONE"
	// 作成日時（初期表示時のデフォルト）
	TodoOrderKeyCreatedAt TodoOrderKey = "CREATED_AT"
	// ユーザー名
	TodoOrderKeyUserName TodoOrderKey = "USER_NAME"
)

var AllTodoOrderKey = []TodoOrderKey{
	TodoOrderKeyText,
	TodoOrderKeyDone,
	TodoOrderKeyCreatedAt,
	TodoOrderKeyUserName,
}

func (e TodoOrderKey) IsValid() bool {
	switch e {
	case TodoOrderKeyText, TodoOrderKeyDone, TodoOrderKeyCreatedAt, TodoOrderKeyUserName:
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
