package block

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(imageData []byte) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock, imageData)
	bc.Blocks = append(bc.Blocks, newBlock)
}
