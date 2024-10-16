.PHONY=build

build:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

build-parser:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/name_parser cmd/name_parser/main.go

build-sub:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/sub_commands cmd/sub_commands/main.go

build-self:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/self_define_cmd cmd/self_define_cmd/main.go

build-tour:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/tour cmd/tour/main.go

build-template:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/template_sample cmd/template_sample/main.go

run: build
	@./bin/main

coverage:
	@go test -v -cover ./...

test:
	@go test -v ./...

