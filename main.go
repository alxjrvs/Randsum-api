package main

import (
	"fmt"
)

func main() {
	params := RollParams{2, 20}
	result := RollResult(params)
	fmt.Println(result)
}
