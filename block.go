package main

import (
	"bytes"
	"crypto/sha512"
	"strconv"
)

type Block struct {
	blockHeight   uint64
	Data          []byte
	prevBlockHash [sha512.Size]byte
	Hash          [sha512.Size]byte
}

func (b *Block) SetHash() {
	blockHeightBytes := []byte(strconv.FormatUint(b.blockHeight, 10))
	headers := bytes.Join([][]byte{b.prevBlockHash[:], b.Data, blockHeightBytes}, []byte{})
	b.Hash = sha512.Sum512(headers)
}

func NewBlock(data string, prevBlockHash [sha512.Size]byte, prevBlock Block) *Block {
	newBlock := &Block{blockHeight: prevBlock.blockHeight + 1, Data: []byte(data), prevBlockHash: prevBlock.Hash}
	newBlock.SetHash()
	return newBlock
}
