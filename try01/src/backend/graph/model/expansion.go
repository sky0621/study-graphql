package model

import "math"

func (c *TextFilterCondition) MatchString() string {
	if c == nil {
		return ""
	}
	if c.FilterWord == "" {
		return ""
	}
	matchStr := "%" + c.FilterWord + "%"
	if c.MatchingPattern == MatchingPatternExactMatch {
		matchStr = c.FilterWord
	}
	return matchStr
}

func (c *PageCondition) TotalPage(totalCount int64) int64 {
	if c == nil {
		return 0
	}
	targetCount := 0
	if c.Backward == nil && c.Forward == nil {
		targetCount = c.InitialLimit
	} else {
		if c.Backward != nil {
			targetCount = c.Backward.Last
		}
		if c.Forward != nil {
			targetCount = c.Forward.First
		}
	}
	return int64(math.Ceil(float64(totalCount) / float64(targetCount)))
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

func (o *EdgeOrder) CustomerOrderKeyExists() bool {
	if o == nil {
		return false
	}
	if o.Key == nil {
		return false
	}
	if o.Key.CustomerOrderKey == nil {
		return false
	}
	return o.Key.CustomerOrderKey.IsValid()
}
