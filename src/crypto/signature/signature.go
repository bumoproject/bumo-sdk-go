// account
package signature

import (
	"crypto/sha512"
	"crypto/subtle"
	"encoding/hex"
	"errors"

	"github.com/agl/ed25519/edwards25519"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
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

	var priv [32]byte
	var pri [32]byte
	pri = *PrivateKey
	for i := range priv {
		priv[i] = pri[i]
	}

	_, privateKey, _ := keypair.GenerateKey(priv)

	h := sha512.New()

	h.Write(privateKey[:32])

	var digest1, messageDigest, hramDigest [64]byte
	var expandedSecretKey [32]byte
	h.Sum(digest1[:0])
	copy(expandedSecretKey[:], digest1[:])
	expandedSecretKey[0] &= 248
	expandedSecretKey[31] &= 63
	expandedSecretKey[31] |= 64

	h.Reset()
	h.Write(digest1[32:])
	h.Write(message)
	h.Sum(messageDigest[:0])

	var messageDigestReduced [32]byte
	edwards25519.ScReduce(&messageDigestReduced, &messageDigest)
	var R edwards25519.ExtendedGroupElement
	edwards25519.GeScalarMultBase(&R, &messageDigestReduced)

	var encodedR [32]byte
	R.ToBytes(&encodedR)

	h.Reset()
	h.Write(encodedR[:])
	h.Write(privateKey[32:])
	h.Write(message)
	h.Sum(hramDigest[:0])
	var hramDigestReduced [32]byte
	edwards25519.ScReduce(&hramDigestReduced, &hramDigest)

	var s [32]byte
	edwards25519.ScMulAdd(&s, &hramDigestReduced, &expandedSecretKey, &messageDigestReduced)

	signature := new([64]byte)
	copy(signature[:], encodedR[:])
	copy(signature[32:], s[:])
	var sig []byte
	sig = signature[:]
	sign = hex.EncodeToString(sig)
	return sign, nil
}

//verify
func Verify(public string, message []byte, Sign string) bool {
	if public == "" {
		return false
	}

	sig, err := hex.DecodeString(Sign)
	if err != nil {
		return false
	}
	PublicKey, errp := keypair.DecodePublicKey(public)
	if errp != nil {
		return false
	}
	var publicKey [32]byte
	var pub [32]byte
	pub = *PublicKey
	for i := range pub {
		publicKey[i] = pub[i]
	}
	if sig[63]&224 != 0 {
		return false
	}

	var A edwards25519.ExtendedGroupElement
	if !A.FromBytes(&publicKey) {
		return false
	}
	edwards25519.FeNeg(&A.X, &A.X)
	edwards25519.FeNeg(&A.T, &A.T)

	h := sha512.New()
	h.Write(sig[:32])
	h.Write(publicKey[:])
	h.Write(message)
	var digest [64]byte
	h.Sum(digest[:0])

	var hReduced [32]byte
	edwards25519.ScReduce(&hReduced, &digest)

	var R edwards25519.ProjectiveGroupElement
	var b [32]byte
	copy(b[:], sig[32:])
	edwards25519.GeDoubleScalarMultVartime(&R, &hReduced, &A, &b)

	var checkR [32]byte
	R.ToBytes(&checkR)
	return subtle.ConstantTimeCompare(sig[:32], checkR[:]) == 1
}
