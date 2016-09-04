package gocrypto

import (
	"testing"
	"math/rand"
)

func TestDes(t *testing.T) {
	message := make([]byte, 60)
	c := 0
	getLen := func() int {
		c++
		return len(message)
	}
	for i := 0; i < getLen(); i++ {
		message[i] = 'a' + byte(rand.Intn(26))
	}
	//t.Logf("明文: %s", string(message))

	key := []byte("12345678")

	data, err := DesEncrypt(message, key)
	if err != nil {
		t.Fatal(err)
	}
	//t.Logf("密文: %s", string(data))

	data, err = DesDecrypt(data, key)
	if err != nil {
		t.Fatal(err)
	}
	//t.Logf("解密的明文: %s", string(data))

	if string(data) != string(message) {
		t.Fatalf("加密前后不同: '%s' != '%s'", string(message), string(data))
	}
}
