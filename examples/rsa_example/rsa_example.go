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

	//Encryption
	cipher, err := gopki.Encrypt(publicKey, message)
	if err != nil {
		log.Println("ERR:", err)
		return
	}

	cipherMessage := string(cipher)
	log.Printf("Cipher: '%x'\n", cipherMessage)

	//Decryption
	decryptedMessage, err := gopki.Decrypt(privateKey, []byte(cipherMessage))
	if err != nil {
		log.Println("ERR:", err)
		return
	}
	log.Println("Result:", decryptedMessage)

}
