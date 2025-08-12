package accounts

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	godotenv "github.com/joho/godotenv"
)

func AccountBalance() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("CREDO_API_KEY")
	if apiKey == "" {
		log.Fatal("CREDO_API_KEY not set in .env")
	}

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the Ethereum network")

	account := common.HexToAddress("0xE276Bc378A527A8792B353cdCA5b5E53263DfB9e")
	fmt.Println("Account:", account)

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, _ := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) 
}
