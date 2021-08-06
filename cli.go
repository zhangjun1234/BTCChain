package main

import (
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA  "add data to blockchain"
	printChain            "print all data from blockchain"
	getBalance --address ADDRESS "返回指定地址的余额"
	send FROM TO AMOUNT MINER DATA "由from 转账给 to 由 miner 挖矿"
	newWallet  "创建一个新的钱包"
`

func (cli *CLI) Run() {
	//1.得到命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("add block")
		if args[2] == "--data" && len(args) == 4 {
			data := args[3]
			cli.AddBlock(data)
		}
	case "printChain":
		cli.printBlockChain()
	case "getBalance":
		fmt.Println("getBalance")
		if args[2] == "--address" && len(args) == 4 {
			address := args[3]
			cli.GetBalance(address)
		}
	case "send":
		fmt.Println("转账开始")
		if len(args) != 7 {
			fmt.Println("参数不足")
			fmt.Println(Usage)
			return
		}
		from := args[2]
		to := args[3]
		amount, _ := strconv.ParseFloat(os.Args[4], 64)
		miner := args[5]
		data := args[6]
		cli.Send(from, to, amount, miner, data)
	case "newWallet":
		fmt.Println("创建钱包开始........")
		cli.NewWallet()
	default:
		fmt.Println("lalalalalalalalal")
	}
}
