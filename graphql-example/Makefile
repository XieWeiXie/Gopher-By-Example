PROJECT=graphql_example




default:
	@go build -o ${PROJECT} main.go


run:
	@go run main.go

sync: default
	@chmod +x ${PROJECT}
	@./${PROJECT} sync

.PHONY: default run sync
