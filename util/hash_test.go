package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHash(t *testing.T) {
	logger := &MockLogger{}
	data := []byte("test data")
	expectedHash := "916f0027a575074ce72a331777c3478d6513f786a591bd892da1a577bf2335f9"
	hash := CalculateHash(data, logger)
	assert.Equal(t, expectedHash, hash)
}
