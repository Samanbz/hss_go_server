package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(pwd string) string {
	hashedPwd := sha256.Sum256([]byte(pwd))
	hashedPwdStr := hex.EncodeToString(hashedPwd[:])
	return hashedPwdStr
}
