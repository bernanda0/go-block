package main

import (
	"blkchn/block"
	"blkchn/display"
)

func main() {
	bc := block.NewBlockChain()

	// Add sample image data
	imageData1 := []byte("Image 1 data")
	bc.AddBlock(imageData1)

	imageData2 := []byte("Image 2 data")
	bc.AddBlock(imageData2)

	display.StartCard()
	for i, block := range bc.Blocks {
		display.BlockCard(i, block)
	}
	display.EndCard()
}
