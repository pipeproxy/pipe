

generate:
	go generate ./...
	make format

format:
	goimports -w .