package main

import (
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/sky0621/study-graphql/try01/src/backend/graph"
	"github.com/sky0621/study-graphql/try01/src/backend/graph/generated"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	// MEMO: ローカルでしか使わないので、ベタ書き
	dsn := "host=localhost port=25432 dbname=study-graphql-local-db user=postgres password=yuckyjuice sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	boil.SetLocation(loc)

	r := chi.NewRouter()
	r.Use(corsHandlerFunc())
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query",
		handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{
					Resolvers: &graph.Resolver{
						DB: db,
					},
				},
			),
		),
	)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func corsHandlerFunc() func(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler
}
