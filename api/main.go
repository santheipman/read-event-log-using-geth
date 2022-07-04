package main

import (
	"context"
	market "example/marketAPI/contracts"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URL                 = "http://127.0.0.1:8545/"
	MARKET_ADDRESS      = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	DATABASE            = "market"
	COLLECTION          = "order"
	MONGODB_URI_ENV_VAR = "MONGODB_URI"
)

func main() {
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err)
	}

	coll := collectionHandler()

	contractAddress := common.HexToAddress(MARKET_ADDRESS)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(20),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error)
	// func (_Market *MarketFilterer) ParseCancelSell(log types.Log) (*MarketCancelSell, error)

	marketFilterer, err := market.NewMarketFilterer(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	logCreateSellOrderSigHash := sigHash("CreateSellOrder(uint256,address,uint256,uint256)")
	logBuySigHash := sigHash("Buy(uint256,address)")
	logCancelSellSigHash := sigHash("CancelSell(uint256)")

	for _, vLog := range logs {
		fmt.Printf("Log Block number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logCreateSellOrderSigHash.Hex():
			fmt.Print("Log name: CreateSellOrder\n")

			createSellOrderEvent, err := marketFilterer.ParseCreateSellOrder(vLog)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("---- orderId: %s\n", createSellOrderEvent.OrderId.String())
			fmt.Printf("---- seller: %s\n", createSellOrderEvent.Seller.Hex())
			fmt.Printf("---- tokenId: %s\n", createSellOrderEvent.TokenId.String())
			fmt.Printf("---- price: %s\n\n", createSellOrderEvent.Price.String())

			err = insertOrderToDB(coll, createSellOrderEvent)
			if err != nil {
				log.Fatal(err)
			}

		case logBuySigHash.Hex():
			fmt.Print("Log name: Buy\n")

			buyEvent, err := marketFilterer.ParseBuy(vLog)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("---- orderId: %s\n\n", buyEvent.OrderId.String())

			// change status to sold, add buyer field
			err = updateOrderAfterBuying(coll, buyEvent)
			if err != nil {
				log.Fatal(err)
			}

		case logCancelSellSigHash.Hex():
			fmt.Print("Log name: CancelSellSigHash\n")

			cancelSellSigHashEvent, err := marketFilterer.ParseCancelSell(vLog)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("---- orderId: %s\n\n", cancelSellSigHashEvent.OrderId.String())

			// change status to canceled
			err = updateOrderAfterCanceling(coll, cancelSellSigHashEvent)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func collectionHandler() *mongo.Collection {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not specified")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database(DATABASE).Collection(COLLECTION)

	return coll
}

func insertOrderToDB(c *mongo.Collection, m *market.MarketCreateSellOrder) error {
	doc := bson.D{
		{Key: "orderId", Value: m.OrderId.String()},
		{Key: "seller", Value: m.Seller.Hex()},
		{Key: "tokenId", Value: m.TokenId.String()},
		{Key: "price", Value: m.Price.String()},
		{Key: "status", Value: "sale"},
	}
	_, err := c.InsertOne(context.TODO(), doc)

	return err
}

func updateOrderAfterBuying(c *mongo.Collection, m *market.MarketBuy) error {
	filter := bson.D{{Key: "orderId", Value: m.OrderId.String()}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "sold"}, {Key: "buyer", Value: m.Buyer.Hex()}}}} // ,
	_, err := c.UpdateOne(context.TODO(), filter, update)

	return err
}

func updateOrderAfterCanceling(c *mongo.Collection, m *market.MarketCancelSell) error {
	filter := bson.D{{Key: "orderId", Value: m.OrderId.String()}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "canceled"}}}}
	_, err := c.UpdateOne(context.TODO(), filter, update)

	return err
}

func sigHash(s string) common.Hash {
	sig := []byte(s)
	return crypto.Keccak256Hash(sig)
}
