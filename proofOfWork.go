package main

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"math/big"
	"strconv"
)

const targetBits = 24

type POW struct {
	block  *Block
	target *big.Int
}

// It is better to make this a method as well for consistency
func NewProofOfWork(b *Block) *POW {
	target := big.NewInt(1)
	target.Lsh(target, uint(512-targetBits))
	return &POW{b, target}
}

func (pow *POW) PrepareData(nonce uint64) []byte {
	//Why target bits are added??
	return bytes.Join([][]byte{
		pow.block.prevBlockHash[:],
		pow.block.Data,
		[]byte(strconv.FormatUint(pow.block.blockHeight, 10)),
		[]byte(strconv.Itoa(int(nonce))),
	}, []byte{})
}

func (pow *POW) Run() (uint64, [sha512.Size]byte) {
	var intHash big.Int
	var hash [sha512.Size]byte
	var nonce uint64
	nonce = 1

	for nonce != 0 {
		hash = sha512.Sum512(pow.PrepareData(nonce))
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}
	fmt.Printf("Nonce is: %d, Hash is : %x\n", nonce, hash)
	fmt.Print("\n\n")
	return nonce, hash
}

func (pow *POW) validate() bool {
	var intHash big.Int
	newHash := sha512.Sum512(pow.PrepareData(pow.block.nonce))
	intHash.SetBytes(pow.block.Hash[:])
	return newHash == pow.block.Hash && intHash.Cmp(pow.target) == -1
}
