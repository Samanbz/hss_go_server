package security

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(pwd string) string {
	hashedPwd := sha256.Sum256([]byte(pwd))
	hashedPwdStr := hex.EncodeToString(hashedPwd[:])
	return hashedPwdStr
}
