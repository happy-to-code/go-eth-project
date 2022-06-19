package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	conn, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
	}
	defer conn.Close()

	contractAddr := "0xC8E0Ffb5a3a18e710D5DBBB9Af160b614C4D6A6A"
	instance, err := NewSimpleStorage(common.HexToAddress(contractAddr), conn)
	if err != nil {
		panic(err)
	}

	message1, err := instance.GetMessage(nil)
	if err != nil {
		fmt.Println("获取信息失败，err:", err)
		return
	}
	fmt.Println("第一次获取信息：", message1)

	// 调用合约setMessage方法 改变值
	privateKey, err := crypto.HexToECDSA("879bbacd37f915823c97729c082d19ac8e7325e2f58453eb68ccc677c8ec592c")
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := conn.PendingNonceAt(context.Background(), common.HexToAddress("0x3b3e7fa456f1F9D55ea796701a32e4c94E17a90a"))
	if err != nil {
		log.Fatal("PendingNonceAt:", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	// auth.GasPrice = gasPrice
	tx, err := instance.SetMessage(auth, "你好，yida")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Errorf("tx:%+v\n", tx)

	message2, err := instance.GetMessage(nil)
	if err != nil {
		fmt.Println("2获取信息失败，err:", err)
		return
	}
	fmt.Println("第二次获取信息：", message2)

}

// 879bbacd37f915823c97729c082d19ac8e7325e2f58453eb68ccc677c8ec592c
func test() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("879bbacd37f915823c97729c082d19ac8e7325e2f58453eb68ccc677c8ec592c")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress:", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xC8E0Ffb5a3a18e710D5DBBB9Af160b614C4D6A6A")
	instance, err := NewSimpleStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SetMessage(auth, "yida")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	result, err := instance.GetMessage(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:]))

}
