run:
	@go run cmd/api/main.go

tests:
	@go test ./...

update:
	@go get .

dependencies:
	@go list -m all
	
build:
	@go build -o bin/main cmd/api/main.go
