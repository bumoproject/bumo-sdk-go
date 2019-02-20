// account
package signature

import (
	"encoding/hex"
	"errors"

	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"golang.org/x/crypto/ed25519"
)

const (
	PublicKeySize  = 32
	PrivateKeySize = 64
	SignatureSize  = 64
)

//signature
func Sign(private string, message []byte) (sign string, err error) {
	if private == "" {
		return "", errors.New("check privateKey error : private is error")
	}
	if !keypair.CheckPrivateKey(private) {
		return "", errors.New("check privateKey error")
	}
	PrivateKey, err := keypair.DecodePrivateKey(private)
	if err != nil {
		return "", err
	}
	signByte := ed25519.Sign((*PrivateKey)[:], message)
	return hex.EncodeToString(signByte), nil
}

//verify
func Verify(public string, message []byte, sign string) bool {
	if public == "" {
		return false
	}

	dePublicKey, err := keypair.DecodePublicKey(public)
	if err != nil {
		return false
	}
	signByte, err := hex.DecodeString(sign)
	if err != nil {
		return false
	}
	var dePub [32]byte = *dePublicKey
	return ed25519.Verify(dePub[:], message, signByte)
}
