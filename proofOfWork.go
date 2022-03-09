package main

import (
	"crypto/sha512"
	"math/big"
)

const targetBits = 24

// Returnes target for proof of work
func GetTarget() *big.Int {
	target := big.NewInt(1)
	target.Lsh(target, uint(sha512.Size-targetBits))
	return target
}
