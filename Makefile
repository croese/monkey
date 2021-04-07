build:
	go build -o cmd/monkey/monkey cmd/monkey/main.go

runInterp:
	go run cmd/monkey/main.go

test:
	go test ./...