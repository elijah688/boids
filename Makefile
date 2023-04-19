.PHONY: test run

test:
	go test -race -v -failfast ./...

run:
	go run main.go

