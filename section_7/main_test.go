package main

import (
	"testing"
	"fmt"
)

func TestCompute(t *testing.T) {
	shift := uint8(128)
	raw := "ABCDE XYZ あいうえお"
	encrypted := shiftEncrypt("ABCDE XYZ あいうえお", shift)

	if encrypted == raw {
		t.Fail()
		return
	}

	decrypted := shiftDecrypt(encrypted, shift)

	fmt.Println(encrypted)
	fmt.Println(decrypted)

	if decrypted != raw {
		t.Fail()
		return
	}
}

func TestCompute2(t *testing.T) {
	encryptor, decryptor := genTable()

	raw := "ABCDE XYZ あいうえお"
	encrypted := tableEncrypt("ABCDE XYZ あいうえお", encryptor)

	if encrypted == raw {
		t.Fail()
		return
	}

	decrypted := tableEncrypt(encrypted, decryptor)

	fmt.Println(encrypted)
	fmt.Println(decrypted)

	if decrypted != raw {
		t.Fail()
		return
	}
}

func TestCompute3(t *testing.T) {
	raw := "ABCDE XYZ あいうえお"
	encrypted, pad := onetimeEncrypt("ABCDE XYZ あいうえお")

	if encrypted == raw {
		t.Fail()
		return
	}

	decrypted := onetimeDecrypt(encrypted, pad)

	fmt.Println(encrypted)
	fmt.Println(pad)
	fmt.Println(decrypted)

	if decrypted != raw {
		t.Fail()
		return
	}
}

func TestCompute4(t *testing.T) {
	raw := "ABCDE XYZ あいうえお"
	encrypted, pad := onetimeBlockEncrypt("ABCDE XYZ あいうえお", 2)

	if encrypted == raw {
		t.Fail()
		return
	}

	decrypted := onetimeBlockDecrypt(encrypted, pad)

	fmt.Println(encrypted)
	fmt.Println(pad)
	fmt.Println(decrypted)

	if decrypted != raw {
		t.Fail()
		return
	}
}
