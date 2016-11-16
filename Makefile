GIT_VER := $(shell git describe --tags)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)

.PHONY: test clean

all: test install clean

install:
	cd cmd/hccc && go build -ldflags "-X main.version=${GIT_VER} -X main.buildDate=${DATE}" && install hccc ${GOPATH}/bin

clean:
	rm -f cmd/hccc/hccc

test:
	@echo ${GOPATH}
	cd hccc && go test
