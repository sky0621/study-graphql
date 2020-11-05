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
