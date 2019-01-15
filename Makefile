.PHONY: test

test:
	go test ./pkg/... -cover -count=1
