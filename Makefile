PACKAGE=commentService
VERSION_PACKAGE=$(PACKAGE)/pkg/build

include mk/git.mk
include mk/golang.mk

COMMENT_DIR := app/grpc/comment
BIN_DIR := ../../bin

build: show.go tidy build.comment

.PHONY: show.go pre lint tidy generate grpc_pre build.comment

show.go:
	@echo "=====> $(GOOS) $(GO_BUILD_FLAGS)"
build.comment:
	$(GO) build -C $(COMMENT_DIR)  $(GO_BUILD_FLAGS) -o $(BIN_DIR)/comment$(GO_OUT_EXT)
pre:
	curl -L https://git.io/vp6lP | sh
	$(GO) install golang.org/x/lint/golint@latest
	$(GO) install github.com/gordonklaus/ineffassign@latest
grpc_pre:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get -d github.com/envoyproxy/protoc-gen-validate
lint:
	golangci-lint run ./...
	#golint ./...
	#ineffassign ./...
	#go vet ./...
tidy:
	@echo "tidy..."
	$(GO) mod tidy

generate: generate.pre
	$(GO) generate ./...

test:
	$(GO) test -cover ./...


.PHONY: generate.pre install.oapi-codegen
generate.pre: install.oapi-codegen install.stringer

install.oapi-codegen:
ifneq ($(shell oapi-codegen -version >/dev/null 2>&1 && echo 0 || echo 1), 0)
	$(GO) install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
endif
install.stringer:
	$(GO) install golang.org/x/tools/cmd/stringer

