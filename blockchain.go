package main

import (
	bolt "go.etcd.io/bbolt"
)

const databaseAddress = "./dbfile"
const blocksBucket = "blocksBucket"

type Blockchain struct {
	db       *bolt.DB
	iterator *blockchainIterator
}

type blockchainIterator struct {
	visited *Block
}

// Generates a new blockchain with a genesis block and returns the result
func NewBlockchain() *Blockchain {
	db, _ := bolt.Open(databaseAddress, 0600, nil)
	_ = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			b, _ := tx.CreateBucket([]byte(blocksBucket))
			genesisBlock := generateGenesisBlock()
			genesisBlockHash := genesisBlock.GetBlockHash()
			b.Put(genesisBlockHash[:], genesisBlock.BlockEncoder())
			b.Put([]byte("t"), genesisBlockHash[:])
		}
		return nil
	})
	return &Blockchain{db: db, iterator: &blockchainIterator{visited: nil}}
}

// AddBlock function appends a new block to current blockchain with given data
func (bc *Blockchain) AddBlock(data string) {

	// This function consists of three parts:
	// 1. Prepare by getting the last mined block in blockchain
	// 2. Generate a new block using block.go
	// 3. Append the new block

	//1. Preparation:
	var hashLastBlock []byte // hash of the last block in blockchain retrived from database using "t" key in []byte format
	var lastBlock *Block     // last block retrived from the database using its hash as a key

	bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		hashLastBlock = b.Get([]byte("t"))
		lastBlock = BlockDecoder(b.Get(hashLastBlock))
		return nil
	})

	// 2. Generation of New Block
	newBlock := NewBlock(data, lastBlock)

	// 3. Append the new block to the blockchain
	bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		newBlockHash := newBlock.GetBlockHash()
		b.Put(newBlockHash[:], newBlock.BlockEncoder()) // Append the new block to database
		b.Put([]byte("t"), newBlockHash[:])             // Update the tail for database
		return nil
	})
}

// Generates the genesis Block and returnes the result
func generateGenesisBlock() *Block {
	genesisBlock := NewBlock("Genesis Block", nil)
	return genesisBlock
}

// Returns the previous block of the last visited block
// If no block has been visited yet, it will start from the latest mined block
func (bc *Blockchain) GetPreviousBlock() *Block {
	var prevBlock *Block
	iter := bc.iterator
	if iter.visited == nil {
		bc.db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blocksBucket))
			prevBlock = BlockDecoder(b.Get([]byte("t")))
			bc.iterator.visited = prevBlock
			return nil
		})
	} else {
		bc.db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blocksBucket))
			prevBlockHash := iter.visited.GetPreviousBlockHash()
			prevBlock = BlockDecoder(b.Get(prevBlockHash[:]))
			bc.iterator.visited = prevBlock
			return nil
		})
	}
	return prevBlock
}
