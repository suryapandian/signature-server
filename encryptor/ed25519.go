package encryptor

import (
	"crypto/ed25519"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"log"
	"reflect"
)

type ed25519KeyPair struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

func NewEd25519Encrytor(privateKeyPem []byte) *ed25519KeyPair {
	var kp ed25519KeyPair
	kp.setKeyPair(privateKeyPem)
	return &kp
}

func (kp *ed25519KeyPair) GetPublicKey() (pubKey []byte) {
	return []byte(kp.PublicKey)
}

func (kp *ed25519KeyPair) Sign(data []byte) []byte {
	return ed25519.Sign(kp.PrivateKey, data)
}

func (kp *ed25519KeyPair) setKeyPair(privateKeyPEM []byte) {
	if len(privateKeyPEM) == 0 {
		kp.PublicKey, kp.PrivateKey, _ = ed25519.GenerateKey(nil)
		return
	}

	if err := kp.setPrivateKey(privateKeyPEM); err != nil {
		log.Fatal("Invalid Daemon key!!!")
	}
	kp.setPublicKey()

	if !kp.isValid() {
		log.Fatal("Invalid Daemon key!!!")
	}
}

var invalidDeamonKey = errors.New("Invalid Daemon key!!!")

func (kp *ed25519KeyPair) setPrivateKey(privateKeyPEM []byte) (err error) {
	var asn1PrivKey struct {
		Version          int
		ObjectIdentifier struct {
			ObjectIdentifier asn1.ObjectIdentifier
		}
		PrivateKey []byte
	}

	var block *pem.Block
	block, _ = pem.Decode(privateKeyPEM)
	if block == nil {
		err = invalidDeamonKey
		return
	}

	if _, err = asn1.Unmarshal(block.Bytes, &asn1PrivKey); err != nil {
		return
	}

	kp.PrivateKey = ed25519.NewKeyFromSeed(asn1PrivKey.PrivateKey[2:])
	return
}

func (kp *ed25519KeyPair) setPublicKey() {
	kp.PublicKey = kp.PrivateKey.Public().(ed25519.PublicKey)
	return
}

func (kp *ed25519KeyPair) isValid() bool {
	return reflect.DeepEqual(ed25519.NewKeyFromSeed(kp.PrivateKey.Seed()), kp.PrivateKey)
}
