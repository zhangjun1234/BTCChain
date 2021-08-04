package main

import (
	"BTClearn/bolt"
	"log"
)

const (
	blockChainDb = "blockChain.db"
	blockBucket  = "blockBucket"
	lastHashKey  = "lastHashKey"
)

//4.intro blockcain
type BlockChain struct {
	//操作数据库的句柄
	db *bolt.DB
	//最后一个区块的hash
	tail []byte
}

//5 create BlockChain
func CreateBlockChain() *BlockChain {
	var db *bolt.DB
	var lastHash []byte
	/*1. 打开数据库(没有的话就创建)
	2. 找到抽屉（bucket），如果找到，就返回bucket，如果没有找到，我们要创建bucket，通过名字创建
		a. 找到了
			1. 通过"last"这个key找到我们最好⼀个区块的哈希。
		b. 没找到创建
			1. 创建bucket，通过名字
			2. 添加创世块数据 3. 更新"last"这个key的value（创世块的哈希值） */
	db, err := bolt.Open(blockChainDb, 0600, nil)
	//defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败！")
	}

	db.Update(func(tx *bolt.Tx) error {
		//创建或者找到bucket
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket失败")
			}
			genesisBlock := GenesisBlock()
			bucket.Put(genesisBlock.Hash, []byte(genesisBlock.Serialize()))
			bucket.Put([]byte(lastHashKey), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte(lastHashKey))
		}
		return nil
	})
	return &BlockChain{db: db, tail: lastHash}
}

//6.GenesisBlock
func GenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

//7.add Block
func (bc *BlockChain) AddBlock(data string) {
	db := bc.db
	lashHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {
		//创建或者找到bucket
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("打开bucket失败")
		} else {
			//创建新区块
			block := NewBlock(data, lashHash)
			//将区块加入数据库
			bucket.Put(block.Hash, block.Serialize())
			bucket.Put([]byte(lastHashKey),block.Hash)
			//更新内存中的区块尾巴
			bc.tail = block.Hash
		}
		return nil
	})

}
