start:
	go run ./bin/web

build:
	go build ./bin/web

test:
	go test -v ./...

coverage:
	go test -coverprofile='coverage.out' ./...
	go tool cover -html='coverage.out'
	del coverage.out

swagger:
	swag init -g ./bin/web/main.go