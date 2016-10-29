install:
	cd cmd/hccc && go build && install hccc ${GOPATH}/bin

clean:
	rm -f cmd/hccc/hccc

test:
	@echo ${GOPATH}
	cd hccc && go test
