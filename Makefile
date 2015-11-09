default: deps build test

deps:
	go install github.com/hashicorp/terraform

build:
	go build .

test:
	go test -v .
