package pkg

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateUniqueID(str1, str2 string) string {
	// Concatenate the strings
	combinedStr := str1 + str2

	// Generate SHA-256 hash
	hash := sha256.New()
	hash.Write([]byte(combinedStr))

	// Convert the hash to a hex string
	hashedStr := hex.EncodeToString(hash.Sum(nil))

	return hashedStr
}
