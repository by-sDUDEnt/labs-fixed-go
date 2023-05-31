package tokens

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (Impl) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}

func (Impl) PasswordMatches(password string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s Impl) CreateToken() (string, []byte, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", nil, fmt.Errorf("generate random bytes: %w", err)
	}

	plaintext := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := s.HashToken(plaintext)

	return plaintext, hash, nil
}

func (Impl) HashToken(plaintext string) []byte {
	hash := sha256.Sum256([]byte(plaintext))
	return hash[:]
}
