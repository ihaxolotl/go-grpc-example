#!/bin/bash

protoc --proto_path=proto \
    --go_out=internal/transport --go_opt=paths=source_relative \
    --go-grpc_out=internal/transport --go-grpc_opt=paths=source_relative \
    user.proto
