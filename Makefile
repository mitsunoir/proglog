compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go_grpc_out=. \
		--go_opt=paths=source_relative \
		--go_grpc_opt=paths=source_relative \
		--proto_path=.

test:
	go test -race ./...
