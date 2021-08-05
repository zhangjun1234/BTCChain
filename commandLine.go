package main

import "fmt"

func (cli *CLI) printBlockChain() {
	bc := cli.bc
	it := bc.NewIterator()
	for {
		block := it.Next()
		fmt.Printf("================================================\n\n\n")
		fmt.Printf("preHash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Transactions[0].TXInputs[0].Sig)

		if len(block.PrevHash) == 0 {
			fmt.Println("printf end !!!!!!!!!!!")
			break
		}
	}
}

func (cli *CLI) AddBlock(data string) {
	//cli.bc.AddBlock(data) TODO
	fmt.Println("add block success!!!!")
}
func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)
	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("address = %s have Balance : %f\n", address, total)
}

func (cli *CLI) Send(from string, to string, amount float64, miner string, data string) {
	//fmt.Printf("from : %s\n",from)
	//	//fmt.Printf("to : %s\n",to)
	//	//fmt.Printf("amount : %f\n",amount)
	//	//fmt.Printf("miner : %s\n",miner)
	//	//fmt.Printf("data : %s\n",data)
	// TODO
	coinbase := NewCoinBaseTX(miner,data)
	tx :=NewTransaction(from,to,amount,cli.bc)
	if tx ==nil{
		fmt.Println("无效的交易")
		return
	}
	cli.bc.AddBlock([]*Transaction{coinbase,tx})
	fmt.Println("转账成功！！！！")
}
