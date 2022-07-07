package xeqtr

import (
	"fmt"
	"runtime"
	"sync"
)

type xeqtr_core struct {
	Capacity int
}

var initialised_xeqtr_core *xeqtr_core
var once sync.Once

func Create(capacity int) (*xeqtr_core, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be positive and greater than 0")
	}

	if capacity > runtime.NumCPU() {
		fmt.Printf("WARN: capacity %d is greater than CPU's, you might experience a performance issue \n", capacity)
	}

	once.Do(func() {
		initialised_xeqtr_core = &xeqtr_core{capacity: capacity}
	})
	fmt.Printf("capacity: %d\n", capacity)
	return initialised_xeqtr_core, nil
}

func GetCapacity() (int, error) {
	if initialised_xeqtr_core == nil {
		return 0, fmt.Errorf("unable to get capacity")
	}
	return initialised_xeqtr_core.capacity, nil
}
