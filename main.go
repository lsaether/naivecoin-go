package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type transaction string

// Block contains data of the blockchain.
type Block struct {
	Index        uint16
	Timestamp    uint16
	Nonce        uint16
	Transactions []transaction
	PreviousHash string
	Hash         string
}

func (b Block) ToHash() string {
	blockMarshalled, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(blockMarshalled)
	return fmt.Sprintf("%x", sum)
}

func (b *Block) GetDifficulty() int {
	return 32
}

func main() {
	b := Block{PreviousHash: "somedatah3re", Index: 2}
	val := b.ToHash()
	fmt.Println(val)
}
