package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"geth-demo/cont"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math"
	"math/big"
	"strings"
)

func main() {

}

// 理论上来讲，这个函数应该只被创建一次即可
// 创建一个账户
func CreateWallet() (key, addr string) {
	ks := keystore.NewKeyStore("/Users/zhuzhenfeng/Documents/github/gowork/src/geth-demo/",
		keystore.StandardScryptN, keystore.StandardScryptP)
	account, _ := ks.NewAccount("password")
	key_json, err := ks.Export(account, "password", "password")
	if err != nil {
		log.Fatalln("导出账户错误: ", err)
		panic(err)
	}
	key = string(key_json)
	addr = account.Address.Hex()
	return
}

func connectRPC() (*ethclient.Client, error) {
	// 连接测试链的节点
	//rpcClient, err := rpc.Dial("https://rinkeby.infura.io/v3/6c81fb1b66804f0698d49f2ec242afc9")
	rpcClient, err := rpc.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	conn := ethclient.NewClient(rpcClient)
	return conn, nil
}

func TransferToken() {
	key, to_address := CreateWallet()

	client, err := connectRPC()
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), "password")
	if err != nil {
		log.Fatalln("读取keystore失败", err)
		panic(err)
	}

	token, err := cont.NewToken(common.HexToAddress("0x75a26aaaecda412bd914e8fbfaed586a467fa8b5"), client)
	if err != nil {
		log.Fatalln("获取token失败", err)
		panic(err)
	}

	balance, err := token.BalanceOf(nil, common.HexToAddress(to_address))
	if err != nil {
		log.Fatalln("token balance of", err)
	}

	log.Println("to address balance: ", balance)

	amount := big.NewFloat(10.00)
	//这是处理位数的代码段
	tenDecimal := big.NewFloat(math.Pow(10, 18))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})

	tx, err := token.Transfer(auth, common.HexToAddress(to_address), convertAmount)
	if nil != err {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Printf("result: %v\n", tx)
}

func GetBalance(address string) (float64, error) {
	client, err := connectRPC()
	if err != nil {
		log.Fatalln("err: ", err)
		panic(err)
	}

	balance, err := client.BalanceAt(context.TODO(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalln(balance)
		return 0, err
	}
	balanceV := float64(balance.Int64()) * math.Pow(10, -18)
	return balanceV, nil
}

func CreateAccount() (string, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln(err)
		return "", nil
	}

	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	log.Println("address: ", address)

	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Println("privateKey: ", privateKey)
	return address, nil
}
