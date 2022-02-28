package main

import (
	"bytes"
	"crypto/sha512"
	"strconv"
)

type Block struct {
	timestamp     uint64
	Data          []byte
	prevBlockHash [sha512.Size]byte
	Hash          [sha512.Size]byte
}

func (b *Block) SetHash() {
	timestampBytes := []byte(strconv.FormatUint(b.timestamp, 10))
	headers := bytes.Join([][]byte{b.prevBlockHash[:], b.Data, timestampBytes}, []byte{})
	b.Hash = sha512.Sum512(headers)
}

func NewBlock(data string, prevBlockHash [sha512.Size]byte, prevBlock Block) *Block {
	newBlock := &Block{timestamp: prevBlock.timestamp + 1, Data: []byte(data), prevBlockHash: prevBlock.Hash}
	newBlock.SetHash()
	return newBlock
}
