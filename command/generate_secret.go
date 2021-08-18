package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	length     = 32
	characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
)

func main() {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))

		if err != nil {
			panic(err)
		}

		result[i] = characters[num.Int64()]
	}

	fmt.Println(string(result))
}
