package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlcokchain()

	bc.AddBlock("send 1 btc to zhu")
	bc.AddBlock("send 2 more btc to zhen")

	for _, block := range bc.blcoks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		fmt.Println()
	}
}
