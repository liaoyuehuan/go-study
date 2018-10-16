package main

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
	"fmt"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
}

type BlockChain struct {
	blocks []*Block
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	datas := bytes.Join([][]byte{block.PreBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(datas)
	block.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func (bc *BlockChain) AddBlock(data string) {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data,preBlock.Hash)
	bc.blocks = append(bc.blocks,newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block创始块", []byte{})
}

/*
    用创始块创建一个区块链的函数
*/
func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}} /*我不太理解这种语法,层层深入？
YuloYu 在评论中说:这个地方是，把[]*Block{NewGenesisBlock() 生成的块，把块放到这个数组里面。*/
}

func main() {
	bc := NewBlockchain()

	/*
	   添加记录
	*/
	bc.AddBlock("Send 1 BTC TO L")
	bc.AddBlock("Send 2 BTC to R")

	/*
	  查看我们最终存储结果
	*/
	for _, block := range bc.blocks {
		fmt.Printf("Prev: hash:%x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
