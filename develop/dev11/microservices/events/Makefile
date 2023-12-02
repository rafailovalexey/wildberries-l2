# Variables

MOCKS_OUTPUT_DIRECTORY = mocks

MOCKS_FILES = \
	internal/repository/repository.go

# Mocks

mocks-generate:
	bin/mocks-generate.sh $(MOCKS_OUTPUT_DIRECTORY) $(MOCKS_FILES)

# Generate

generate:
	make mocks-generate

# Download

download:
	go mod download

# Build

build:
	go build -o build/main main.go

# Tests

tests:
	go test -v ./...

.PHONY: mocks-generate, generate, download, build, tests