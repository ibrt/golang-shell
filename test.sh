#!/usr/bin/env bash

set -ex
cd "$(dirname "${BASH_SOURCE[0]}")"

go mod tidy
go generate ./...
go build -v ./...
diff -u <(echo -n) <(gofmt -d ./)
go run golang.org/x/lint/golint@latest -set_exit_status ./...
go vet ./...
go run honnef.co/go/tools/cmd/staticcheck@release.2022.1 ./...
go test -v -race -failfast -shuffle=on -covermode=atomic -coverprofile=coverage.txt ./...