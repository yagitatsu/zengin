.PHONY: test
test:
	GO111MODULE=on go test -v ./... -count 1
