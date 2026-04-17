package utils

import (
	"time"
)

// Возвращает время в RFC3339
func NowRFC3339() string {
	return time.Now().Format(time.RFC3339)
}
