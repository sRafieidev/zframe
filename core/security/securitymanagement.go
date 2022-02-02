package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

/*
DES CBC encryption
 The length of Key is 8 bytes, and IV must be the same length.
*/
func EncryptDES_CBC(src, key, iv string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	data = PKCS5Padding(data, block.BlockSize())
	//Get CBC encryption mode
	//iv := keyByte //Use a key as vector (not recommended to use this)
	ivByte := []byte(iv)
	mode := cipher.NewCBCEncrypter(block, ivByte)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

//DESC CBC decryption
func DecryptDES_CBC(src, key, iv string) string {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	ivBye := []byte(iv)
	mode := cipher.NewCBCDecrypter(block, ivBye)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext)
}

//ECB encryption
func EncryptDES_ECB(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	//Make fill in the plain text data
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Piennium is subjected to blocks in blocksize
		//You can use the Go key to encrypt in parallel when necessary.
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out)
}

//ECB decryption
func DecryptDES_ECB(src, key string) (string, error) {
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out), nil
}

/*
   The length IV of the Key parameter must be the same length
   16 bytes - AES-128
   24 bytes - AES-192
   32 bytes - AES-256
*/
func EncryptAES_CBC(src, key, iv string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	data = PKCS5Padding(data, block.BlockSize())
	//Get CBC encryption mode
	//iv := keyByte //Use a key as vector (not recommended to use this)
	ivByte := []byte(iv)
	mode := cipher.NewCBCEncrypter(block, ivByte)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

//AES CBC decryption
func DecryptAES_CBC(src, key, iv string) string {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	//iv := keyByte //Use a key as vector (not recommended to use this)
	ivBye := []byte(iv)
	mode := cipher.NewCBCDecrypter(block, ivBye)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext)
}

//
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//Ming text reduction algorithm
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func DecryptAes128Ecb(data, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if (err != nil) {
		return nil, err
	}

	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted, nil
}

//Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
//cipher.init(Cipher.ENCRYPT_MODE, secretKey);
//return Base64.getEncoder().encodeToString(cipher.doFinal(strToEncrypt.getBytes("UTF-8")));
//setKey(secret);
//Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5PADDING");
//cipher.init(Cipher.DECRYPT_MODE, secretKey);
//return new String(cipher.doFinal(Base64.getDecoder().decode(strToDecrypt)));
//https://www.programmerall.com/article/20782087278/
