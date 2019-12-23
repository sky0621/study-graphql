package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/jinzhu/gorm"
	"github.com/sky0621/study-graphql/backend"

	_ "github.com/go-sql-driver/mysql"
)

const dataSource = "localuser:localpass@tcp(127.0.0.1:3306)/localdb?charset=utf8&parseTime=True&loc=Local"
const defaultPort = "5050"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	defer func() {
		if db != nil {
			if err := db.Close(); err != nil {
				panic(err)
			}
		}
	}()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(backend.NewExecutableSchema(backend.Config{Resolvers: &backend.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
