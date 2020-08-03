.PHONY: build
build:
	go build ./cmd/mapiter/main.go

.PHONY: install
install:
	go install ./cmd/...

.PHONY: test
test:
	go test -failfast

.PHONY: clean
clean:
	rm -rf main
