package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reward = 12.5

type Transaction struct {
	TXID      []byte
	TXInputs  []TXInput
	TXOutputs []TXOutput
}

type TXInput struct {
	TXid  []byte
	Index int64
	Sig   string
}

type TXOutput struct {
	Value      float64
	PubKeyHash string
}

func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

func NewCoinBaseTX(address string, data string) *Transaction {
	input := TXInput{TXid: []byte{}, Index: -1, Sig: data}
	output := TXOutput{Value: reward, PubKeyHash: address}
	tx := Transaction{TXID: []byte{}, TXInputs: []TXInput{input}, TXOutputs: []TXOutput{output}}
	tx.SetHash()
	return &tx
}

func NewTransaction(from string, to string, amount float64, bc *BlockChain) *Transaction {
	utxos, resVal := bc.FindNeedUtxos(from, amount)
	if resVal < amount {
		fmt.Println("not enough balance")
		return nil
	}
	var inputs []TXInput
	var outputs []TXOutput

	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{TXid: []byte(id), Index: int64(i), Sig: from}
			inputs = append(inputs, input)
		}
	}

	output := TXOutput{Value: amount, PubKeyHash: to}
	outputs = append(outputs, output)
	if resVal > amount {
		output := TXOutput{Value: resVal - amount, PubKeyHash: from}
		outputs = append(outputs, output)
	}
	tx := Transaction{TXID: []byte{}, TXInputs: inputs, TXOutputs: outputs}
	tx.SetHash()
	return &tx
}

func (tx *Transaction) IsCoinBase() bool {
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 || tx.TXInputs[0].Index == -1 {
		return true
	}
	return false
}


