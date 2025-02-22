package helpers

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

var angka = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateRekening(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = angka[int(b[i])%len(angka)]
	}
	return string(b)
}

func EncryptedHash(plaintext string) string {
	passwordHash := sha1.New()
	passwordHash.Write([]byte(plaintext))
	passwordSha1 := hex.EncodeToString(passwordHash.Sum(nil))
	return passwordSha1

}
