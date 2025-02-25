package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"os"
)

type Encryption struct {
	environment *Environment
	logger      *helpers.Logger
	privateKey  *rsa.PrivateKey
}

func NewEncryption(environment *Environment) (*Encryption, error) {
	logger := environment.GetLogger()
	key, err := os.ReadFile(environment.GetKeyPath())
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("Failed to parse PEM block containing the private key.")
	}

	var rsaPrivateKey *rsa.PrivateKey
	switch block.Type {
	case "RSA PRIVATE KEY":
		logger.Info("Private key is in PKCS#1 format.")
		rsaPrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse PEM block containing the private PKCS#1 key: %v", err)
		}
	case "PRIVATE KEY":
		logger.Info("The key is in PKCS#8 format.")
		rsaPrivateKeyBinary, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse PEM block containing the private PKCS#8 key: %v", err)
		}
		var ok bool
		rsaPrivateKey, ok = rsaPrivateKeyBinary.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("Not an RSA private key in file %s.", environment.GetKeyPath())
		}
	default:
		return nil, fmt.Errorf("Unknown key format.")
	}

	return &Encryption{environment: environment, logger: logger, privateKey: rsaPrivateKey}, nil
}

func NewEncryptionPkcs1(environment *Environment) *Encryption {
	key, error := os.ReadFile(environment.GetKeyPath())
	if error != nil {
		environment.GetLogger().Fatalf("Error reading file: %v", error)
	}

	block, _ := pem.Decode(key)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		environment.GetLogger().Fatal("Failed to parse PEM block containing the private key.")
	}

	rsaPrivateKey, error := x509.ParsePKCS1PrivateKey(block.Bytes)
	if error != nil {
		environment.GetLogger().Fatalf("Failed to parse PKCS#1 private key: %v", error)
	}

	return &Encryption{environment: environment, logger: environment.GetLogger(), privateKey: rsaPrivateKey}
}

func NewEncryptionPkcs8(environment *Environment) *Encryption {
	key, error := os.ReadFile(environment.GetKeyPath())
	if error != nil {
		environment.GetLogger().Fatalf("Error reading file: %v", error)
	}

	block, _ := pem.Decode(key)
	if block == nil || block.Type != "PRIVATE KEY" {
		environment.GetLogger().Fatal("Failed to parse PEM block containing the private key.")
	}

	priv, error := x509.ParsePKCS8PrivateKey(block.Bytes)
	if error != nil {
		environment.GetLogger().Fatalf("Failed to parse PKCS#8 private key: %v", error)
	}

	rsaPrivKey, ok := priv.(*rsa.PrivateKey)
	if !ok {
		environment.GetLogger().Fatalf("Not an RSA private key in file %s.", environment.GetKeyPath())
	}

	return &Encryption{environment: environment, logger: environment.GetLogger(), privateKey: rsaPrivKey}
}

func (this *Encryption) Encrypt(plainValue string) (encryptedValue string, err error) {
	publicKey := this.privateKey.PublicKey
	var encryptedValueByteSlices []byte
	encryptedValueByteSlices, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte(plainValue), nil)
	if err != nil {
		this.logger.Fatalf("Encryption failed: %v", err)
		return "", err
	}
	return string(encryptedValueByteSlices), nil
}

func (this *Encryption) Decrypt(encryptedValue string) (decryptedValue string, err error) {
	plainValue, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, this.privateKey, []byte(encryptedValue), nil)
	if err != nil {
		this.logger.Errorf("Decryption failed: %v", err)
		return "", err
	}
	return string(plainValue), nil
}
