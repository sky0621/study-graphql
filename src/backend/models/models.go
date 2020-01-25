package models

type Todo struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"createdAt"`
	Done      bool   `json:"done"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
