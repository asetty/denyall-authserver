GOOS ?= linux
GOARCH ?= amd64

build:
	@echo "Building to ./bin"
	# maybe add other flags like -ldflags
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/denyall-authserver

clean:
	@echo "Cleaning"
	rm -rf ./bin


.PHONY: build clean
