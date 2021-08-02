package main

//4.intro blockcain
type BlockChain struct {
	//slice Block
	blocks []*Block
}

//5 create BlockChain
func CreateBlockChain() *BlockChain {
	// add gensisBlock to BlockChain
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

//6.GenesisBlock
func GenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

//7 add Block
func (bc *BlockChain) AddBlock(data string) {
	preHash := bc.blocks[len(bc.blocks)-1].Hash
	bc.blocks = append(bc.blocks, NewBlock(data, preHash))
}
