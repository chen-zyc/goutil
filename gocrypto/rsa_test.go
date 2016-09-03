package gocrypto

import (
	"encoding/base64"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

/*
创建私钥:

openssl genrsa -out private.pem 1024 //密钥长度，1024觉得不够安全的话可以用2048，但是代价也相应增大

创建公钥:

openssl rsa -in private.pem -pubout -out public.pem
*/

var canTest bool

func init() {
	exist := func(fname string) bool {
		_, err := os.Stat(fname)
		if err == nil || err == os.IsExist(err) {
			return true
		}
		return false
	}
	canTest = exist("public.pem") && exist("private.pem")
}

var readFile = func(fname string) []byte {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return data
}

func TestRSA(t *testing.T) {
	if !canTest {
		return
	}
	publicKey := readFile("public.pem")
	privateKey := readFile("private.pem")

	message := "github.com/zhangyuchen0411"

	data, err := RSAEncrypt(publicKey, []byte(message))
	if err != nil {
		t.Fatal(err)
	}
	transfMsg := base64.StdEncoding.EncodeToString(data)
	//t.Logf("加密后的字符串: %s", transfMsg)

	// ========== 传输中 ==================

	data, err = base64.StdEncoding.DecodeString(transfMsg)
	if err != nil {
		t.Fatal(err)
	}
	data, err = RSADecrypt(privateKey, data)
	if err != nil {
		t.Fatal(err)
	}
	//t.Logf("解密后的字符串: %s", string(data))

	if string(data) != message {
		t.Fatalf("加密前后字符串不同, '%s' != '%s'", message, string(data))
	}
}

func TestRSALongMessage(t *testing.T) {
	if !canTest {
		return
	}
	publicKey := readFile("public.pem")
	privateKey := readFile("private.pem")

	message := make([]byte, 1024)
	for i := 0; i < len(message); i++ {
		message[i] = 'a' + byte(rand.Intn(26))
	}
	//t.Logf("消息: %s", string(message))
	//t.Logf("公钥长度: %d, 私钥长度: %d, 消息长度: %d", len(publicKey), len(privateKey), len(message))

	data, err := RSAEncrypt(publicKey, message)
	if err != nil {
		t.Fatal(err)
	}
	//t.Logf("加密后的消息长度: %d", len(data))
	data, err = RSADecrypt(privateKey, data)
	if err != nil {
		t.Fatal(err)
	}
	//t.Logf("解密后的字符串: %s", string(data))

	if string(data) != string(message) {
		t.Fatalf("加密前后字符串不同, '%s' != '%s'", string(message), string(data))
	}
}

func BenchmarkRSAEncrypt(b *testing.B) {
	if !canTest {
		return
	}
	publicKey := readFile("public.pem")
	message := make([]byte, 1024)
	for i := 0; i < len(message); i++ {
		message[i] = 'a' + byte(rand.Intn(26))
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RSAEncrypt(publicKey, message)
	}
}

func BenchmarkRSADecrypt(b *testing.B) {
	if !canTest {
		return
	}
	publicKey := readFile("public.pem")
	privateKey := readFile("private.pem")
	message := make([]byte, 1024)
	for i := 0; i < len(message); i++ {
		message[i] = 'a' + byte(rand.Intn(26))
	}

	data, err := RSAEncrypt(publicKey, []byte(message))
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RSADecrypt(privateKey, data)
	}
}
