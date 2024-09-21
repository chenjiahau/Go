package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func CreateMD5Hash(userId int64) string {
	userIdAndTime := strconv.FormatInt(userId, 10) + time.Now().String()
	hash := md5.Sum([]byte([]byte(userIdAndTime)))
	return hex.EncodeToString(hash[:])
}