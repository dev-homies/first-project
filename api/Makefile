build:
	go build -o .bin/main main.go

run:
	go run main.go

watch:

format:
	gofmt -s -l -w .

format-ci:
	go mod tidy && git diff --exit-code go.mod

lint:
	go mod tidy

lint-ci:
	test -z $(gofmt -l .)
