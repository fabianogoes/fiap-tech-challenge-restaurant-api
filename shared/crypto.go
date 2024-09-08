package shared

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"
)

type Crypto struct {
	key []byte
}

func NewCrypto(key []byte) *Crypto {
	return &Crypto{key}
}

func (c *Crypto) EncryptAES(data string) string {
	// Cria uma instância do cifrador AES
	block, err := aes.NewCipher(c.key)
	if err != nil {
		panic(err)
	}

	// Adiciona padding ao plainText
	plainText := pkcs7Padding([]byte(data), aes.BlockSize)

	// Gera um IV aleatório
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Cria um encriptador AES no modo CBC
	stream := cipher.NewCBCEncrypter(block, iv)

	// Criptografa os dados
	stream.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText)
}

func (c *Crypto) DecryptAES(encryptedText string) string {
	// Decodifica o texto cifrado da string hexadecimal
	cipherText, err := hex.DecodeString(encryptedText)
	if err != nil {
		panic(err)
	}

	// Cria uma instância do cifrador AES
	block, err := aes.NewCipher(c.key)
	if err != nil {
		panic(err)
	}

	// Verifica o tamanho do texto cifrado
	if len(cipherText) < aes.BlockSize {
		panic(err)
	}

	// Separa o IV do texto cifrado
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Cria um decriptador AES no modo CBC
	stream := cipher.NewCBCDecrypter(block, iv)

	// Descriptografa os dados
	stream.CryptBlocks(cipherText, cipherText)

	// Remove padding após a descriptografia
	plainText := pkcs7Unpadding(cipherText)

	return string(plainText)
}

func pkcs7Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padText...)
}

func pkcs7Unpadding(paddedText []byte) []byte {
	length := len(paddedText)
	unpadding := int(paddedText[length-1])
	return paddedText[:(length - unpadding)]
}

func MaskSensitiveData(data string) string {
	visibleLength := 2

	if len(data) <= visibleLength*2 {
		return strings.Repeat("*", len(data))
	}

	return data[:visibleLength] + strings.Repeat("*", len(data)-visibleLength*2) + data[len(data)-visibleLength:]
}
