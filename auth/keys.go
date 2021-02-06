package auth

import (
	"crypto/rsa"
	"io/ioutil"

	Rsa "github.com/dvsekhvalnov/jose2go/keys/rsa"
	"github.com/j45k4/talktocow/config"
)

func loadPrivateKey() *rsa.PrivateKey {
	privateKeyBytes, err := ioutil.ReadFile(config.PrivateKeyPath)

	if err != nil {
		panic("Reading private key failed")
	}

	privateKey, err := Rsa.ReadPrivate(privateKeyBytes)

	if err != nil {
		panic("Reading private key failed")
	}

	return privateKey
}

func loadPublicKey() *rsa.PublicKey {
	publicKeyBytes, err := ioutil.ReadFile(config.PublicKeyPath)

	if err != nil {
		panic("No public key found")
	}

	publicKey, err := Rsa.ReadPublic(publicKeyBytes)

	if err != nil {
		panic("Read public key failed")
	}

	return publicKey
}
