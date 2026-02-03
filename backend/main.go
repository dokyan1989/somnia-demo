package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv" // You'll need: go get github.com/joho/godotenv
)

func main() {
	// 1. Load .env to get configuration (Optional, if you want to keep address there)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Connect to Somnia Shannon Testnet
	client, err := ethclient.Dial(os.Getenv("SOMNIA_WEBSOCKET_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to Somnia: %v", err)
	}
	fmt.Println("✅ Connected to Somnia Network")

	// 3. Bind to the Deployed Contract
	// REPLACE THIS with the address you got from the deploy script!
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	instance, err := NewReactiveCounter(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to bind to contract: %v", err)
	}

	// 4. Setup Event Watcher
	// We make a channel to receive the events
	sink := make(chan *ReactiveCounterIncrement)

	// Watch for the event
	// nil, nil means "listen to all addresses" and "all counts" (no filtering)
	sub, err := instance.WatchIncrement(&bind.WatchOpts{}, sink)
	if err != nil {
		log.Fatalf("Failed to subscribe to events: %v", err)
	}
	fmt.Println("👀 Listening for Counter updates...")

	// 5. The Event Loop
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case event := <-sink:
			// "event" is now a strongly typed struct!
			fmt.Printf("\n🔔 EVENT RECEIVED!\n")
			fmt.Printf("   By: %s\n", event.By.String())
			// fmt.Printf("   New Count: %s\n", event.NewCount.String())
			// fmt.Printf("   Updated By: %s\n", event.UpdatedBy.Hex())
			// fmt.Printf("   Timestamp: %s\n", event.Timestamp.String())
		}
	}
}
