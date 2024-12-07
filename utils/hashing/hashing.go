package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))

	bs := hex.EncodeToString(h.Sum(nil))
	return bs
}
