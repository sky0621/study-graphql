package model

type Todo struct {
	// ID
	ID int64 `json:"id"`
	// TODO
	Task string `json:"task"`
	// ユーザーID
	UserID int64 `json:"user_id"`
}

func (Todo) IsNode() {}

// ページングを伴う結果返却用
type TodoConnection struct {
	// ページ情報
	PageInfo *PageInfo `json:"pageInfo"`
	// 検索結果一覧（※カーソル情報を含む）
	Edges []*TodoEdge `json:"edges"`
	// 検索結果の全件数
	TotalCount int `json:"totalCount"`
}

func (TodoConnection) IsConnection() {}

// 検索結果一覧（※カーソル情報を含む）
type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor string `json:"cursor"`
}

func (TodoEdge) IsEdge() {}
