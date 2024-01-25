BINARY_NAME=echo-boilerplate
build:
	go build \
		-ldflags="-s -w" \
		-trimpath \
		-o ./bin/$(BINARY_NAME) \
		./cmd/$(BINARY_NAME)/main.go
	