package xeqtr

import (
	"fmt"
)

func Create(capacity int) bool {
	if capacity <= 0 {
		return false
	}
	fmt.Printf("capacity: %d\n", capacity)
	return true
}
