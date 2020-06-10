package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func CheckErr(str string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", str, err.Error())
		os.Exit(1)
	}
}

func GenerateKeyAndNonce() (string, string, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", "", err
	}

	// Never use more than 2^32 random nonces with a given key because of
	// the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	return fmt.Sprintf("%x", key), fmt.Sprintf("%x", nonce), nil

}

func ValidateKeyAndNonce(keyHexStr, nonceHexStr string) ([]byte, []byte, error) {
	key, err := hex.DecodeString(keyHexStr)
	if err != nil {
		return nil, nil, err
	}

	nonce, err := hex.DecodeString(nonceHexStr)
	if err != nil {
		return nil, nil, err
	}

	return key, nonce, nil
}

func Encrypt(key []byte, nonce []byte, plainText string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesgcm.Seal(nil, nonce, []byte(plainText), nil)

	return fmt.Sprintf("%x", cipherText), nil
}

func Decrypt(key []byte, nonce []byte, cipherHexStr string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(cipherHexStr)
	if err != nil {
		return "", err
	}

	plainText, err := aesgcm.Open(nil, nonce, []byte(cipherText), nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func Test() (string, string, string, string, string, error) {
	keyHexStr, nonceHexStr, err := GenerateKeyAndNonce()
	if err != nil {
		return "", "", "", "", "", err
	}

	key, nonce, err := ValidateKeyAndNonce(keyHexStr, nonceHexStr)
	if err != nil {
		return "", "", "", "", "", err
	}

	plainTextBytes := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, plainTextBytes); err != nil {
		return "", "", "", "", "", err
	}

	plainText := fmt.Sprintf("%x", plainTextBytes)

	cipherText, err := Encrypt(key, nonce, plainText)
	if err != nil {
		return "", "", "", "", "", err
	}

	plainText2, err := Decrypt(key, nonce, cipherText)
	if err != nil {
		return "", "", "", "", "", err
	}

	return keyHexStr, nonceHexStr, plainText, cipherText, plainText2, nil
}

func main() {
	keyPtr := flag.String("key", "", "cipher key (hex string)")
	noncePtr := flag.String("nonce", "", "nonce (hex string)")
	generatePtr := flag.Bool("generate", false, "generate cipher key and nonce (hex strings)")
	decryptPtr := flag.Bool("decrypt", false, "decrypt instead of encrypt")
	plainTextPtr := flag.String("plaintext", "", "plaintext to encrypt (string)")
	cipherTextPtr := flag.String("ciphertext", "", "ciphertext to decrypt (hex string)")
	testPtr := flag.Bool("test", false, "test mode")
	flag.Parse()

	switch {
	case *testPtr:
		key, nonce, plainText, cipherText, plainText2, err := Test()
		CheckErr("test", err)

		if plainText != plainText2 {
			CheckErr("test", fmt.Errorf("wrong decrypted text"))
		}

		fmt.Printf("key: %s\n", key)
		fmt.Printf("nonce: %s\n", nonce)
		fmt.Printf("plaintext: %s\n", plainText)
		fmt.Printf("ciphertext: %s\n", cipherText)
		fmt.Printf("plaintext2: %s\n", plainText2)
	case *generatePtr:
		key, nonce, err := GenerateKeyAndNonce()
		CheckErr("generate key/nonce", err)

		fmt.Printf("key: %s\n", key)
		fmt.Printf("nonce: %s\n", nonce)
	case *decryptPtr:
		key, nonce, err := ValidateKeyAndNonce(*keyPtr, *noncePtr)
		CheckErr("validate key/nonce", err)

		plainText, err := Decrypt(key, nonce, *cipherTextPtr)
		CheckErr("decrypt", err)

		fmt.Printf("plaintext: %s\n", plainText)
	default:
		key, nonce, err := ValidateKeyAndNonce(*keyPtr, *noncePtr)
		CheckErr("validate key/nonce", err)

		cipherText, err := Encrypt(key, nonce, *plainTextPtr)
		CheckErr("encrypt", err)

		fmt.Printf("ciphertext: %s\n", cipherText)

	}
}
