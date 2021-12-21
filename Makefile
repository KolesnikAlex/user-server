# make <commandName>

gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
 	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
 	grpc_user_service.proto

migration-up:
	goose postgres "user=postgres password=password dbname=postgres sslmode=disable" up

migration-down:
	goose postgres "user=postgres password=password dbname=postgres sslmode=disable" down
