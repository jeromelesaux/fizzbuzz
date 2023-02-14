CC=go
RM=rm
GIT_COMMIT := $(shell git rev-list -1 HEAD)
VERSION_DATE := $(shell date "+%Y-%m-%d_%H:%M:%S")
HTTP_PORT=8080

build: swagger
		$(CC)  build  -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.VersionDate=${VERSION_DATE}" -o fizzbuzz main.go

test:
		$(CC) test ./... -cover

lint:
		golangci-lint run --timeout 2m0s

clean:
		$(RM) -f fizzbuzz

swagger:
		go install github.com/swaggo/swag/cmd/swag@latest
		swag init

docker:
		docker build . -t fizzbuzz

docker-run: docker
		docker run -e PORT="${HTTP_PORT}" -p ${HTTP_PORT}:${HTTP_PORT} -i  fizzbuzz

help:
	@echo "	Makefile help"
	@echo "	possible actions :"
	@echo ""
	@echo " make build (simple local server build)"
	@echo " make test (execute local unit tests)"
	@echo " make lint (execute linter on the project)"
	@echo " make clean (clean workspace)"
	@echo " make swagger (build swagger contract)"
	@echo " make docker (build docker image for the app)"
	@echo " make docker-run (build and run docker app)"
