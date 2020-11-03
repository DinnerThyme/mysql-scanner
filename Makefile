test:
	go test -v ./...

build:
	go build -o mysql-scanner cmd/scanner/main.go

run:
	go run cmd/scanner/main.go