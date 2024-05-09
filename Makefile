LOCAL_BIN:=$(CURDIR)/bin

bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/go-swagger/go-swagger/cmd/swagger@latest

gen-docs:
	find . -name "*swagger.json" | xargs $(LOCAL_BIN)/swagger mixin -o total_docs/total.swagger.json


.PHONY: local_up
local-up:
	docker-compose up -d --build

.PHONY: local-down
local-down:
	docker-compose down
