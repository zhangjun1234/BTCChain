package main

import "fmt"

func main() {
	bc := CreateBlockChain()
	bc.AddBlock("hello")
	bc.AddBlock("no no no")
	for i, block := range bc.blocks {
		fmt.Printf("heigh: %d\n", i)
		fmt.Printf("preHash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
	}
}
