package smileidentity

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

// calculateSecretKey calculates the secret key added to the request body using the specifications gotten from the smileidentity documentation.
// link for reference
func calculateSecretKey(key, partnerID string, timestamp int64) (string, error) {
	msg := fmt.Sprintf("%s:%d", strings.TrimSpace(partnerID), timestamp)
	sha := sha256.New()
	_, err := sha.Write([]byte(msg))
	if err != nil {
		return "", err
	}

	hash := hex.EncodeToString(sha.Sum(nil))

	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	encryptedHash, err := rsaEncrypt([]byte(hash), decodedKey)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedHash) + "|" + hash, nil
}

func rsaEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	pkixKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pkixKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to assert public key type")
	}
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, origData)
}
