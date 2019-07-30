package router

import (
	"Gopher-By-Example/graphql-example/web/mutation"
	"Gopher-By-Example/graphql-example/web/query"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func RegisterSchema() *graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query.Query,
			Mutation: mutation.Mutation,
		})
	if err != nil {
		panic("schema init fail")
	}
	return &schema

}

func Register() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   RegisterSchema(),
		Pretty:   true,
		GraphiQL: true,
	})
}
func StartWebServer() {
	log.Println("Start Web Server...")
	http.Handle("/graphql", Register())
	log.Fatal(http.ListenAndServe(":7878", nil))
}
