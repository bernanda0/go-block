package block

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Timestamp     int64
	ImageData     []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(fmt.Sprint(b.Timestamp))
	headers := append(timestamp, append(b.ImageData, b.PrevBlockHash...)...)
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
