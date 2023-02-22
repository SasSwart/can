APP_GO_FILES := $(shell find . -name '*.go')

all: linux_amd64.zip

linux_amd64.zip: build/linux_amd64 LICENSE
	zip -r linux_amd64 build/linux_amd64 LICENSE config.yaml

build/linux_amd64: build/linux_amd64/can build/linux_amd64/templates

build/linux_amd64/can: $(APP_GO_FILES)
	go build -o ./build/linux_amd64/ ./cmd/...

build/linux_amd64/templates:
	cp -r templates build/linux_amd64/templates

clean:
	rm -rf build linux_amd64.zip

.PHONY: test test_coverage
test:
	go test ./openapi/...

test_coverage:
	go test ./openapi -coverprofile=coverage.out
	go tool cover -html=coverage.out
