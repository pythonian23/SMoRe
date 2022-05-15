install: format
	go install ./cmd/smore
format:
	gofmt -w .
	goimports -w .
	go mod tidy
