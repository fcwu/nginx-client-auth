package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	rng := rand.Reader
	message, _ := ioutil.ReadFile("sign.txt")
	hashed := sha256.Sum256(message)
	fmt.Printf("hash256\n%s", hex.Dump(hashed[:]))

	privateKey, _ := ioutil.ReadFile("final.key")
	block, _ := pem.Decode(privateKey)
	pk, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKCS8PrivateKey: %s\n", err.Error())
		os.Exit(0)
	}
	signature, err := rsa.SignPKCS1v15(rng, pk.(*rsa.PrivateKey), crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Printf("rsa.SignPKCS1v15: %s\n", err.Error())
		os.Exit(0)
	}
	signatureHex := hex.EncodeToString(signature)
	fmt.Printf("signature hex\n")
	fmt.Printf("%s\n", signatureHex)
	fmt.Printf("signature len(signature): %d\n", len(signature))

	// publicKey, _ := ioutil.ReadFile("final.pub")
	// block, _ = pem.Decode(publicKey)
	// pubkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	// if err != nil {
	// 	fmt.Printf("x509.ParsePKIXPublicKey: %s\n", err.Error())
	// 	os.Exit(0)
	// }
	b, _ := ioutil.ReadFile("final.crt")
	block, _ = pem.Decode(b)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKIXPublicKey: %s\n", err.Error())
		os.Exit(0)
	}
	pubkey := cert.PublicKey
	err = rsa.VerifyPKCS1v15(pubkey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Printf("rsa.VerifyPKCS1v15: %s\n", err.Error())
		os.Exit(0)
	}
	errBytes := hashed[:]
	errBytes[0] = '\x00'
	err = rsa.VerifyPKCS1v15(pubkey.(*rsa.PublicKey), crypto.SHA256, errBytes, signature)
	if err == nil {
		fmt.Printf("rsa.VerifyPKCS1v15: force error\n")
		os.Exit(0)
	}
}
