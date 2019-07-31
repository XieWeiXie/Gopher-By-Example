package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/graphql-go/handler"
)

func Auth(ctx context.Context, next *handler.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		bearer := strings.Split(auth, " ")
		if len(bearer) != 2 {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("Add Authorization")))
			return
			ctx = context.WithValue(ctx, "Token", bearer[1])
		}
		next.ContextHandler(ctx, writer, request)
	}
}
