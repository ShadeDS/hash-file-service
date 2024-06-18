package util

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/heroiclabs/nakama-common/runtime"
)

// CalculateHash calculates the SHA-256 hash of the input data and returns it as a hexadecimal string.
func CalculateHash(data []byte, logger runtime.Logger) string {
	logger.Debug("Calculating SHA-256 hash")
	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])
	logger.Debug("Hash calculated successfully", "hash", hashString)
	return hashString
}
