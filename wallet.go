package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"log"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

func  NewWallet() *Wallet {
	curve := elliptic.P256()
	privateKey ,err:= ecdsa.GenerateKey(curve,rand.Reader)
	if err != nil{
		log.Panic()
	}
	publicKeyOrig := privateKey.PublicKey
	publicKey := append(publicKeyOrig.X.Bytes(),publicKeyOrig.Y.Bytes()...)
	return &Wallet{PrivateKey: privateKey,PublicKey: publicKey}
}

func (w *Wallet) NewAddress()string{
	publicKey := w.PublicKey
	rip160hash := btcutil.Hash160(publicKey)
	//hash := sha256.Sum256(publicKey)
	//riphasher := crypto.RIPEMD160.New()
	//_,err :=riphasher.Write(hash[:])
	//if err != nil{
	//	log.Panic(err)
	//}
	//rip160hash := riphasher.Sum(nil)

	version := byte(00)
	payload := append([]byte{version},rip160hash...)

	hash1 :=sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	checkCode := hash2[:4]

	payload = append(payload,checkCode...)
	address := base58.Encode(payload)
	return address
}