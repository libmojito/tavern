grpc_targets := proto/command.pb.go proto/command_grpc.pb.go

.PHONY: build
build: test vet
	go build .

$(grpc_targets): proto/command.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/command.proto

.PHONY: test
test: $(grpc_targets)
	go test .

.PHONY: vet
vet: $(grpc_targets)
	go vet

.PHONY: clean
clean:
	rm proto/*.pb.go
	go clean ./...

.PHONY: update
update:
	go get -u
	go mod tidy

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


.PHONY: update-examples
update-examples:
	go -C examples/openai get -u
	go -C examples/openai vet
	go -C examples/hello get -u
	go -C examples/hello vet
