package e2e_simple

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

// E2e_enc generates an RSA encrypted ciphertext and private key for a given message.
//
// message: The message to be encrypted.
// Returns:
// - ciphertext: The RSA encrypted ciphertext.
// - privKey: The generated RSA private key.
// - error: Any error that occurred during the encryption process.
func E2e_enc(message string) ([]byte, *rsa.PrivateKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	pubKey := privKey.PublicKey

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &pubKey, []byte(message))
	if err != nil {
		return nil, nil, err
	}

	return ciphertext, privKey, nil
}

// E2e_dec decrypts the given ciphertext using the provided private key.
//
// It takes a *rsa.PrivateKey and a []byte ciphertext as parameters.
// The function returns a []byte plaintext and an error.
func E2e_dec(privKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, ciphertext)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return plaintext, nil
}
