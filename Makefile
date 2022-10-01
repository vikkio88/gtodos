build:
	go build -o bin/
run:
	go run main.go
test:
	go test ./...
clean:
	rm -rf bin/ db_files/