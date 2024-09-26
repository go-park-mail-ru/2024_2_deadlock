PKG := `go list -mod=mod -f {{.Dir}} ./...`

init: mod-tidy install-gci install-lint

mod-tidy:
	go mod tidy

install-gci:
	go install github.com/daixiang0/gci@latest

install-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin latest

pre-commit: lint test

fmt:
	go fmt ./...
	gci write -s standard -s default -s "Prefix(github.com/go-park-mail-ru/2024_2_deadlock)" -s blank -s dot $(PKG)

lint: fmt
	golangci-lint run

.PHONY: test
test:
	go test ./... -count=1 -p=1

.PHONY: cover
cover:
	mkdir -p .coverage
	go test ./... -count=1 -p=1 -coverprofile .coverage/cover.out
	go tool cover -html=.coverage/cover.out -o .coverage/cover.html
