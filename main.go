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

// TODO
func (b *Block) GetDifficulty() int {
	return 32
}

// NewBlock parses the Json data of a new block and creates an instance
// of a Block struct.
func NewBlock(data string) *Block {
	block := new(Block)
	err := json.Unmarshal([]byte(data), &block)
	if err != nil {
		panic(err)
	}
	block.Hash = block.ToHash()
	return block
}

func main() {
	b := Block{PreviousHash: "somedatah3re", Index: 2}
	val := b.ToHash()
	fmt.Println(val)
}
