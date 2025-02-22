# grpc-server-client-app


To generate the protobuf files

protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative protobuf-spec/customer-spec.proto

## To run this application
``` go mod tidy ``` To get the dependancy

### Run the server:
``` go run server/main.go ```

### Run the client:
``` go run client/main.go ```
