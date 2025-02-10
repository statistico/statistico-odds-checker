# Makefile
.PHONY: all

NETWORK_NAME=statistico_internal

check-network:
	@echo "Checking if network '$(NETWORK_NAME)' exists..."
	@if [ -z "$$(docker network ls --filter name=^$(NETWORK_NAME)$$ --format="{{ .Name }}")" ]; then \
  		echo "Creating network '$(NETWORK_NAME)'..."; \
    	docker network create $(NETWORK_NAME); \
    else \
		echo "Network '$(NETWORK_NAME)' already exists."; \
	fi

docker-build: check-network
	docker compose -f docker-compose.build.yml up -d --build

docker-up: check-network
	docker compose -f docker-compose.build.yml -f docker-compose.dev.yml up -d --build --force-recreate

docker-down:
	docker compose -f docker-compose.build.yml -f docker-compose.dev.yml down -v

docker-run-console:
	docker compose -f docker-compose.build.yml -f docker-compose.dev.yml run console ./console $(args)

test:
	docker compose -f docker-compose.build.yml run test gotestsum -f short-verbose $(args)

docker-logs:
	docker compose -f docker-compose.build.yml -f docker-compose.dev.yml logs -f $(service)
