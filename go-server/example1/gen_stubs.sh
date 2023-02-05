#!/usr/bin/env bash
protoc --go_out=./server --python_betterproto_out=./client/messaging messages.proto
