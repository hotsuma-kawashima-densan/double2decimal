GO := go
GO_BUILD := $(GO) build
GO_ENV := GOOS=linux GOARCH=arm64 CGO_ENABLED=0
GO_FLAGS := -ldflags="-s -w" -tags lambda.norpc -trimpath
BIN_ROOT := ./release
TARGETS := \
	batch
.PHONY: gen-db
gen-db:
	@go run tools/gorm-gen/main.go
