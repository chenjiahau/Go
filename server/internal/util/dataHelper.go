package util

import "time"

func GetNow() time.Time {
	return time.Now().UTC()
}