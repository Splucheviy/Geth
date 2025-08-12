package client

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// GetDefaultRPCURL возвращает URL по умолчанию для разработки
func GetDefaultRPCURL() string {
	return "https://rpc.sepolia.org"
}

func Client() {
	// Используем URL из конфигурации
	client, err := ethclient.Dial(GetDefaultRPCURL())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the Ethereum Sepolia testnet")
	_ = client
}
