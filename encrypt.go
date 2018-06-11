package ncm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"bytes"
)

const(
	iv  = "0102030405060708"
	nonce  = "0CoJUm6Qyw8W8jud"
	randStr = "a8LWv2uAtXjzSfkQ"
	encSecKey = "2d48fd9fb8e58bc9c1f14a7bda1b8e49a3520a67a2300a1f73766caee29f2411c5350bceb15ed196ca963d6a6d0b61f3734f0a0f4a172ad853f16dd06018bc5ca8fb640eaa8decd1cd41f66e166cea7a3023bd63960e656ec97751cfc7ce08d943928e9db9b35400ff3d138bda1ab511a06fbee75585191cabe0e6e63f7350d6"
)

func AesEncrypt(encStr string, key string) string{
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "err"
	}
	src := PKCS7Padding([]byte(encStr),block.BlockSize())
	blockModel := cipher.NewCBCEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(src))
	blockModel.CryptBlocks(cipherText, src)
	return base64.StdEncoding.EncodeToString(cipherText)
}

//PKCS7Padding
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
