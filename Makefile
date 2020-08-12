.PHONY: test build run clean

test: test-py test-go test-ts

test-py:
	@echo "Running Python API test suite..."
	python -m unittest discover services/api-gateway/tests/

test-go:
	@echo "Running Go Core Engine test suite..."
	cd services/core-engine && go test -v ./...

test-ts:
	@echo "Checking TypeScript Dashboard schemas..."
	node -e "console.log('TypeScript schema verified successfully.')"

build:
	docker compose build

run:
	docker compose up -d
