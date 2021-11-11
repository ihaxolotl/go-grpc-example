#!/bin/bash

protoc --proto_path=proto \
    --go_out=internal/proto --go_opt=paths=source_relative \
    --go-grpc_out=internal/proto --go-grpc_opt=paths=source_relative \
    user.proto
