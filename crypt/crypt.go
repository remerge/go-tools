package crypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"hash"
	"math/rand"
)

const (
	initializationVectorSize = 16
	integritySignatureSize   = 4
	minimalMessageSize       = initializationVectorSize + integritySignatureSize
)

func DecryptHmacXorWithIntegrity(message, encryptKey, integrityKey []byte) ([]byte, error) {
	size := len(message)
	if size < minimalMessageSize {
		return nil, fmt.Errorf("message length too short(<%d), message: %v, length: %d", minimalMessageSize, message, len(message))
	}

	initializationVector := message[0:initializationVectorSize]
	cipherText := message[initializationVectorSize : size-integritySignatureSize]
	cipherTextSize := len(cipherText)
	integritySignature := message[size-integritySignatureSize : size]

	mac := hmac.New(sha1.New, encryptKey)
	if _, err := mac.Write(initializationVector); err != nil {
		return nil, err
	}

	pad := mac.Sum(nil)
	if len(pad) < cipherTextSize {
		return nil, fmt.Errorf("pad length too short, pad: %v, pad-length: %d, message: %v, cipher-text-length: %d", pad, len(pad), message, len(cipherText))
	}

	unciphered := make([]byte, cipherTextSize)
	for i := 0; i < cipherTextSize; i++ {
		unciphered[i] = pad[i] ^ cipherText[i]
	}

	mac = hmac.New(sha1.New, integrityKey)
	if _, err := mac.Write(unciphered); err != nil {
		return nil, err
	}

	if _, err := mac.Write(initializationVector); err != nil {
		return nil, err
	}

	signature := mac.Sum(nil)[0:integritySignatureSize]

	// check signature
	for i := 0; i < integritySignatureSize; i++ {
		if signature[i] != integritySignature[i] {
			return nil, fmt.Errorf("signature %#v does not match integrity %#v", signature, integritySignature)
		}
	}

	return unciphered, nil
}

func CryptHmacXorWithIntegrity(message, encryptKey, integrityKey []byte) ([]byte, error) {
	return cryptHmacXorWithIntegrity(message, sha1.New, encryptKey, integrityKey)
}

func cryptHmacXorWithIntegrity(message []byte, h func() hash.Hash, encryptKey []byte, integrityKey []byte) ([]byte, error) {
	var err error

	initializationVector := make([]byte, initializationVectorSize)
	for i := 0; i < initializationVectorSize; i++ {
		initializationVector[i] = uint8(rand.Int() % 255)
	}

	// create cipher text
	padMac := hmac.New(h, encryptKey)
	_, err = padMac.Write(initializationVector)
	if err != nil {
		return nil, err
	}

	pad := padMac.Sum(nil)
	if len(pad) < len(message) {
		return nil, fmt.Errorf("pad length to short, pad: %v, pad-size: %d, message: %v, message-text-size: %d", pad, len(pad), message, len(message))
	}

	ciphered := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		ciphered[i] = pad[i] ^ message[i]
	}

	// create signature
	sigMac := hmac.New(h, integrityKey)
	_, err = sigMac.Write(message)
	if err != nil {
		return nil, err
	}

	_, err = sigMac.Write(initializationVector)
	if err != nil {
		return nil, err
	}

	signature := sigMac.Sum(nil)[0:integritySignatureSize]

	// create result
	var result []byte
	result = append(result, initializationVector...)
	result = append(result, ciphered...)
	result = append(result, signature...)

	return result, nil
}
