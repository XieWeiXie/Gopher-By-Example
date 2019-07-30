package mutation

import (
	"Gopher-By-Example/graphql-example/web/vote"
	"log"
	"strconv"

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
					Type: graphql.NewNonNull(graphql.NewList(vote.OptionInput)),
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
				var options []vote.OptionParams
				for _, i := range p.Args["options"].([]interface{}) {
					var one vote.OptionParams
					k := i.(map[string]interface{})
					one.Name = k["name"].(string)
					options = append(options, one)
				}
				params.Options = options
				log.Println(params)
				result, err := vote.CreateVote(params)
				if err != nil {
					return nil, err
				}
				return result, nil

			},
		},
		"updateVote": &graphql.Field{
			Type: vote.Vote,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				var params vote.UpdateVoteParams
				id := p.Args["id"]
				ID, _ := strconv.Atoi(id.(string))
				params.Id = int64(ID)
				params.Title = p.Args["title"].(string)
				params.Description = p.Args["description"].(string)
				return vote.UpdateOneVote(params)
			},
		},
	},
})
