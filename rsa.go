package gopki

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func Encrypt(publicKey *rsa.PublicKey, message string) ([]byte, error) {
	secretMessage := []byte(message)
	rng := rand.Reader
	return rsa.EncryptOAEP(sha256.New(), rng, publicKey, secretMessage, nil)
}

func Decrypt(privateKey *rsa.PrivateKey, cipherMessage []byte) (string, error) {
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, privateKey, cipherMessage, nil)
	return string(plaintext), err
}

func Signing(privateKey *rsa.PrivateKey, message string) ([]byte, error) {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(message))
	if err != nil {
		return nil, err
	}
	hashed := msgHash.Sum(nil)

	rng := rand.Reader
	return rsa.SignPKCS1v15(rng, privateKey, crypto.SHA256, hashed)
}

func VerifySign(publicKey *rsa.PublicKey, message string, signature []byte) error {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(message))
	if err != nil {
		return err
	}
	hashed := msgHash.Sum(nil)

	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, signature)
}
