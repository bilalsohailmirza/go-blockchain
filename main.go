package main

import (
	"fmt"
	"strconv"

	"github.com/bilalsohailmirza/golang-blockchain/blockchain"
)

func main() {

	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for i, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println(i)
		pow := blockchain.NewProof(block)
		fmt.Printf("ProofOfWork: %s\n", strconv.FormatBool(pow.Validate()))
	}
}