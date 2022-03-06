package main

import (
	"bytes"
	"crypto/sha512"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

// Genesis Block Hash: c4b64ddbaff431c9bf1e5b35752d8960507336098a2c612fd1374dda944e9e8b04fbefda96a7153ff82ab8b350ebfa9c66696c3a47196fbdc44ed1ce13b89f2f
// Genesis Block Hash in Bytes: [196 182 77 219 175 244 49 201 191 30 91 53 117 45 137 96 80 115 54 9 138 44 97 47 209 55 77 218 148 78 158 139 4 251 239 218 150 167 21 63 248 42 184 179 80 235 250 156 102 105 108 58 71 25 111 189 196 78 209 206 19 184 159 47]

const databaseAddress = "./dbfile"
const blocksBucket = "blocksBucket"

type Blockchain struct {
	head []byte
	db   *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var tail []byte
	var tailBlock *Block

	// Get the last block
	_ = bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		tail = b.Get([]byte("t"))
		tailBlock = Deserialize(b.Get(tail))
		return nil
	})

	// Generate the new block
	newBlock := NewBlock(data, tailBlock)

	// Add the new block to database and make the tail point to the last block
	_ = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		b.Put(newBlock.Hash[:], Serialize(newBlock))
		b.Put([]byte("t"), newBlock.Hash[:])
		return nil
	})
}

func NewBlockchain() *Blockchain {
	var head []byte
	db, _ := bolt.Open(databaseAddress, 0600, nil)
	_ = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesisBlock := &Block{blockHeight: 0, Data: []byte("Genesis Block"), prevBlockHash: [sha512.Size]byte{}, nonce: 0}
			genesisBlock.Hash = sha512.Sum512(bytes.Join([][]byte{genesisBlock.prevBlockHash[:], genesisBlock.Data, []byte(strconv.FormatUint(genesisBlock.blockHeight, 10))}, []byte{}))

			b, _ := tx.CreateBucket([]byte(blocksBucket))
			_ = b.Put(genesisBlock.Hash[:], Serialize(genesisBlock))
			_ = b.Put([]byte("t"), genesisBlock.Hash[:])
			head = genesisBlock.Hash[:]
		} else {
			head = b.Get([]byte("t"))
		}
		return nil
	})
	bc := &Blockchain{head, db}
	return bc
}
