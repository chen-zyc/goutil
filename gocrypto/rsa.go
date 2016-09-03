package gocrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RSAEncrypt RSA加密,将消息使用公钥加密后传给对方,由对方使用私钥解密
func RSAEncrypt(publicKey, message []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubIntf, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubIntf.(*rsa.PublicKey)
	maxLen := (pub.N.BitLen()+7)/8 - 11 // 每次加密明文的最大长度

	var data []byte
	// 计算密文的最大长度(加密后密文要比明文长,maxLen+11就是加密后密文的最大长度)
	if m := len(message); m%maxLen == 0 {
		data = make([]byte, 0, m/maxLen*(maxLen+11))
	} else {
		data = make([]byte, 0, (m/maxLen+1)*(maxLen+11))
	}
	for {
		if l := len(message); l == 0 {
			break
		} else if l < maxLen {
			maxLen = l
		}
		d, err := rsa.EncryptPKCS1v15(rand.Reader, pub, message[0:maxLen])
		if err != nil {
			return nil, err
		}
		message = message[maxLen:]
		data = append(data, d...)
	}
	return data, nil
}

// RSADecrypt RSA解密
func RSADecrypt(privateKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	maxLen := (pri.PublicKey.N.BitLen() + 7) / 8 // 每次解密 密文的最大长度

	var message []byte
	// 计算明文的最大长度(解密后明文要比密文短,maxLen-11就是加密是明文的最大长度)
	if m := len(data); m%maxLen == 0 {
		message = make([]byte, 0, m/maxLen*(maxLen-11))
	} else {
		message = make([]byte, 0, (m/maxLen+1)*(maxLen-11))
	}
	for {
		if l := len(data); l == 0 {
			break
		} else if l < maxLen {
			maxLen = l
		}
		d, err := rsa.DecryptPKCS1v15(rand.Reader, pri, data[0:maxLen])
		if err != nil {
			return nil, err
		}
		data = data[maxLen:]
		message = append(message, d...)
	}
	return message, nil
}
