# Variables
BACKEND_DIR=backend
CONTRACT_JSON=blockchain/artifacts/contracts/ReactiveCounter.sol/ReactiveCounter.json
ABI_FILE=$(BACKEND_DIR)/counter.abi
GO_OUT=$(BACKEND_DIR)/counter.go

.PHONY: all build-contract run-backend install-tools

all: build-contract

# Extract ABI and Generate Go Bindings
build-contract:
	@echo "Extracting ABI from Hardhat artifacts..."
	cat $(CONTRACT_JSON) | jq .abi > $(ABI_FILE)
	@echo "Generating Go bindings..."
	abigen --abi=$(ABI_FILE) --pkg=main --type=ReactiveCounter --out=$(GO_OUT)
	@echo "Done: $(GO_OUT) generated."

# Helper to run the backend
run-backend:
	cd $(BACKEND_DIR) && go mod tidy && go mod vendor && go run main.go counter.go

# Install abigen tool
install-tools:
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest