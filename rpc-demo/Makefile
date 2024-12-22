# build protoc --go_out=. --go_opt=paths=source_relative --go_opt=Mtelemetry.proto=github.com/namikmesic/rpc-demo/proto --go-grpc_out=. telemetry.proto


proto:
    protoc --go_out=proto --go_opt=paths=source_relative \
           --go-grpc_out=proto --go-grpc_opt=paths=source_relative \
           proto/telemetry.proto