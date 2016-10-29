GIT_VER := $(shell git describe --tags)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)

install:
	cd cmd && go build -ldflags "-X main.version=${GIT_VER} -X main.buildDate=${DATE}" && install hccc ${GOPATH}/bin

