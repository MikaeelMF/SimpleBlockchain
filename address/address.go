package address

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const addressChecksumLen = 4

type KeyPair struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewKeyPair() KeyPair {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	pubkey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return KeyPair{PrivateKey: *private, PublicKey: pubkey}
}

// Generates a public address in Bitcoin format
func (pair KeyPair) GetAddress() []byte {

	pubKeyHash := HashPubKey(pair.PublicKey)

	versiondPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versiondPayload)

	fullPayload := append(versiondPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

func HashPubKey(pubkey []byte) []byte {
	publicSha256 := sha256.Sum256(pubkey)

	ripemdHasher := ripemd160.New()
	_, err := ripemdHasher.Write(publicSha256[:])
	if err != nil {
		panic(err)
	}

	publicRIPEMD160 := ripemdHasher.Sum(nil)
	return publicRIPEMD160
}

func checksum(payload []byte) []byte {
	firstSum := sha256.Sum256(payload)
	scndSum := sha256.Sum256(firstSum[:])

	return scndSum[:addressChecksumLen]
}
