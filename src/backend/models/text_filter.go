package models

import (
	"fmt"
	"io"
	"strconv"
)

// 文字列フィルタ条件
type TextFilterCondition struct {
	// フィルタ文字列
	FilterWord string `json:"filterWord"`
	// マッチングパターン（※オプション。指定無しの場合は「部分一致」となる。）
	MatchingPattern *MatchingPattern `json:"matchingPattern"`
}

func (c *TextFilterCondition) NoFilter() bool {
	return c == nil || c.FilterWord == ""
}

func (c *TextFilterCondition) ExistsFilter() bool {
	return !c.NoFilter()
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
