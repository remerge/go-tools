package dm

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func DecryptHmacXorWithIntegrity(message, encryptKey, integrityKey []byte) ([]byte, error) {
	size := len(message)
	initializationVector := message[0:16]
	cipherText := message[16 : size-4]
	cipherTextSize := len(cipherText)
	integritySignature := message[size-4 : size]

	mac := hmac.New(sha1.New, encryptKey)
	if _, err := mac.Write(initializationVector); err != nil {
		return nil, err
	}

	pad := mac.Sum(nil)
	if len(pad) < cipherTextSize {
		return nil, fmt.Errorf("pad length to short, pad: %v, pad-size: %d, message: %v, cipher-text-size: %d", pad, len(pad), message, len(cipherText))
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

	signature := mac.Sum(nil)[0:4]

	// check signature
	for i := 0; i < 4; i++ {
		if signature[i] != integritySignature[i] {
			return nil, fmt.Errorf("signature %#v does not match integrity %#v", signature, integritySignature)
		}
	}

	return unciphered, nil
}
