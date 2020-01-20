package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"

	"github.com/99designs/gqlgen/handler"
	"github.com/jinzhu/gorm"
	"github.com/sky0621/study-graphql/src/backend"

	_ "github.com/lib/pq"
)

const dataSourceFormat = "dbname=%s user=%s password=%s sslmode=disable"
const defaultPort = "5050"

func init() {
	viper.SetEnvPrefix("STUDY_GRAPHQL")
	viper.AutomaticEnv()
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("PGPASSWORD", "localpass")
	viper.SetDefault("POSTGRES_DB", "localdb")
}

func main() {
	u := viper.Get("POSTGRES_USER")
	fmt.Println(u)

	dataSource := os.Getenv("CLOUDSQL_DATASOURCE")
	if dataSource == "" {
		//dataSource = defaultDataSource
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open("postgres", dataSource)
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
	db.LogMode(true)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(backend.NewExecutableSchema(backend.Config{Resolvers: &backend.Resolver{DB: db}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
