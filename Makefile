APP_GO_FILES := $(shell find . -name '*.go')

all: linux.zip

linux.zip: build/linux_amd64/can LICENSE
	zip -r linux build/linux_amd64/can LICENSE

build/linux_amd64/can: $(APP_GO_FILES)
	go build -o ./build/linux_amd64/ ./cmd/...