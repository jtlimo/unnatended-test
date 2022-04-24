start:
	make build
	go run main.go
build:
	go build
test:
	go test ./...