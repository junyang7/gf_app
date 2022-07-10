#!/usr/bin/env bash

protoc --go-grpc_out=. --go_out=. service.proto
