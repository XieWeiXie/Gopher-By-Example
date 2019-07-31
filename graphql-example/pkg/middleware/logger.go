package middleware

import (
	"Gopher-By-Example/graphql-example/pkg/log"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/graphql-go/handler"
)

type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

func LoggerMiddleWare(ctx context.Context, h *handler.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		message := fmt.Sprintf("%s | %s | %s | %s", request.Method, request.Host, request.RequestURI, time.Now().Format(time.RFC3339))
		log_for_project.Println(message)
		bodyBytes, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		var opts RequestOptions
		json.Unmarshal(bodyBytes, &opts)
		log_for_project.Println(opts.Query)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		h.ContextHandler(ctx, writer, request)
	}

}
