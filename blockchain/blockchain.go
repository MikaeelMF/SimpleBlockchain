package blockchain

import (
	"../block"
	"github.com/syndtr/goleveldb/leveldb"
)

const databaseAddress = "./database"

type Blockchain struct {
	db       *leveldb.DB
	iterator *blockchainIterator
}

type blockchainIterator struct {
	visited *block.Block
}

// Generates a new blockchain with a genesis block and returns the result
func InitBlockchain() *Blockchain {
	ldb, _ := leveldb.OpenFile(databaseAddress, nil)
	_, err := ldb.Get([]byte("t"), nil)
	if err != nil {
		gb := genesisBlockGenerator()
		gbHash := gb.GetBlockHash()
		ldb.Put([]byte("t"), gbHash[:], nil)
		ldb.Put(gbHash[:], gb.BlockEncoder(), nil)
	}
	return &Blockchain{db: ldb, iterator: &blockchainIterator{nil}}
}

// Generates the genesis Block and returnes the result
func genesisBlockGenerator() *block.Block {
	return block.NewBlock("Genesis Block", nil)
}

func (bc *Blockchain) AddBlock(data string) {
	// This function consists of three parts:
	// 1. Prepare by getting the last mined block in blockchain
	// 2. Generate a new block using block.go
	// 3. Append the new block

	//1. Preparation:
	var hashLastBlock []byte    // hash of the last block in blockchain retrived from database using "t" key in []byte format
	var lastBlockEncoded []byte // last block in encoded format retrived from the database using its hash as a key
	var lastBlock *block.Block  // last block decoded

	hashLastBlock, _ = bc.db.Get([]byte("t"), nil)
	lastBlockEncoded, _ = bc.db.Get(hashLastBlock, nil)
	lastBlock = block.BlockDecoder(lastBlockEncoded)

	// 2. Generation of New Block
	newBlock := block.NewBlock(data, lastBlock)

	// 3. Append the new block to the blockchain
	newBlockHash := newBlock.GetBlockHash()
	bc.db.Put(newBlockHash[:], newBlock.BlockEncoder(), nil)
	bc.db.Put([]byte("t"), newBlockHash[:], nil)
}

// Returns the previous block of the last visited block
// If no block has been visited yet, it will start from the latest mined block

func (bc *Blockchain) GetPreviousBlock() *block.Block {
	var prevBlockHash []byte
	var prevBlock *block.Block
	iter := bc.iterator

	if iter.visited == nil {
		prevBlockHash, _ = bc.db.Get([]byte("t"), nil)
		if prevBlockHash == nil {
			panic("No Block in the Blockchain. Check your Initialization")
		}
	} else {
		prevBlockHashTemp := iter.visited.GetPreviousBlockHash()
		prevBlockHash = prevBlockHashTemp[:]
	}
	prevBlockEncoded, _ := bc.db.Get(prevBlockHash[:], nil)
	prevBlock = block.BlockDecoder(prevBlockEncoded)
	bc.iterator.visited = prevBlock

	return prevBlock
}
