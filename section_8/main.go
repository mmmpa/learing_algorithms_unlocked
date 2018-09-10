package main

import (
	"math"
	"math/rand"
)

const (
	SHIFT_LIMIT = uint16(math.MaxUint8)
)

func shiftEncrypt(a string, shift uint8) string {
	bytes := []byte(a)

	for i, b := range bytes {
		bytes[i] = uint8((uint16(b) + uint16(shift)) % SHIFT_LIMIT)
	}

	return string(bytes)
}

func shiftDecrypt(a string, shift uint8) string {
	bytes := []byte(a)

	for i, b := range bytes {
		bytes[i] = uint8((uint16(b) + SHIFT_LIMIT - uint16(shift)) % SHIFT_LIMIT)
	}

	return string(bytes)
}

func genTable() ([]uint8, []uint8) {
	table := make([]uint8, math.MaxUint8)
	for i := uint8(0); i < math.MaxUint8; i++ {
		table[i] = i
	}

	encryptor := shuffle(table)
	decryptor := make([]uint8, math.MaxUint8)
	for i, v := range encryptor {
		decryptor[v] = uint8(i)
	}

	return encryptor, decryptor
}

func shuffle(data []uint8) []uint8 {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}

	return data
}

func tableEncrypt(a string, table []uint8) string {
	bytes := []byte(a)

	for i, b := range bytes {
		bytes[i] = table[b]
	}

	return string(bytes)
}

func onetimeEncrypt(a string) (string, []byte) {
	bytes := []byte(a)
	pad := make([]uint8, len(bytes))

	for i, b := range bytes {
		p := uint8(rand.Intn(math.MaxInt8))
		bytes[i] = b ^ p
		pad[i] = p
	}

	return string(bytes), pad
}

func onetimeDecrypt(a string, pad []byte) string {
	bytes := []byte(a)

	for i, b := range bytes {
		bytes[i] = b ^ pad[i]
	}

	return string(bytes)
}

func onetimeBlockEncrypt(a string, block int) (string, []byte) {
	bytes := []byte("0" + a)
	pad := make([]uint8, block)

	for i, _ := range pad {
		pad[i] = uint8(rand.Intn(math.MaxInt8))
	}

	for i := 1; i < len(bytes); i++ {
		bytes[i] = bytes[i] ^ pad[i%block] ^ bytes[i-1]
	}

	return string(bytes[1:]), pad
}

func onetimeBlockDecrypt(a string, pad []byte) string {
	bytes := []byte("0" + a)
	block := len(pad)

	for i := len(bytes) - 1; i >= 1; i-- {
		bytes[i] = bytes[i] ^ pad[i%block] ^ bytes[i-1]
	}

	return string(bytes[1:])
}
