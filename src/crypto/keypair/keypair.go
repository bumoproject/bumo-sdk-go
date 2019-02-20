// toll
package keypair

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/bumoproject/bumo-sdk-go/src/crypto/base58"
	"github.com/myENA/secureRandom"
	"golang.org/x/crypto/ed25519"
)

const (
	PublicKeySize  = 32
	PrivateKeySize = 32
	SignatureSize  = 64
)
const (
	DePublicKeySize  = 32
	DePrivateKeySize = 64
	DeAddressSize    = 20
)

//Create
func Create() (publicKey string, privateKey string, address string, err error) {
	ranstr, err := secureRandom.New(32)
	if err != nil {
		return "", "", "", err
	}

	dePublic, dePriv, err := generateKey([]byte(ranstr))
	if err != nil {
		return "", "", "", err
	}

	publicKey, err = encodePublicKey(dePublic)
	if err != nil {
		return "", "", "", err
	}

	privateKey, err = encodePrivateKey(dePriv)
	if err != nil {
		return "", "", "", err
	}

	address, err = encodeAddress(dePublic)
	if err != nil {
		return "", "", "", err
	}

	return publicKey, privateKey, address, nil

}

//The private key gets the public key
func GetEncPublicKey(privateKey string) (publicKey string, err error) {
	if CheckPrivateKey(privateKey) == false {
		return "", errors.New("privateKey error")
	}

	PrivateKey, err := DecodePrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	PublicKey, _, err := generateKey((*PrivateKey)[:32])
	if err != nil {
		return "", err
	}

	return encodePublicKey(PublicKey)
}

//The public key gets the address
func GetEncAddress(publicKey string) (address string, err error) {
	if CheckPublicKey(publicKey) == false {
		return "", errors.New("publicKey error")
	}

	PublicKey, err := DecodePublicKey(publicKey)
	if err != nil {
		return "", err
	}

	return encodeAddress(PublicKey)
}

//Verify the public key
func CheckPublicKey(publicKey string) bool {
	if publicKey == "" {
		return false
	}

	var pub []byte
	var err error
	pub, err = hex.DecodeString(publicKey)
	if err != nil {
		return false
	}

	if len(pub) != (DePublicKeySize+6) || pub[0] != 0xb0 || pub[1] != 1 {
		return false
	}

	var hash1, hash2 []byte

	dePub := pub[:DePublicKeySize+2]
	h1 := sha256.New()
	h1.Write([]byte(dePub))
	hash1 = h1.Sum(nil)
	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)
	if !(hash2[0] == pub[DePublicKeySize+2] && hash2[1] == pub[DePublicKeySize+3] && hash2[2] == pub[DePublicKeySize+4] && hash2[3] == pub[DePublicKeySize+5]) {
		return false
	}

	return true
}

//Verify the private key
func CheckPrivateKey(privateKey string) bool {
	if privateKey == "" {
		return false
	}

	priv, err := base58.Decode(privateKey)
	if err != nil {
		return false
	}

	if !(len(priv) == (PrivateKeySize+9) && priv[0] == 0xDA && priv[1] == 0x37 && priv[2] == 0x9F && priv[3] == 1) {
		return false
	}

	if !(priv[PrivateKeySize+4] == 0x00) {
		return false
	}

	var hash1, hash2 []byte

	dpriv := priv[:PrivateKeySize+5]

	h1 := sha256.New()
	h1.Write([]byte(dpriv))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	if !(hash2[0] == priv[PrivateKeySize+5] && hash2[1] == priv[PrivateKeySize+6] && hash2[2] == priv[PrivateKeySize+7] && hash2[3] == priv[DePrivateKeySize-32+8]) {
		return false
	}

	return true

}

//Verify the address key
func CheckAddress(Saddress string) bool {
	if Saddress == "" {
		return false
	}
	var addre []byte
	var ret bool
	var err error
	addre, err = base58.Decode(Saddress)
	if err != nil {
		return false
	}

	if !(addre[0] == 0X01 && addre[1] == 0X56) {
		return false
	} else if !(addre[2] == 1) {
		return false
	}
	var hash1, hash2 []byte

	daddr := addre[:DeAddressSize+3]

	h1 := sha256.New()
	h1.Write([]byte(daddr))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	if !(hash2[0] == addre[DeAddressSize+3] && hash2[1] == addre[DeAddressSize+4] && hash2[2] == addre[DeAddressSize+5] && hash2[3] == addre[DeAddressSize+6]) {
		return false
	}
	ret = true

	return ret

}

//Generate Key
func generateKey(ranbuf []byte) (*[DePublicKeySize]byte, *[DePrivateKeySize]byte, error) {
	var publicKey [DePublicKeySize]byte
	var privateKey [DePrivateKeySize]byte
	ranBytes := make([]byte, 32)
	copy(ranBytes[:], ranbuf[:])
	rand := strings.NewReader(string(ranBytes))
	pubKey, priKey, err := ed25519.GenerateKey(rand)
	if err != nil {
		return nil, nil, err
	}

	copy(publicKey[:], pubKey[:])
	copy(privateKey[:], priKey[:])
	return &publicKey, &privateKey, nil
}

func bytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

//Encoding public Key
func encodePublicKey(dePublicKey *[DePublicKeySize]byte) (publicKey string, err error) {
	if dePublicKey == nil {
		return "", errors.New("encode publicKey is error")
	}
	var pblic [DePublicKeySize]byte = *dePublicKey
	var resultStr []byte
	var hash1, hash2 []byte
	resultStr = append(resultStr, 0xb0)
	resultStr = append(resultStr, 1)
	resultStr = bytesCombine(resultStr, pblic[:DePublicKeySize])
	h1 := sha256.New()
	h1.Write([]byte(resultStr))
	hash1 = h1.Sum(nil)
	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)
	resultStr = bytesCombine(resultStr, hash2[:4])
	publicKey = hex.EncodeToString(resultStr)
	return publicKey, nil
}

//Encoding private Key
func encodePrivateKey(dePrivateKey *[DePrivateKeySize]byte) (privateKey string, err error) {
	if dePrivateKey == nil {
		return "", errors.New("encode privateKey is error")
	}
	var priv [32]byte
	copy(priv[:], (*dePrivateKey)[:32])
	var resultStr []byte
	var hash1, hash2 []byte
	resultStr = append(resultStr, 0xDA)
	resultStr = append(resultStr, 0x37)
	resultStr = append(resultStr, 0x9F)
	resultStr = append(resultStr, 1)

	resultStr = bytesCombine(resultStr, priv[:])
	resultStr = append(resultStr, 0x00)

	h1 := sha256.New()
	h1.Write([]byte(resultStr))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	resultStr = bytesCombine(resultStr, hash2[:4])
	privateKey = base58.Encode(resultStr)
	return privateKey, nil
}

//Encoding address
func encodeAddress(dePublicKey *[DePublicKeySize]byte) (address string, err error) {
	if dePublicKey == nil {
		return "", errors.New("encode publicKey is error")
	}

	var resultStr []byte
	var hash1, hash2, pubSha []byte
	resultStr = append(resultStr, 0X01)
	resultStr = append(resultStr, 0X56)
	resultStr = append(resultStr, 1)

	ShaPub := sha256.New()
	ShaPub.Write((*dePublicKey)[:])
	pubSha = ShaPub.Sum(nil)

	resultStr = bytesCombine(resultStr, pubSha[12:DeAddressSize+12])

	h1 := sha256.New()
	h1.Write([]byte(resultStr))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	resultStr = bytesCombine(resultStr, hash2[:4])
	address = base58.Encode(resultStr)

	return

}

//Decode public Key
func DecodePublicKey(publicKey string) (decodePublicKey *[DePublicKeySize]byte, err error) {
	if publicKey == "" {
		return nil, errors.New("decode publicKey error :publicKey is nil")
	}
	if !(CheckPublicKey(publicKey)) {
		return nil, errors.New("check publicKey error")
	}
	var pub []byte
	pub, err = hex.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	var dePub [DePublicKeySize]byte
	copy(dePub[:], pub[2:DePublicKeySize+2])
	return &dePub, nil

}

//Decode private Key
func DecodePrivateKey(privateKey string) (decodePrivateKey *[DePrivateKeySize]byte, err error) {
	if !(CheckPrivateKey(privateKey)) {
		return nil, errors.New("check privateKey error")
	}
	ranbuf, err := base58.Decode(privateKey)
	if err != nil {
		return nil, err
	}
	_, decodePrivateKey, err = generateKey(ranbuf[4:36])
	if err != nil {
		return nil, err
	}
	return decodePrivateKey, nil
}
