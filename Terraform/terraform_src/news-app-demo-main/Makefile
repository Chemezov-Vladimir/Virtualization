SERVICE_NAME := news-demo
CURRENT_DIR = $(shell pwd)

GOOS?=darwin
GOARCH?=amd64

RELEASE := $(shell git tag -l | tail -1 | grep -E "v.+"|| echo devel)
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	cd ${CURRENT_DIR}/app && GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build -mod=readonly\
		-ldflags "-s -w -X 'news-demo/internal/version.Version=${RELEASE}' \
		-X 'news-demo/internal/version.Commit=${COMMIT}' \
		-X 'news-demo/internal/version.BuildTime=${BUILD_TIME}'" \
		-o ${CURRENT_DIR}/.bin/${SERVICE_NAME} ./cmd/${SERVICE_NAME}/*.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build