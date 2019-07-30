package query

import (
	"Gopher-By-Example/graphql-example/web/ping"
	"Gopher-By-Example/graphql-example/web/vote"
	"strconv"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ping": &graphql.Field{},
	},
})

func init() {
	Query.AddFieldConfig("ping", &graphql.Field{
		Type: ping.Ping,
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			if p.Args["data"] == nil {
				return ping.Default, nil
			}
			return ping.MakeResponseForPing(p.Args["data"].(string)), nil
		},
	})
}

func init() {
	Query.AddFieldConfig("vote", &graphql.Field{
		Type: vote.Vote,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id := p.Args["id"]
			ID, _ := strconv.Atoi(id.(string))
			return vote.GetOneVote(int64(ID))
		},
	})
}
