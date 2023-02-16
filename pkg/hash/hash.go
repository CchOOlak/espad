package hash

import (
	"encoding/base64"
	"encoding/hex"

	"crypto/sha256"
)

type HashGenerator interface {
	GetHash(string) string
}

type hashGen struct{}

func New() HashGenerator {
	return &hashGen{}
}

func (h hashGen) GetHash(input string) string {
	sha256Hash := sha256Of(input)
	hexHash := hexOf(input)

	return sha256Hash[:4] + hexHash[:2]
}

func hexOf(input string) string {
	return hex.EncodeToString([]byte(input))
}

func sha256Of(input string) string {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return base64.URLEncoding.EncodeToString(algorithm.Sum(nil))
}
