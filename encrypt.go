package mutt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"io"
	"errors"
)

var (
	ERR_ENC_CIPHER_SHORT = errors.New("encryption cipher too short")
)

func PasswdHash(pwd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
}

func PasswdCompare(hashedPaswd []byte, plainPasswd []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPaswd, plainPasswd)
}

// TODO: encryption needed
func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, ERR_ENC_CIPHER_SHORT
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
func ByteToB64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func B64ToByte(b string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(b)
}
