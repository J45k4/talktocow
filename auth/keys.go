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
		panic("no private key found")
	}

	privateKey, err := Rsa.ReadPrivate(privateKeyBytes)

	if err != nil {
		panic("parsing private key failed")
	}

	return privateKey
}

func loadPublicKey() *rsa.PublicKey {
	publicKeyBytes, err := ioutil.ReadFile(config.PublicKeyPath)

	if err != nil {
		panic("no public key found")
	}

	publicKey, err := Rsa.ReadPublic(publicKeyBytes)

	if err != nil {
		panic("parsing public key failed")
	}

	return publicKey
}
