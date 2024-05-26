LOCAL_BIN:=$(CURDIR)/bin

bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@v1.16.3

.PHONY: docs
docs:
# либа к swag https://github.com/swaggo/swag
	$(LOCAL_BIN)/swag fmt
	$(LOCAL_BIN)/swag init --outputTypes "go,json" --parseInternal -g cmd/main.go