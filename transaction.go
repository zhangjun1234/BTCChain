package main

type Transaction struct {
	TXID      []byte
	TXInputs  []TXInput
	TXOutputs []TXOutput
}

type TXInput struct {
	TXid  []byte
	Index uint64
	Sig   string
}

type TXOutput struct {
	Value      float64
	PubKeyHash string
}
