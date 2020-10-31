

regenerate: generate format
	go mod tidy

generate: bin/go-bindata
	-rm -f init/init.go
	-rm -rf components/common/gen/*/
	-rm -f bind/bind.go
	go generate \
		./init \
		./components/common/gen
	go generate \
		./init \
		./bind
	go generate \
		./examples

format: bin/goimports
	bin/goimports -w .

bin/goimports:
	mkdir -p $(@D)
	GOBIN=$$PWD/$(@D) go get golang.org/x/tools/cmd/goimports

bin/go-bindata:
	mkdir -p $(@D)
	GOBIN=$$PWD/$(@D) go get github.com/wzshiming/go-bindata/cmd/go-bindata