.PHONY: server
server: format
	go run main.go

.PHONY: format
format:
	gofmt -w .

.PHONY: build
build:
	go build -o bin/server .

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif

.PHONY: deps
deps: glide
	glide install

.PHONY: db/migrate
db/migrate:
	go run db/migrate.go
