

regenerate: generate format
	go mod tidy

generate: bin/go-bindata
	-rm -f init/init_gen.go
	go generate \
		./hack/gen_common \
		./components/common/gen \
		./init \
		./bind \
		./examples

format: bin/goimports
	bin/goimports -w .

bin/goimports:
	mkdir -p $(@D)
	GOBIN=$$PWD/$(@D) go get golang.org/x/tools/cmd/goimports

bin/go-bindata:
	mkdir -p $(@D)
	GOBIN=$$PWD/$(@D) go get github.com/wzshiming/go-bindata/cmd/go-bindata