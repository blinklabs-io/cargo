BINARY=cargo
GOMODULE=github.com/cloudstruct/cargo

# Determine root directory
ROOT_DIR=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# Gather all .go files for use in dependencies below
GO_FILES=$(shell find $(ROOT_DIR) -name '*.go')

# Set version strings based on git tag and current ref
GO_LDFLAGS=-ldflags "-X '$(GOMODULE)/version.Version=$(shell git describe --tags --exact-match 2>/dev/null)' -X '$(GOMODULE)/version.CommitHash=$(shell git rev-parse --short HEAD)'"

# Build our program binary
# Depends on GO_FILES to determine when rebuild is needed
$(BINARY): $(GO_FILES)
	# Needed to fetch new dependencies and add them to go.mod
	go mod tidy
	go build \
		$(GO_LDFLAGS) \
		-o $(BINARY) \
		./cmd/$(BINARY)

.PHONY: build image clean

# Alias for building program binary
build: $(BINARY)

# Build docker image
image: build
	docker build -t cloudstruct/$(BINARY) .

clean:
	rm -f $(BINARY)
