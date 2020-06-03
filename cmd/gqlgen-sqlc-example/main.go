package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cleyonpiv/gqlgen-sqlc-example/dataloaders"
	"github.com/cleyonpiv/gqlgen-sqlc-example/gqlgen"
	"github.com/cleyonpiv/gqlgen-sqlc-example/pg"
)

func main() {
	// initialize the db
	// Format: host=foo port=5432 user=bar password=baz dbname=qux sslmode=disable"
	db, err := pg.Open("dbname=gqlgen_sqlc_example_db user=gqlgen sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize the repository
	repo := pg.NewRepository(db)

	// initialize the dataloaders
	dl := dataloaders.NewRetriever() // <- here we initialize the dataloader.Retriever

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	dlMiddleware := dataloaders.Middleware(repo)     // <- here we initialize the middleware
	queryHandler := gqlgen.NewHandler(repo, dl)      // <- use dataloader.Retriever here
	mux.Handle("/query", dlMiddleware(queryHandler)) // <- use dataloader.Middleware here

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
