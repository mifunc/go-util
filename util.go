package util

import (
	"time"
)

func Tool() int {
	return time.Now().Year()
}

func GetMonth() int {
	return int(time.Now().Month())
}
