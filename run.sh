#!/bin/bash

cleanup() {
    echo "Cleaning up..."
    kill %1
    kill %2
    kill %3
}

trap cleanup SIGINT

go run cmd/users/main.go &
go run cmd/auth/main.go &
go run cmd/app/main.go &
go run cmd/s3/main.go &

wait