package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func generateSalt() ([]byte, error) {
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    return salt, err
}

func HashPassword(password string) (string, error) {
    salt, err := generateSalt()
    if err != nil {
        return "", err
    }

    hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
    return base64.RawStdEncoding.EncodeToString(append(salt, hash...)), nil
}

func VerifyPassword(storedHash, password string) bool {
    decodedHash, err := base64.RawStdEncoding.DecodeString(storedHash)
    if err != nil {
        return false
    }

    salt := decodedHash[:16]
    hash := decodedHash[16:]
    passwordHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
    return base64.RawStdEncoding.EncodeToString(hash) == base64.RawStdEncoding.EncodeToString(passwordHash)
}
