run:
	go run main.go
build:
	go build main.go
test:
	go clean -testcache
	go test ./...
cover:
	go test --short -coverprofile="coverage.out" ./...
	go tool cover -html="coverage.out"
	rm coverage.out