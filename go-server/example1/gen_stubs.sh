#!/usr/bin/env bash
protoc --go_out=. --go_opt=paths=source_relative --python_betterproto_out=./client/messaging ./server/messages.proto
