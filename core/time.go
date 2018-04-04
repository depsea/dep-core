package core

import "time"

// GetNowTimestamp -
func GetNowTimestamp() int {
	return int(time.Now().Unix())
}
