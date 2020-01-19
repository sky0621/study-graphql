package models

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
