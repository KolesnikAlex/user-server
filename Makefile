
gen-proto:
	$protoc --go_out=. --go_opt=paths=source_relative \
 	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
 	grpc/grpc_user_service.proto
