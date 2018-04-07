package main

import (
	"crypto/sha256"
	"fmt"
)

type transaction string

// Block contains data of the blockchain.
type Block struct {
	index        int
	previousHash string
	timestamp    int
	nonce        int
	transactions []transaction
	hash         string
}

func (b *Block) ToHash() string {
	sum := sha256.Sum256([]byte("hello world\n"))
	return fmt.Sprintf("%x", sum)
}

func main() {
	b := new(Block)
	val := b.ToHash()
	fmt.Println(val)
}
