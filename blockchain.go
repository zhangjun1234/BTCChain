package main

import (
	"BTClearn/bolt"
	"fmt"
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
func CreateBlockChain(address string) *BlockChain {
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
			genesisBlock := GenesisBlock(address)
			fmt.Printf("genesis block = %s\n", genesisBlock)
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
func GenesisBlock(address string) *Block {
	coinbase := NewCoinBaseTX(address, "GenesisBlock")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//7.add Block
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	db := bc.db
	lashHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {
		//创建或者找到bucket
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("打开bucket失败")
		} else {
			//创建新区块
			block := NewBlock(txs, lashHash)
			//将区块加入数据库
			bucket.Put(block.Hash, block.Serialize())
			bucket.Put([]byte(lastHashKey), block.Hash)
			//更新内存中的区块尾巴
			bc.tail = block.Hash
		}
		return nil
	})

}

//找到指定地址的所有的UTXO
func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	it := bc.NewIterator()
	spentoutputs := make(map[string][]int64)
	for {
		block := it.Next()
		for _, tx := range block.Transactions {
			fmt.Printf("txid : %x\n", tx.TXID)

			for i, output := range tx.TXOutputs {
				fmt.Printf("current index : %d\n", i)
				flag := true
				if spentoutputs[string(tx.TXID)] != nil {
					for _, j := range spentoutputs[string(tx.TXID)] {
						if int64(i) == j {
							flag = false
							break
						}
					}
				}

				if (output.PubKeyHash == address) && (flag == true) {
					UTXO = append(UTXO, output)
				}
			}

			if !tx.IsCoinBase() {
				for _, input := range tx.TXInputs {
					if input.Sig == address {
						indexArray := spentoutputs[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
						spentoutputs[string(input.TXid)] = indexArray
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}
	return UTXO
}

func (bc *BlockChain) FindNeedUtxos(from string, amount float64) (map[string][]uint64, float64) {
	utxos := make(map[string][]uint64)
	var cal float64

	it := bc.NewIterator()
	spentoutputs := make(map[string][]int64)
	for {
		block := it.Next()
		for _, tx := range block.Transactions {
			fmt.Printf("txid : %x\n", tx.TXID)

			for i, output := range tx.TXOutputs {
				fmt.Printf("current index : %d\n", i)
				flag := true
				if spentoutputs[string(tx.TXID)] != nil {
					for _, j := range spentoutputs[string(tx.TXID)] {
						if int64(i) == j {
							flag = false
							break
						}
					}
				}

				if (output.PubKeyHash == from) && (flag == true) {
					if cal < amount {
						utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(i))
						cal += output.Value
						if cal >= amount {
							fmt.Println("找到了满足的金额")
							return utxos, cal
						}
					} else {
						fmt.Println("不满足转账金额")
					}
				}
			}

			if !tx.IsCoinBase() {
				for _, input := range tx.TXInputs {
					if input.Sig == from {
						indexArray := spentoutputs[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
						spentoutputs[string(input.TXid)] = indexArray
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}

	return nil, 0
}

func (bc *BlockChain)GetUTXOTransaction(address string)[]*Transaction{
	var txs []*Transaction

	it := bc.NewIterator()
	spentoutputs := make(map[string][]int64)
	for {
		block := it.Next()
		for _, tx := range block.Transactions {
			fmt.Printf("txid : %x\n", tx.TXID)

			for i, output := range tx.TXOutputs {
				fmt.Printf("current index : %d\n", i)
				flag := true
				if spentoutputs[string(tx.TXID)] != nil {
					for _, j := range spentoutputs[string(tx.TXID)] {
						if int64(i) == j {
							flag = false
							break
						}
					}
				}

				if (output.PubKeyHash == address) && (flag == true) {
					txs = append(txs, tx)
				}
			}

			if !tx.IsCoinBase() {
				for _, input := range tx.TXInputs {
					if input.Sig == address {
						indexArray := spentoutputs[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
						spentoutputs[string(input.TXid)] = indexArray
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}
	return txs
}
