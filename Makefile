.DEFAULT_GOAL := local-dev-all

.PHONY: go-doc
go-doc:
	go install golang.org/x/tools/cmd/godoc@latest
	godoc -http :8081

.PHONY: go-fmt
go-fmt:
	gofmt -d -s -w .

.PHONY: go-lint
go-lint:
	golangci-lint run ./... examples/multiple-jobs

.PHONY: go-test
go-test:
	go test ./...

.PHONY: local-dev-all
local-dev-all: go-fmt go-test go-lint
