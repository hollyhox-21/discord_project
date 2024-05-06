.PHONY: local_up
local-up:
	docker-compose up -d --build

.PHONY: local-down
local-down:
	docker-compose down

