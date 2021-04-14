package utils

import (
	"crypto/sha256"
	"fmt"
)

func Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%X", h.Sum(nil))
}
