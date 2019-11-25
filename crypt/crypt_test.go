package crypt

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecryptHmacXorWithIntegrity(t *testing.T) {
	var testEncryptionKeys = []byte{
		0xb0, 0x8c, 0x70, 0xcf, 0xbc,
		0xb0, 0xeb, 0x6c, 0xab, 0x7e,
		0x82, 0xc6, 0xb7, 0x5d, 0xa5,
		0x20, 0x72, 0xae, 0x62, 0xb2,
		0xbf, 0x4b, 0x99, 0x0b, 0xb8,
		0x0a, 0x48, 0xd8, 0x14, 0x1e,
		0xec, 0x07,
	}

	var testIntegrityKeys = []byte{
		0xbf, 0x77, 0xec, 0x55, 0xc3,
		0x01, 0x30, 0xc1, 0xd8, 0xcd,
		0x18, 0x62, 0xed, 0x2a, 0x4c,
		0xd2, 0xc7, 0x6a, 0xc3, 0x3b,
		0xc0, 0xc4, 0xce, 0x8a, 0x3d,
		0x3b, 0xbd, 0x3a, 0xd5, 0x68,
		0x77, 0x92,
	}

	t.Run("check pad lenght", func(t *testing.T) {
		t.Run("check encryption and decryption", func(t *testing.T) {
			bytes := []byte{0, 0, 0, 24, 255, 22, 1, 120}
			cryptMsg, err := CryptHmacXorWithIntegrity(bytes, testEncryptionKeys, testIntegrityKeys)
			require.NoError(t, err)

			result, err := DecryptHmacXorWithIntegrity(cryptMsg, testEncryptionKeys, testIntegrityKeys)
			require.NoError(t, err)
			assert.EqualValues(t, bytes, result)
		})

		t.Run("check encryption and decryption with 20 byte", func(t *testing.T) {
			bytes := []byte{0, 145, 0, 0, 0, 0, 0, 120, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 17} // 20 bytes
			require.Len(t, bytes, 20)

			cryptMsg, err := CryptHmacXorWithIntegrity(bytes, testEncryptionKeys, testIntegrityKeys)
			require.NoError(t, err)

			result, err := DecryptHmacXorWithIntegrity(cryptMsg, testEncryptionKeys, testIntegrityKeys)
			require.NoError(t, err)
			assert.EqualValues(t, bytes, result)
		})

		t.Run("check encryption and decryption with 21 byte", func(t *testing.T) {
			bytes := []byte{0, 0, 0, 234, 0, 0, 0, 120, 0, 0, 6, 0, 0, 0, 10, 0, 0, 0, 11, 0, 12} // 21 bytes
			require.Len(t, bytes, 21)

			// sha1 give a 20 byte pad --> can't build a cipher text longer than 20 bytes
			// sha256 give a 32 byte long pad --> can build it with a 21 byte array
			cryptMsg, err := cryptHmacXorWithIntegrity(bytes, sha256.New, testEncryptionKeys, testIntegrityKeys)
			require.NoError(t, err)

			_, err = DecryptHmacXorWithIntegrity(cryptMsg, testEncryptionKeys, testIntegrityKeys)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "pad length too short") // the error is not because of the wrong hash function
		})
	})

	t.Run("check message length", func(t *testing.T) {
		t.Run("message too short", func(t *testing.T) {
			cryptMsg := []byte{0, 0, 0, 24, 255, 22, 1, 120}
			require.True(t, len(cryptMsg) < minimalMessageSize)

			_, err := DecryptHmacXorWithIntegrity(cryptMsg, testEncryptionKeys, testIntegrityKeys)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "message length too short")
		})

		t.Run("minimal message length", func(t *testing.T) {
			bytes := make([]byte, 0)
			cryptMsg, err := CryptHmacXorWithIntegrity(bytes, testEncryptionKeys, testIntegrityKeys)
			require.NoError(t, err)
			require.Len(t, cryptMsg, minimalMessageSize)

			result, err := DecryptHmacXorWithIntegrity(cryptMsg, testEncryptionKeys, testIntegrityKeys)
			require.NoError(t, err)
			assert.Empty(t, result)
		})
	})

}
