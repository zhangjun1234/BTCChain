package main

import (
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}
const Usage = `
	addBlock --data DATA  "add data to blockchain"
	printChain            "print all data from blockchain"
`

func (cli *CLI)Run()  {
	//1.得到命令
	args := os.Args
	if len(args) < 2{
		fmt.Printf(Usage)
		return
	}
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("add block")
		if args[2]=="--data" && len(args)==4{
			data := args[3]
			cli.AddBlock(data)
		}
	case "printChain":
		cli.printBlockChain()
	default:
		fmt.Println("lalalalalalalalal")
	}
}