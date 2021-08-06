package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "wallet.dat"
type Wallets struct {
	WalletMap map[string]*Wallet
}

func NewWallets() *Wallets{
	var wallets Wallets
	//wallets.WalletMap = make(map[string]*Wallet)
	wallets.loadFile()
	return &wallets
}

func (ws *Wallets)CreateWallet()string{
	wallet:= NewWallet()
	address :=wallet.NewAddress()

	ws.WalletMap[address] = wallet
	ws.saveToFile()
	return address
}

func (ws *Wallets) loadFile(){
	_,err :=os.Stat(walletFile)
	if os.IsNotExist(err){
		ws.WalletMap = make(map[string]*Wallet)
		return
	}

	buffer ,err := ioutil.ReadFile(walletFile)
	if err != nil{
		log.Panic(err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(buffer))
	err  = decoder.Decode(&wallets)
	if err != nil{
		log.Panic(err)
	}
	 ws.WalletMap = wallets.WalletMap
}






func (ws *Wallets) saveToFile() {
	var buffer bytes.Buffer
	gob.Register(elliptic.P256())
	encoder :=gob.NewEncoder(&buffer)
	err := encoder.Encode(&ws)
	if err != nil{
		log.Panic(err)
	}
	ioutil.WriteFile(walletFile,buffer.Bytes(),0600)
}

func (ws *Wallets) GetAllAddresses() []string{
	var addresses []string
	for address :=range ws.WalletMap{
		addresses = append(addresses, address)
	}
	return  addresses
}