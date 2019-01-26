# Parameters
MAIN_FILE=main.go
SERVER=players_finder

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o $(SERVER) $(MAIN_FILE)

build-linux:
	GOARCH=amd64 GOOS=linux go build -o $(SERVER) $(MAIN_FILE)

run:
	./$(SERVER)

test:
	go test -race -cover -v ./...