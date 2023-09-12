package block

import "time"

func NewBlock(prevBlock *Block, imageData []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		ImageData:     imageData,
		PrevBlockHash: prevBlock.Hash,
	}
	block.SetHash()
	return block
}

func NewBlockChain() *Blockchain {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		ImageData:     []byte("START GO-BLOCK"),
		PrevBlockHash: []byte{},
	}
	block.SetHash()
	return &Blockchain{[]*Block{block}}
}
