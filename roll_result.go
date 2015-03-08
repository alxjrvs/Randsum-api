package main

import (
	"crypto/rand"
	"math/big"
)

type RollParams struct {
	NumberOfRolls int64
	DieSides      int64
}

type TotalResult struct {
	Rolls []*big.Int `json:"rolls"`
	Total *big.Int   `json:"total"`
}

func RollResult(params RollParams) (result TotalResult) {
	rollsArray := make([]*big.Int, params.NumberOfRolls)
	total := big.NewInt(0)
	for i, _ := range rollsArray {
		roll := RollSingleD(params.DieSides)
		rollsArray[i] = roll
		total.Add(total, roll)
	}
	r := TotalResult{rollsArray, total}
	return r
}

func RollSingleD(sides int64) (n *big.Int) {
	bigSides := big.NewInt(sides)
	num, _ := rand.Int(rand.Reader, bigSides)
	offset := big.NewInt(1)
	num.Add(num, offset)
	return num
}
