init:
	go mod tidy

dev:
	go mod tidy
	go run main.go

test:
	go test ./...

build:
	go build -o ccwc

count_bytes:
	./ccwc -c example.txt

count_lines:
	./ccwc -l example.txt

count_words:
	./ccwc -w example.txt