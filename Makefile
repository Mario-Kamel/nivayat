build:
	go build -o main.exe

run: build
	./main.exe

test:
	go test -v ./...