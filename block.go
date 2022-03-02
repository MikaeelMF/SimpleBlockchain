package main

import (
	"crypto/sha512"
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
