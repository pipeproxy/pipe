

generate:
	go generate ./...
	make format
	go mod tidy

format:
	goimports -w .