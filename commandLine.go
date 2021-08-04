package main

import "fmt"

func (cli *CLI)printBlockChain() {
	bc := cli.bc
	it := bc.NewIterator()
	for {
		block := it.Next()
		fmt.Printf("================================================\n\n\n")
		fmt.Printf("preHash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)

		if len(block.PrevHash) == 0 {
			fmt.Println("printf end !!!!!!!!!!!")
			break
		}
	}
}

func (cli *CLI)AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("add block success!!!!")
}
