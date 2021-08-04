package main

import (
	"BTClearn/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain)NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

func (it *BlockChainIterator) Next() *Block{
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			log.Panic("bucket is nil")
		}
		blockTmp :=bucket.Get(it.currentHashPointer)
		block = Deserialize(blockTmp)
		it.currentHashPointer=block.PrevHash
		return nil
	})
	return &block
}

