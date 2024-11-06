package root

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//go:generate sh -c "protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative protos/*/*.proto"
