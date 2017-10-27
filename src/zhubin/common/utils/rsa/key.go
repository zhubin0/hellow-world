package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"strings"
)

type Key interface {
	PublicKey() *rsa.PublicKey
	PrivateKey() *rsa.PrivateKey
	Modulus() int
}

//parse public key
func ParsePKCS8PubKey(publicKey []byte) (Key, error) {
	puk, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return &key{publicKey: puk.(*rsa.PublicKey)}, nil
}

//parse private key
func ParsePKCS8PriKey(privateKey []byte) (Key, error) {
	prk, err := x509.ParsePKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &key{privateKey: prk.(*rsa.PrivateKey)}, nil
}

func ParsePKCS8Key(publicKey, privateKey []byte) (Key, error) {
	puk, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	prk, err := x509.ParsePKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &key{publicKey: puk.(*rsa.PublicKey), privateKey: prk.(*rsa.PrivateKey)}, nil
}

func ParsePKCS1Key(publicKey, privateKey []byte) (Key, error) {
	puk, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	prk, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &key{publicKey: puk.(*rsa.PublicKey), privateKey: prk}, nil
}

func LoadKeyFromPEMFile(publicKeyFilePath, privateKeyFilePath string, ParseKey func([]byte, []byte) (Key, error)) (Key, error) {

	publicKeyFilePath = strings.TrimSpace(publicKeyFilePath)

	pukBytes, err := ioutil.ReadFile(publicKeyFilePath)
	if err != nil {
		return nil, err
	}

	puk, _ := pem.Decode(pukBytes)
	if puk == nil {
		return nil, errors.New("publicKey is not pem format")
	}

	privateKeyFilePath = strings.TrimSpace(privateKeyFilePath)

	prkBytes, err := ioutil.ReadFile(privateKeyFilePath)
	if err != nil {
		return nil, err
	}

	prk, _ := pem.Decode(prkBytes)
	if prk == nil {
		return nil, errors.New("privateKey is not pem format")
	}

	return ParseKey(puk.Bytes, prk.Bytes)
}

type key struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// either public key or private key must be provided when encryption or decryption, but not both.
func (key *key) Modulus() int {
	if key.privateKey != nil {
		return len(key.privateKey.N.Bytes())
	} else {
		return len(key.publicKey.N.Bytes())
	}
}

func (key *key) PublicKey() *rsa.PublicKey {
	return key.publicKey
}

func (key *key) PrivateKey() *rsa.PrivateKey {
	return key.privateKey
}
