mod:
	go mod download
	go mod vendor
.PHONY: mod

test:
	go test -v `go list ./... | grep -v mocks | grep -v test` -cover
.PHONY: test

mock:
	cd ./somepkg && mockery --all
.PHONY: mock

lint:
	golangci-lint run
.PHONY: lint
