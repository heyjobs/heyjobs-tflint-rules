PLUGIN_NAME = heyjobs-tflint-rules
VERSION = 0.1.0

TARGETS = \
	darwin_amd64 \
	darwin_arm64 \
	linux_amd64 \
	linux_arm64

.PHONY: default build test clean dist install release

default: build

test:
	go test ./...

clean:
	rm -rf dist
	rm -f $(PLUGIN_NAME)

dist: clean
	mkdir -p dist
	for target in $(TARGETS); do \
		GOOS=$${target%_*} GOARCH=$${target#*_} go build -o dist/$(PLUGIN_NAME); \
		cd dist && zip $(PLUGIN_NAME)_$${target}.zip $(PLUGIN_NAME) && rm $(PLUGIN_NAME) && cd ..; \
	done
	cd dist && sha256sum *.zip > checksums.txt
	cd dist && gpg --detach-sign --armor --output checksums.txt.sig checksums.txt

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./$(PLUGIN_NAME) ~/.tflint.d/plugins

release: dist
	gh release create v$(VERSION) \
		--title "Release v$(VERSION)" \
		--notes "Release v$(VERSION)" \
		dist/*.zip dist/checksums.txt dist/checksums.txt.sig
