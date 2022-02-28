package main

import (
	"bytes"
	"crypto/sha512"
	"strconv"
)

// Genesis Block Hash: c4b64ddbaff431c9bf1e5b35752d8960507336098a2c612fd1374dda944e9e8b04fbefda96a7153ff82ab8b350ebfa9c66696c3a47196fbdc44ed1ce13b89f2f
// Genesis Block Hash in Bytes: [196 182 77 219 175 244 49 201 191 30 91 53 117 45 137 96 80 115 54 9 138 44 97 47 209 55 77 218 148 78 158 139 4 251 239 218 150 167 21 63 248 42 184 179 80 235 250 156 102 105 108 58 71 25 111 189 196 78 209 206 19 184 159 47]

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	genesisBlock := &Block{blockHeight: 0, Data: []byte("Genesis Block"), prevBlockHash: [sha512.Size]byte{}}
	genesisBlock.Hash = sha512.Sum512(bytes.Join([][]byte{genesisBlock.prevBlockHash[:], genesisBlock.Data, []byte(strconv.FormatUint(genesisBlock.blockHeight, 10))}, []byte{}))
	bc.blocks = append(bc.blocks, genesisBlock)
	return bc
}
