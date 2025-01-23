default: build

test:
	go test ./...

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./heyjobs-tflint-rules ~/.tflint.d/plugins
