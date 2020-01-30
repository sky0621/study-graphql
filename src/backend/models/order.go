package models

import (
	"fmt"
	"io"
	"strconv"
)

// 並び替え条件
type EdgeOrder struct {
	// 並べ替えキー項目
	Key *OrderKey `json:"key"`
	// ソート方向
	Direction OrderDirection `json:"direction"`
}

func (o *EdgeOrder) NoSort() bool {
	return o == nil
}

func (o *EdgeOrder) ExistsSort() bool {
	return !o.NoSort()
}

// 並べ替えのキー
// 汎用的な構造にしたいが以下はGraphQLの仕様として不可だった。
// ・enum・・・汎化機能がない。
// ・interface・・・inputには実装機能がない。
// ・union・・・inputでは要素に持てない。
// とはいえ、並べ替えも共通の仕組みとして提供したく、結果として機能毎に enum フィールドを列挙
type OrderKey struct {
	// TODO一覧の並べ替えキー
	TodoOrderKey *TodoOrderKey `json:"todoOrderKey"`
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
