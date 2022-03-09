package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/gob"
	"errors"
	"math/big"
	"strconv"
	"time"
)

// Block structure contains necessary fields fro a block
type Block struct {
	blockHeight   uint64            // blockHeight is the distance of current block from the genesis block
	timeStamp     time.Time         // current time that the block is mined
	data          []byte            // data is the messages that we want to save in a block
	prevBlockHash [sha512.Size]byte // prevBlockHash referes to the previous block that current block is mined on top of
	hash          [sha512.Size]byte // hash is the sha_512 of nonce + blockHeight + data + prevBlockHash such that hash < target
	nonce         uint64            // a nonce that lets the hash be less that target
}

// A getter method for current block height @returns uint64
func (b *Block) GetBlockHeight() uint64 {
	return b.blockHeight
}

// A getter method for current block timestamp @returns time.Time
func (b *Block) GetTimeStamp() time.Time {
	return b.timeStamp
}

// A getter method for data @returns string
func (b *Block) GetData() string {
	return string(b.data)
}

// A getter method for previous block hash @returns [sha512.Size]byte
func (b *Block) GetPreviousBlockHash() [sha512.Size]byte {
	return b.prevBlockHash
}

// A getter method for current block hash @returns [sha512.Size]byte
func (b *Block) GetBlockHash() [sha512.Size]byte {
	return b.hash
}

// A getter method for current block nonce @returns uint64
func (b *Block) GetNonce() uint64 {
	return b.nonce
}

// returns all information correspond to a block in order of:
// height, data, previousBlockHash, currentBlockHash, nonce
func (b *Block) GetBlockInfo() (uint64, time.Time, string, [sha512.Size]byte, [sha512.Size]byte, uint64) {
	return b.GetBlockHeight(), b.GetTimeStamp(), b.GetData(), b.GetPreviousBlockHash(), b.GetBlockHash(), b.GetNonce()
}

func NewBlock(data string, prevBlock *Block) *Block {
	replaceNonce := "false"
	newBlock := &Block{blockHeight: prevBlock.GetBlockHeight() + 1, timeStamp: time.Now(), data: []byte(data), prevBlockHash: prevBlock.GetBlockHash()}

findNonce:
	res, err := newBlock.findNonce(replaceNonce)

	if res != true {
		if err.Error() == "Already contains a nonce" {

		} else if err.Error() == "Could not find a nonce" {
			newBlock.timeStamp = time.Now()
			replaceNonce = "true"
			goto findNonce
		}
	}

	return newBlock
}

// Returnes true if a block is hashed according to protocols
func (b *Block) ValidateNonce() bool {

	// Get the target
	target := GetTarget()

	// convert current block hash into big Int to compare
	var currentBlockHash big.Int
	currentBlockHash.SetBytes(b.hash[:])

	if currentBlockHash.Cmp(target) == -1 {
		return true
	}
	return false
}

// Encodes a block into a slice of bytes using gob.NewEncoder
func (b *Block) BlockEncoder() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	_ = encoder.Encode(b)
	return result.Bytes()
}

// Decodes an encoded block using gob.NewDecoder
func BlockDecoder(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	_ = decoder.Decode(&block)
	return &block
}

// Prepares Block for hashing
func prepareBlockForPOW(b *Block, nonce uint64) []byte {
	return bytes.Join([][]byte{
		b.prevBlockHash[:],
		b.data,
		[]byte(strconv.FormatUint(b.blockHeight, 10)),
		[]byte(strconv.Itoa(int(nonce))),
	}, []byte{})
}

// This function tries to find a valid nonce and returns true, nil if it does so
func (b *Block) findNonce(replace string) (bool, error) {

	// in case the work has been done before and it is not requested to redo it return
	if replace != "true" && b.GetNonce() != 0 {
		return false, errors.New("Already contains a nonce")
	}

	// Get the target
	target := GetTarget()

	// initialize temporary nonce
	var tempNonce uint64 = 1

	// this loop will look for a nonce unless it searches all uint64 and cannot find an appropriate nonce
	for tempNonce != 0 {
		tempHash := sha512.Sum512(prepareBlockForPOW(b, tempNonce))
		var intHash big.Int // converts current temporary hash into big.Int to compare it to the target
		intHash.SetBytes(tempHash[:])
		if intHash.Cmp(target) == -1 {
			b.nonce = tempNonce
			b.hash = tempHash
			break
		}
		tempNonce++
	}

	// if no nonce could be found it will return flase and error
	if tempNonce == 0 {
		return false, errors.New("Could not find a nonce")
	}

	return true, nil
}
