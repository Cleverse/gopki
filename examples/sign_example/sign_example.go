package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"github.com/Cleverse/gopki"
)

func main() {
	//import or generate public/private key
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	message := "HELLO TEST"

	//Signing
	signature, err := gopki.Signing(privateKey, message)
	if err != nil {
		log.Println("ERR:", err)
		return
	}

	signatureMessage := string(signature)
	log.Printf("Signature: '%x'\n", signatureMessage)

	//Decryption
	if err := gopki.VerifySign(publicKey, message, []byte(signatureMessage)); err != nil {
		log.Println("ERR:", err)
		return
	}
	log.Println("Result: VALIDATED")

}
