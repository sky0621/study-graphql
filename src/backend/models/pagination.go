package models

import "math"

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

func (c *PageCondition) TotalPage(totalCount int64) int64 {
	if c == nil {
		return 0
	}
	if c.InitialLimit == nil {
		return 0
	}
	return int64(math.Ceil(float64(totalCount) / float64(*c.InitialLimit)))
}

func (c *PageCondition) MoveToPageNo() int {
	if c == nil {
		return 1 // 想定外のため初期ページ
	}
	if c.Backward == nil && c.Forward == nil {
		return c.NowPageNo // 前にも後ろにも遷移しないので
	}
	if c.Backward != nil {
		if c.NowPageNo <= 2 {
			return 1
		}
		return c.NowPageNo - 1
	}
	if c.Forward != nil {
		return c.NowPageNo + 1
	}
	return 1 // 想定外のため初期ページ
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
