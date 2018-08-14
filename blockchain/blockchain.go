package main

type Blockchain struct {
	blcoks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blcoks[len(bc.blcoks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blcoks = append(bc.blcoks, newBlock)
}

func NewBlcokchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
