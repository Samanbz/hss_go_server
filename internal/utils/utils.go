package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func Hash(pwd string) string {
	hashedPwd := sha256.Sum256([]byte(pwd))
	hashedPwdStr := hex.EncodeToString(hashedPwd[:])
	return hashedPwdStr
}

func InNDaysAt(n int, hour int, minute int) time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+n, hour, minute, 0, 0, time.Local)
}
