package mutation

import (
	"Gopher-By-Example/graphql-example/web/vote"
	"log"

	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createVote": &graphql.Field{
			Type: vote.Vote,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"options": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(vote.OptionInput),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"deadline": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"class": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(vote.Class),
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				log.Println(p.Args)
				var params vote.CreateVoteParams
				params.Title = p.Args["title"].(string)
				if p.Args["description"] != nil {
					params.Description = p.Args["description"].(string)
				}
				params.Deadline = p.Args["deadline"].(string)
				params.Class = p.Args["class"].(int)
				return nil, nil

			},
		},
	},
})
