package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

type CryptoUtil struct {
}

func (c *CryptoUtil) EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (c *CryptoUtil) DecodeBase64(data string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return decoded
}

func (c CryptoUtil) GeneratePrivatePublicKeyPair() ([]byte, []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	publicKey := &privateKey.PublicKey

	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)

	return privateKeyBytes, publicKeyBytes
}

func (c CryptoUtil) GeneratePKCS8PrivatePublicKeyPair() ([]byte, []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	return privateKeyBytes, publicKeyBytes
}

func (c *CryptoUtil) SignWihPrivateKey(msg string, privateKeyBytes []byte) []byte {
	msgHash := sha256.Sum256([]byte(msg))
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, msgHash[:])
	if err != nil {
		panic(err)
	}
	return signature
}

func (c *CryptoUtil) SignWihPKCS8PrivateKey(msg string, privateKeyBytes []byte) []byte {
	msgHash := sha256.Sum256([]byte(msg))
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
	privateKeyRSA := privateKey.(*rsa.PrivateKey)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKeyRSA, crypto.SHA256, msgHash[:])
	if err != nil {
		panic(err)
	}
	return signature
}

func (c *CryptoUtil) VerifyWithPublicKey(signature []byte, msg string, publicKeyBytes []byte) bool {
	msgHash := sha256.Sum256([]byte(msg))
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		panic(err)
	}
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, msgHash[:], signature)
	if err != nil {
		panic(err)
	}

	return true
}

func (c *CryptoUtil) VerifyWithPublicKeyPKCS8(signature []byte, msg string, publicKeyBytes []byte) bool {
	msgHash := sha256.Sum256([]byte(msg))
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		panic(err)
	}
	/*switch pub := publicKey.(type) {
	case *rsa.PublicKey:
		fmt.Println("pub is of type RSA:", pub)
	case *dsa.PublicKey:
		fmt.Println("pub is of type DSA:", pub)
	case *ecdsa.PublicKey:
		fmt.Println("pub is of type ECDSA:", pub)
	default:
		panic("unknown type of public key")
	}*/

	publicKeyRSA, isRSAPublicKey := publicKey.(*rsa.PublicKey)
	if !isRSAPublicKey {
		fmt.Println("Public key parsed is not an RSA public key")
		return false
	}

	err = rsa.VerifyPKCS1v15(publicKeyRSA, crypto.SHA256, msgHash[:], signature)
	if err != nil {
		panic(err)
	}

	return true
}
