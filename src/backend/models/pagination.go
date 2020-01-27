package models

// ページング条件
type PageCondition struct {
	// 前ページ遷移条件
	Backward *BackwardPagination `json:"backward"`
	// 次ページ遷移条件
	Forward *ForwardPagination `json:"forward"`
	// 現在ページ番号（今回のページング実行前の時点のもの）
	NowPageNo int `json:"nowPageNo"`
	// １ページ表示件数
	InitialLimit *int `json:"initialLimit"`
}

// 前ページ遷移条件
type BackwardPagination struct {
	// 取得件数
	Last int `json:"last"`
	// 取得対象識別用カーソル（※前ページ遷移時にこのカーソルよりも前にあるレコードが取得対象）
	Before *string `json:"before"`
}

// 次ページ遷移条件
type ForwardPagination struct {
	// 取得件数
	First int `json:"first"`
	// 取得対象識別用カーソル（※次ページ遷移時にこのカーソルよりも後ろにあるレコードが取得対象）
	After *string `json:"after"`
}
