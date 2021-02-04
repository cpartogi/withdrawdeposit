package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	b64 "encoding/base64"
	"os"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

//ToSnakeCase is function to convert camelCase to snake_case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

//SetLowerAndAddSpace is function to convert camelCase to snake_case
func SetLowerAndAddSpace(str string) string {
	lower := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	lower = matchAllCap.ReplaceAllString(lower, "${1} ${2}")
	return strings.ToLower(lower)
}

// GetEnv returns app envorinment : e.g. development, production, staging, testing, etc
func GetEnv() string {
	return os.Getenv("APP_ENV")
}

// IsProductionEnv returns whether the app is running using production env
func IsProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}

//HashPassword will generate hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

//CheckPasswordHash will
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

// Encrypt text with the passphrase
func Encrypt(text string, passphrase string) string {
	salt := []byte("g7gHa3LQe6Txg6zpypft") // salt must 20 chars

	key, iv := deriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	pad := pKCS7Padding([]byte(text), block.BlockSize())
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(pad))
	ecb.CryptBlocks(encrypted, pad)

	return b64.StdEncoding.EncodeToString([]byte("Salted__" + string(salt) + string(encrypted)))
}

// Decrypt encrypted text with the passphrase
func Decrypt(encrypted string, passphrase string) string {
	ct, _ := b64.StdEncoding.DecodeString(encrypted)
	if len(ct) < 28 || string(ct[:8]) != "Salted__" {
		return ""
	}

	salt := ct[8:28]
	ct = ct[28:]
	key, iv := deriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	dst := make([]byte, len(ct))
	cbc.CryptBlocks(dst, ct)

	return string(pKCS7Trimming(dst))
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func deriveKeyAndIv(passphrase string, salt string) (string, string) {
	salted := ""
	dI := ""

	for len(salted) < 48 {
		md := md5.New()
		md.Write([]byte(dI + passphrase + salt))
		dM := md.Sum(nil)
		dI = string(dM[:16])
		salted = salted + dI
	}

	key := salted[0:32]
	iv := salted[32:48]

	return key, iv
}
