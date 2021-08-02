package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

//0.create block struct
type Block struct {
	//version
	Version uint64
	//pre block hash
	PrevHash []byte
	//MerkelRoot
	MerkelRoot []byte
	//timeStamp
	TimeStamp uint64
	//Difficulty
	Difficulty uint64
	//nonce
	Nonce uint64
	//cur block hash
	Hash []byte
	//data
	Data []byte
}

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//2.create block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{}, //empty,later compu //TODO
		Data:       []byte(data),
	}

	block.SetHash()
	return &block
}

//3.create hash
func (block *Block) SetHash() {
	//1.add data
	//blockinfo = append(blockinfo, Uint64ToByte(block.Version)...)
	//blockinfo = append(blockinfo, block.PrevHash...)
	//blockinfo = append(blockinfo, block.MerkelRoot...)
	//blockinfo = append(blockinfo, Uint64ToByte(block.TimeStamp)...)
	//blockinfo = append(blockinfo, Uint64ToByte(block.Difficulty)...)
	//blockinfo = append(blockinfo, Uint64ToByte(block.Nonce)...)
	//blockinfo = append(blockinfo, block.Data...)
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	blockinfo := bytes.Join(tmp, []byte{})

	//2.SHA256
	hash := sha256.Sum256(blockinfo)
	block.Hash = hash[:]
}
