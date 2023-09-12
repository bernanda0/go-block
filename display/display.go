package display

import (
	"blkchn/block"
	"fmt"
	"strings"
)

func StartCard() {
	// Display the blockchain as a card
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║        Blockchain Explorer       ║")
	fmt.Println("╚══════════════════════════════════╝")
}

func EndCard() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║           End of Chain           ║")
	fmt.Println("╚══════════════════════════════════╝")
}

func BlockCard(i int, b *block.Block) {
	fmt.Printf("Block %d\n", i+1)
	fmt.Printf("Timestamp		: %d\n", b.Timestamp)
	fmt.Printf("Image Data		: %s\n", b.ImageData)
	if i != 0 {
		fmt.Printf("Prev Block Hash		: %x\n", b.PrevBlockHash)
	}
	fmt.Printf("Block Hash		: %x\n", b.Hash)
	fmt.Println(strings.Repeat("=", 90))
}
