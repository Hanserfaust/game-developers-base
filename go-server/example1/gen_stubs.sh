#!/usr/bin/env bash
protoc --go_out=./server --python_betterproto_out=./client/net ./messages.proto
