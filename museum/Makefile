.PHONY: all
all: build
FORCE: ;

.PHONY: build

install:
	npm install
	go mod tidy

tidy:
	go mod tidy

generate:
	go generate ./...

run:
	go run cmd/museum/main.go

build:
	go build -o dist/museum cmd/museum/main.go

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "netgo" -installsuffix netgo -o dist/museum cmd/museum/main.go

clean:
	@rm -rf dist tmp
