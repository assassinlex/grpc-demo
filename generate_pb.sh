function pbGenerator {
    DOMAIN=$1
    PROTO_PATH=./${DOMAIN}/api/v1/proto
    GO_OUT_PATH=./${DOMAIN}/api/v1/pb
    mkdir -p $GO_OUT_PATH

    protoc -I=$PROTO_PATH --go_out=plugins=grpc,paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
    protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/${DOMAIN}.yaml:$GO_OUT_PATH ${DOMAIN}.proto
}

pbGenerator auth
pbGenerator rental

