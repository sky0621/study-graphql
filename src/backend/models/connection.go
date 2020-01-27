package models

// ページングを伴う結果返却用
type Connection interface {
	IsConnection()
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

// 検索結果一覧（※カーソル情報を含む）
type Edge interface {
	IsEdge()
}
