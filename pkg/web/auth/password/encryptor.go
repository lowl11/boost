package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/lowl11/boost/pkg/system/types"
	"io"
)

type Encryptor struct {
	key []byte
}

func NewEncryptor(key string) *Encryptor {
	return &Encryptor{
		key: types.ToBytes(key),
	}
}

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
