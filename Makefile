build:
	go build -o ./bin/atta cmd/main.go

run: build
	./bin/atta
