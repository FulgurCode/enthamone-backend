build:
	go build -o ./bin/enthamone cmd/main.go

run: build
	./bin/enthamone
