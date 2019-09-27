// toll
package keypair

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"

	"github.com/agl/ed25519/edwards25519"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/base58"
	"github.com/myENA/secureRandom"
)

const (
	PublicKeySize  = 32
	PrivateKeySize = 64
	SignatureSize  = 64
)
const (
	DePublicKeySize  = 32
	DePrivateKeySize = 32
	DeAddressSize    = 20
)

//Create
func Create() (publicKey string, privateKey string, address string, err error) {

	var ranbuf [32]byte
	ranstr, err := secureRandom.New(32)
	if err != nil {
		return "", "", "", err
	}
	ranbyte := []byte(ranstr)
	for i := range ranbuf {
		ranbuf[i] = ranbyte[i]
	}
	public, _, err := GenerateKey(ranbuf)
	if err != nil {
		return "", "", "", err
	}

	publicKey, err = encodePublicKey(public)
	if err != nil {
		return "", "", "", err
	}
	privateKey, err = encodePrivateKey(&ranbuf)
	if err != nil {
		return "", "", "", err
	}
	address, err = encodeAddress(public)
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

	PublicKey, _, err := GenerateKey(*PrivateKey)
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
	var epub []byte
	var ret bool
	var err error
	epub, err = hex.DecodeString(publicKey)
	if err != nil {
		return false
	}

	if len(epub) != (DePublicKeySize+6) || epub[0] != 0xb0 || epub[1] != 1 {
		return false
	}

	var hash1, hash2 []byte

	dpub := epub[:DePublicKeySize+2]
	h1 := sha256.New()
	h1.Write([]byte(dpub))
	hash1 = h1.Sum(nil)
	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)
	if !(hash2[0] == epub[DePublicKeySize+2] && hash2[1] == epub[DePublicKeySize+3] && hash2[2] == epub[DePublicKeySize+4] && hash2[3] == epub[DePublicKeySize+5]) {
		return false
	}

	ret = true
	return ret
}

//Verify the private key
func CheckPrivateKey(SprivateKey string) bool {
	var ret bool

	if SprivateKey == "" {
		return false
	}

	epriv, err := base58.Decode(SprivateKey)
	if err != nil {
		return false
	}

	if !(len(epriv) == (DePrivateKeySize+9) && epriv[0] == 0xDA && epriv[1] == 0x37 && epriv[2] == 0x9F && epriv[3] == 1) {
		return false
	}

	if !(epriv[DePrivateKeySize+4] == 0x00) {
		return false
	}

	var hash1, hash2 []byte

	dpriv := epriv[:DePrivateKeySize+5]

	h1 := sha256.New()
	h1.Write([]byte(dpriv))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	if !(hash2[0] == epriv[DePrivateKeySize+5] && hash2[1] == epriv[DePrivateKeySize+6] && hash2[2] == epriv[DePrivateKeySize+7] && hash2[3] == epriv[DePrivateKeySize+8]) {
		return false
	}

	ret = true
	return ret

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
func GenerateKey(ranbuf [32]byte) (publicKey *[PublicKeySize]byte, privateKey *[PrivateKeySize]byte, err error) {
	if len(ranbuf) == 0 {
		return nil, nil, errors.New("buf is nil")
	}

	privateKey = new([PrivateKeySize]byte)
	publicKey = new([PublicKeySize]byte)
	copy(privateKey[:PublicKeySize], ranbuf[:])

	h := sha512.New()
	h.Write(privateKey[:PublicKeySize])
	digest := h.Sum(nil)

	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	var A edwards25519.ExtendedGroupElement
	var hBytes [PublicKeySize]byte
	copy(hBytes[:], digest)
	edwards25519.GeScalarMultBase(&A, &hBytes)
	A.ToBytes(publicKey)

	copy(privateKey[PublicKeySize:], publicKey[:])
	return
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
func encodePublicKey(publicKey *[PublicKeySize]byte) (GpublicKey string, err error) {
	if publicKey == nil {
		return "", errors.New("encode publicKey is error")
	}
	var ppblic [PublicKeySize]byte = *publicKey
	var str_result []byte
	var hash1, hash2 []byte
	str_result = append(str_result, 0xb0)
	str_result = append(str_result, 1)
	str_result = bytesCombine(str_result, ppblic[:PublicKeySize])
	h1 := sha256.New()
	h1.Write([]byte(str_result))
	hash1 = h1.Sum(nil)
	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)
	str_result = bytesCombine(str_result, hash2[:4])

	str := hex.EncodeToString(str_result)
	return str, nil
}

//Encoding private Key
func encodePrivateKey(privateKey *[PublicKeySize]byte) (GPrivateKey string, err error) {
	if privateKey == nil {
		return "", errors.New("encode privateKey is error")
	}
	var ppriv [32]byte = *privateKey
	var str_result []byte
	var hash1, hash2 []byte
	str_result = append(str_result, 0xDA)
	str_result = append(str_result, 0x37)
	str_result = append(str_result, 0x9F)
	str_result = append(str_result, 1)

	str_result = bytesCombine(str_result, ppriv[:])
	str_result = append(str_result, 0x00)

	h1 := sha256.New()
	h1.Write([]byte(str_result))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	str_result = bytesCombine(str_result, hash2[:4])
	var result string
	result = base58.Encode(str_result)

	return result, nil
}

//Encoding address
func encodeAddress(publicKey *[PublicKeySize]byte) (GAccoun string, err error) {
	if publicKey == nil {
		return "", errors.New("encode publicKey is error")
	}
	var ppbilc [PublicKeySize]byte = *publicKey

	var str_result []byte
	var hash1, hash2, pubSha []byte
	str_result = append(str_result, 0X01)
	str_result = append(str_result, 0X56)
	str_result = append(str_result, 1)
	var ppbilc1 []byte
	ppbilc1 = ppbilc[:]

	ShaPub := sha256.New()
	ShaPub.Write([]byte(ppbilc1))
	pubSha = ShaPub.Sum(nil)

	str_result = bytesCombine(str_result, pubSha[12:DeAddressSize+12])

	h1 := sha256.New()
	h1.Write([]byte(str_result))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	str_result = bytesCombine(str_result, hash2[:4])
	var result string
	result = base58.Encode(str_result)

	return result, nil

}

//Decode public Key
func DecodePublicKey(publicKey string) (decodePublicKey *[PublicKeySize]byte, err error) {
	if publicKey == "" {
		return nil, errors.New("decode publicKey error :publicKey is nil")
	}
	ispub := CheckPublicKey(publicKey)
	if !(ispub) {
		return nil, errors.New("check publicKey error")
	}
	var epub []byte
	epub, err = hex.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	var dpub [PublicKeySize]byte
	copy(dpub[:], epub[2:PublicKeySize+2])
	return &dpub, nil

}

//Decode private Key
func DecodePrivateKey(privateKey string) (decodePrivateKey *[DePrivateKeySize]byte, err error) {
	if privateKey == "" {
		return nil, errors.New("decode privateKey error :privateKey is nil")
	}
	ispri := CheckPrivateKey(privateKey)
	if !(ispri) {
		return nil, errors.New("check privateKey error")
	}

	epriv, err := base58.Decode(privateKey)
	if err != nil {
		return nil, err
	}

	var dpriv [DePrivateKeySize]byte
	copy(dpriv[:], epriv[4:DePrivateKeySize+4])

	return &dpriv, nil
}
