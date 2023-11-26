BINARY_NAME=video-concatter

build:
	go build -o bin/${BINARY_NAME} cmd/main.go

run: build
	./bin/${BINARY_NAME}

clean:
	go clean
	rm bin/${BINARY_NAME}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

lint:
	golangci-lint run
