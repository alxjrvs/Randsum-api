package main

import (
	"crypto/rand"
	"math/big"
)

func RollRandomD6() (n *big.Int) {
	sides := big.NewInt(6)
	num, err := rand.Int(rand.Reader, sides)
	offset := big.NewInt(1)
	num.Add(num, offset)
	if err != nil {
		panic(err)
	}
	return num
}
