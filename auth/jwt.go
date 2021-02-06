package auth

import (
	"encoding/json"
	"log"

	jose "github.com/dvsekhvalnov/jose2go"
)

func GenerateTokenFromObject(
	in interface{},
) (
	string,
	error,
) {
	bytes, err := json.Marshal(in)

	if err != nil {
		log.Println("json marshal error")

		return "", err
	}

	privateKey := loadPrivateKey()

	token, signBytesError := jose.SignBytes(bytes, jose.RS256, privateKey)

	if signBytesError != nil {
		log.Println("token signing failed ", signBytesError)

		return "", signBytesError
	}

	return token, nil
}

func DecodeObjectFromToken(
	token string,
	out interface{},
) error {
	publicKeyBytes := loadPublicKey()

	sessionJSON, _, decodeError := jose.Decode(token, publicKeyBytes)

	if decodeError != nil {
		return decodeError
	}

	unmarshalError := json.Unmarshal([]byte(sessionJSON), &out)

	if unmarshalError != nil {
		return unmarshalError
	}

	return nil
}
