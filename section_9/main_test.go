package main

import (
	"testing"
	"github.com/k0kubun/pp"
	"fmt"
)

func TestCompute(t *testing.T) {
	s := "AAAVVAABBCDE"

	result, tree, rest := huffman(s)

	fmt.Printf("%v bytes -> %v bytes\n", len([]byte(s)), len(result))

	pp.Println(deHuffman(result, tree, rest))
}
