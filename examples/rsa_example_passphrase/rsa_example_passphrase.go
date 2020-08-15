package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/Cleverse/gopki"
)

func main() {
	//import private key from *.pem with passphrase
	pembytes, err := ioutil.ReadFile("./keys/private_key.pem")
	if err != nil {
		log.Println("ERR:", err)
	}

	block, _ := pem.Decode(pembytes)

	//enter passphrase
	passphrase := []byte("1234567890")
	der, err := x509.DecryptPEMBlock(block, passphrase)
	if err != nil {
		log.Println("ERR:", err)
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		log.Println("ERR:", err)
		return
	}
	publicKey := &privateKey.PublicKey

	//example message
	message := "HELLO TEST WITH IMPORTED PRIVATE KEY"

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
