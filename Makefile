export GO111MODULE=on

build:
	go build

build-linux:
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

run: clean
	go run votingpower.go

clean:
	rm -vf ./votingpower-tracker

.PHONY: build build-linux clean run