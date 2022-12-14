.PHONY: protogen
protogen:
	protoc --proto_path=proto \
	--go_out=pkg --go_opt=paths=source_relative \
	--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
	./proto/tagpro/services/calculator/messages.proto

.PHONY: build
build:
	go build -o bin/calculator ./server/cmd/calculator

.PHONY: run
run: build
	./bin/calculator