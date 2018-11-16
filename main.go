package main

import (
	"fmt"
	"github.com/edwinharly/go-graphql-api/gql"
	"github.com/edwinharly/go-graphql-api/postgres"
	"github.com/edwinharly/go-graphql-api/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	fatal(err)

	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	dbport, err := strconv.Atoi(os.Getenv("DBPORT"))
	fatal(err)

	db, err := postgres.New(
		postgres.ConnString(
			os.Getenv("DBHOST"),
			dbport,
			os.Getenv("DBUSER"),
			os.Getenv("DBPASS"),
			os.Getenv("DBNAME"),
		),
	)
	fatal(err)

	rootQuery := gql.NewRoot(db)
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQL())

	return router, db
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
