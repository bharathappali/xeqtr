package xeqtr

import (
	"fmt"
	"runtime"
	"sync"
)

type xeqtr_core struct {
	capacity int
}

var default_xeqtr_core_capacity int = runtime.NumCPU()
var initialised_xeqtr_core *xeqtr_core = nil
var once sync.Once

func Create(capacity int) error {
	if capacity <= 0 {
		fmt.Printf("WARN: Capacity cannot be negative or zero \n")
		fmt.Printf("Setting capacity to %d\n", default_xeqtr_core_capacity)
	}

	if capacity > default_xeqtr_core_capacity {
		fmt.Printf("WARN: capacity %d is greater than CPU's, you might experience a performance issue \n", capacity)
	}

	once.Do(func() {
		initialised_xeqtr_core = &xeqtr_core{capacity: capacity}
	})
	fmt.Printf("capacity: %d\n", capacity)
	return nil
}

func GetCapacity() (int, error) {
	if initialised_xeqtr_core == nil {
		return -1, fmt.Errorf("unable to get capacity")
	}
	return initialised_xeqtr_core.capacity, nil
}
