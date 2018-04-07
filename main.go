package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type transaction string

// Block contains data of the blockchain.
type Block struct {
	index        uint16
	timestamp    uint16
	nonce        uint16
	transactions []transaction
	previousHash string
	hash         string
}

func (b *Block) ToHash() string {
	uintBytes := make([]byte, 6)
	binary.LittleEndian.PutUint16(uintBytes[0:], b.index)
	binary.LittleEndian.PutUint16(uintBytes[2:], b.timestamp)
	binary.LittleEndian.PutUint16(uintBytes[4:], b.nonce)
	// for _, x := range b.transactions {

	// }
	var buf bytes.Buffer
	buf.Write(uintBytes)
	buf.Write([]byte(b.previousHash))
	sum := sha256.Sum256(buf.Bytes())
	return fmt.Sprintf("%x", sum)
}

func main() {
	b := Block{previousHash: "somedatah3re", index: 3}
	val := b.ToHash()
	fmt.Println(val)
}
