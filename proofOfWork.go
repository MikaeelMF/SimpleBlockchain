package main

import (
	"math/big"
)

const targetBits = 24

// Returnes target for proof of work
func GetTarget() *big.Int {
	target := big.NewInt(1)
	target.Lsh(target, uint(512-targetBits))
	return target
}
