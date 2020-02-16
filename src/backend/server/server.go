package main

import (
	"net/http"
	"os"

	"github.com/sky0621/study-graphql/src/backend"

	"github.com/go-chi/chi"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const defaultDataSource = "localuser:localpass@tcp(localhost:3306)/localdb?charset=utf8&parseTime=True&loc=Local"
const defaultPort = "5050"

func main() {
	/*
	 * RDB（Cloud SQL）
	 */
	dataSource := os.Getenv("CLOUDSQL_DATASOURCE")
	if dataSource == "" {
		dataSource = defaultDataSource
	}
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
	db.LogMode(true)

	/*
	 * タスクキュー（Cloud Tasks）
	 */
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		os.Exit(-1)
	}
	locationID := os.Getenv("GCP_LOCATION_ID")
	if locationID == "" {
		os.Exit(-1)
	}
	credentialPath := os.Getenv("GCP_CREDENTIAL_PATH")
	if credentialPath == "" {
		os.Exit(-1)
	}
	queueIDMap := map[backend.CloudTasksQueueKind]string{
		backend.CloudTasksQueueKindTodo: "my-queue",
	}
	taskExecURL := os.Getenv("GCP_CLOUDTASK_EXEC_URL")
	if taskExecURL == "" {
		os.Exit(-1)
	}

	gcpClient := backend.NewGCPClientWrapper(projectID, locationID, credentialPath, taskExecURL, queueIDMap)

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Handle("/", playgroundHandler())
	r.Handle("/query", graphqlHandler(db, gcpClient))

	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}

func playgroundHandler() http.HandlerFunc {
	h := handler.Playground("study-graphql-playground", "/query")
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func graphqlHandler(db *gorm.DB, gcpClient *backend.GCPClientWrapper) http.HandlerFunc {
	h := handler.GraphQL(backend.NewExecutableSchema(backend.Config{Resolvers: &backend.Resolver{
		DB:        db,
		GCPClient: gcpClient,
	}}))
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
