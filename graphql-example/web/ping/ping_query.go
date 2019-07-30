package ping

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

var Ping = graphql.NewObject(graphql.ObjectConfig{
	Name: "ping",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if response, ok := p.Source.(ResponseForPing); ok {
					return response.Data, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"code": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if response, ok := p.Source.(ResponseForPing); ok {
					return response.Code, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
	},
})

type ResponseForPing struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}

var Default = ResponseForPing{
	Data: "pong",
	Code: http.StatusOK,
}

func MakeResponseForPing(data string) ResponseForPing {
	return ResponseForPing{
		Data: data,
		Code: http.StatusOK,
	}
}
