CMD := semtag

bin/$(CMD): go.* *.go
	make test
	go vet .
	go fmt .
	mkdir -p bin
	go build -o bin/$(CMD) .

.PHONY: test
test:
	go test -cover ./...
	which gocyclo && ./scripts/test_cyclomatic_complexity.sh

.PHONY: install
install: go.* *.go
	go install .

.PHONY: setup-tools
setup-tools:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.6.0
