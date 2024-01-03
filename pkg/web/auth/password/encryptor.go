package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/lowl11/boost/pkg/system/types"
	"io"
)

// Encryptor is a service which can encrypt or decrypt given password (or any string)
type Encryptor struct {
	key []byte
}

// NewEncryptor creates Encryptor instance.
// Argument "key" is key for AES encryption.
// Key should have 16, 24 or 36 bytes length
func NewEncryptor(key string) *Encryptor {
	return &Encryptor{
		key: types.ToBytes(key),
	}
}

// Encrypt encrypts given password (or any string) by using encryptors key
func (encrypt Encryptor) Encrypt(password string) (string, error) {
	plaintext := types.ToBytes(password)

	block, err := aes.NewCipher(encrypt.key)
	if err != nil {
		return "", ErrorEncryptPassword(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts given password (or any string) by using encryptors key
func (encrypt Encryptor) Decrypt(password string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(password)

	block, err := aes.NewCipher(encrypt.key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		return "", ErrorDecryptPassword(err)
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return types.ToString(ciphertext), nil
}
