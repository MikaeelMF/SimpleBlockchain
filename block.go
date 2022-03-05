package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/gob"
)

type Block struct {
	blockHeight   uint64
	Data          []byte
	prevBlockHash [sha512.Size]byte
	Hash          [sha512.Size]byte
	nonce         uint64
}

func NewBlock(data string, prevBlock *Block) *Block {
	newBlock := &Block{blockHeight: prevBlock.blockHeight + 1, Data: []byte(data), prevBlockHash: prevBlock.Hash}
	newPOW := NewProofOfWork(newBlock)
	nonce, hash := newPOW.Run()
	newBlock.nonce = nonce
	newBlock.Hash = hash
	return newBlock
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	_ = encoder.Encode(b)
	return result.Bytes()
}

func Deserialize(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	_ = decoder.Decode(&block)
	return &block
}
