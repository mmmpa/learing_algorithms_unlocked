package main

import (
	"testing"
	"fmt"
	"io/ioutil"
)

func TestCompute(t *testing.T) {
	s := "AAAVV  AABB aaa bbb CD\n\nn aaaa"

	result, tree, rest := huffman(s)


	de := deHuffman(result, tree, rest)
	fmt.Printf("%v bytes -> %v bytes -> %v bytes\n", len([]byte(s)), len(result), len(de))

	if s != de {
		fmt.Println(s, de)
		t.Fail()
	}
}


func TestCompute2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("moby-dick.txt")
	s := string(bytes)

	result, tree, rest := huffman(s)

	de := deHuffman(result, tree, rest)
	fmt.Printf("%v bytes -> %v bytes -> %v bytes\n", len([]byte(s)), len(result), len(de))
}


func TestCompute3(t *testing.T) {
	s := "TATAGATCTTAATATAAVVCCVCVCVCV"
	en := lzw(s)
	de := delzw(en)

	fmt.Println(en)
	fmt.Println(de)

	if s != de {
		t.Fail()
	}
}
