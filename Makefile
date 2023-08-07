BINARY = api

BIN_DIRECTORY = bin


all: $(BINARY)

$(BINARY): clean
	go build -o $(BIN_DIRECTORY)/$(BINARY) cmd/api/*.go

test:
	go test ./... -v

clean:
	rm -f bin/$(BINARY)

.PHONY: all clean
