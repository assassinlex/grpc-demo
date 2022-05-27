# pb.go
protoc -I=. --go_out=plugins=grpc,paths=source_relative:../pb auth.proto
# pb.gateway.go
protoc -I=. --grpc-gateway_out=grpc_api_configuration=auth.yaml,paths=source_relative:../pb auth.proto