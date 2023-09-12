package main

import (
	"blkchn/block"
	"fmt"
)

func main() {
	bc := block.NewBlockChain()

	// Add sample image data
	imageData1 := []byte("Image 1 data")
	bc.AddBlock(imageData1)

	imageData2 := []byte("Image 2 data")
	bc.AddBlock(imageData2)

	// Print the blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("ImageData: %s\n", block.ImageData)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
