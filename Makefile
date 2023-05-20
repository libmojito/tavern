.PHONY: build
build: vet test
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/command.proto
	go build .

.PHONY: test
test:
	go test .

.PHONY: vet
vet:
	go vet

.PHONY: clean
clean:
	rm proto/*.pb.go
	go clean ./...

.PHONY: examples
examples:
	protoc --go_out=examples/openai --go_opt=paths=source_relative \
		--go-grpc_out=examples/openai --go-grpc_opt=paths=source_relative \
		proto/command.proto
	go build -C examples/openai
	protoc --go_out=examples/hello --go_opt=paths=source_relative \
		--go-grpc_out=examples/hello --go-grpc_opt=paths=source_relative \
		proto/command.proto
	go build -C examples/hello

.PHONY: run-examples
run-examples: examples
	go run -C examples/openai . serve --port 50052 & \
	go run -C examples/hello . serve --port 50051
