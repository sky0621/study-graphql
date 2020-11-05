package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := sqlx.MustOpen("sqlite3", "./data.db")
	db.MustExec(schema)

	tx := db.MustBegin()

	users := []string{"Sato", "Suzuki", "Takahashi", "Tanaka", "Ito", "Watanabe", "Yamamoto", "Nakamura", "Kobayashi", "Kato"}
	for idx, user := range users {
		tx.MustExec(insertUser, idx+1, user)
	}

	todoID := 1
	for idx, _ := range users {
		for j := 0; j < 10; j++ {
			tx.MustExec(insertTodo, todoID, fmt.Sprintf("やること%2d", todoID), idx+1)
			todoID++
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

var schema = `
CREATE TABLE todo (
    id integer not null primary key ,
    task text not null,
    user_id integer not null
);

CREATE TABLE user (
    id integer not null primary key ,
    name text not null
);
`

var insertTodo = `INSERT INTO todo VALUES($1, $2, $3)`

var insertUser = `INSERT INTO user VALUES($1, $2)`
