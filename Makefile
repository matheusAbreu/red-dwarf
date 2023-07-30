SWAG = $(shell go env GOPATH)/bin/swag

deps/install:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	$(SWAG) init -generalInfo ./main.go -output ./docs/swagger
