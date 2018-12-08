package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/crypto/nacl/box"
)

func Test_Encrypt(t *testing.T) {
	pubkey := "MC0CAQACBQCkuQVJAgMBAAECBCJXm4ECAwDOtwIDAMv/AgMAvYcCAgWzAgMAqsc="
	s, err := EncryptToPubKey(pubkey, "This a encrypt msessage")
	if err != nil {
		log.Println("couldn't encrypt the secret message \n %s", err)
	}
	log.Println("encrypt message:%s", s)

	prikey := "MC0CAQACBQDZBgQRAgMBAAECBDSXSAkCAwD+8wIDANnrAgMA+U0CAkc5AgMAlq8="
	s, err = DecodeWithPrivKey(prikey, "This a encrypt msessage")
	if err != nil {
		log.Println("couldn't decode the secret message \n %s", err)
	}
	log.Println("decode message:%s", s)

	log.Println("---------finish----------")
}

// EncryptToPubKey takes a NaCL public key in Base64-encoded form, and
// a message, and returns a Base64-encoded message encrypted for that
// key, or an error.
func EncryptToPubKey(pubkey, message string) (s string, err error) {
	var (
		b                []byte
		recipientKey     [32]byte
		nonce            [24]byte
		ephemeralPrivKey *[32]byte
		ephemeralPubKey  *[32]byte
	)
	// decode the public key
	if b, err = base64.StdEncoding.DecodeString(pubkey); err != nil {
		return "", errors.Wrap(err, "couldn't decode public key")
	}
	// if it's not exactly 32 bytes, it can't be a public key
	if len(b) != 32 {
		log.Println("DEBUG:this secrets length:%s", len(b))
		log.Println("DEBUG:this secrets is :%S", b)
		return "", errors.New("secrets must be 32 bytes long")
	}
	// turn the byte slice into a byte array
	copy(recipientKey[:], b[0:32])
	// generate the ephemeral keypair
	if ephemeralPubKey, ephemeralPrivKey, err = box.GenerateKey(rand.Reader); err != nil {
		return "", errors.Wrap(err, "couldn't generate ephemeral key")
	}
	// read a random nonce
	if _, err = io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return "", errors.Wrap(err, "couldn't read random nonce")
	}
	// encrypt the message, and append to the nonce
	b = box.Seal(nonce[:], []byte(message), &nonce, &recipientKey,
		ephemeralPrivKey)
	// prepend the ephemeral public key
	b = append((*ephemeralPubKey)[:], b...)
	// finally, return the encoded, encrypted message
	return base64.StdEncoding.EncodeToString(b), nil
}

func DecodeWithPrivKey(privKey, message string) (m string, err error) {
	var (
		b               []byte
		key             [32]byte
		nonce           [24]byte
		ephemeralPubKey [32]byte
		ok              bool
	)
	// decode the private key
	if b, err = base64.StdEncoding.DecodeString(privKey); err != nil {
		panic(err)
	}
	// if it's not exactly 32 bytes, it can't be a public key
	if len(b) != 32 {
		return "", errors.New("secrets must be 32 bytes long")
	}
	// turn the byte slice into a byte array
	copy(key[:], b[0:32])
	// decode the ciphertext
	if b, err = base64.StdEncoding.DecodeString(message); err != nil {
		return "", errors.Wrap(err, "couldn't decode encrypted message")
	}
	// extract the ephemeral public key
	if len(b) < 32 {
		return "", errors.New("missing ephemeral public key")
	}
	copy(ephemeralPubKey[:], b[:32])
	// extract the nonce
	if len(b) < 56 {
		return "", errors.New("missing nonce")
	}
	copy(nonce[:], b[32:56])
	// decrypt the message
	b, ok = box.Open(nil, b[56:], &nonce, &ephemeralPubKey, &key)
	if !ok {
		return "", errors.New("couldn't decrypt message")
	}
	return string(b), nil
}
