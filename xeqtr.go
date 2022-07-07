package xeqtr

import (
	"fmt"
	"runtime"
	"sync"
)

type xeqtr_core struct {
	capacity int
}

const max_multiplier int = 4

var default_xeqtr_core_capacity int = runtime.NumCPU()
var max_xeqtr_core_capacity_limit int = max_multiplier * default_xeqtr_core_capacity
var initialised_xeqtr_core *xeqtr_core = nil
var once sync.Once

func Create(capacity int) error {
	if capacity <= 0 {
		fmt.Printf("WARN: Capacity cannot be negative or zero \n")
		fmt.Printf("Setting capacity to %d\n", default_xeqtr_core_capacity)
	}

	if capacity > default_xeqtr_core_capacity {
		if capacity <= max_xeqtr_core_capacity_limit {
			fmt.Printf(
				"WARN: current capacity : %d is greater than toatl CPU's - %d , you might experience a performance issue \n",
				capacity,
				default_xeqtr_core_capacity,
			)
		} else {
			return fmt.Errorf(
				"capacity limit exceeded, please set the capacity limit less than or equal to max capacity : %d",
				max_xeqtr_core_capacity_limit,
			)
		}

	}

	once.Do(func() {
		initialised_xeqtr_core = &xeqtr_core{capacity: capacity}
	})
	fmt.Printf("capacity: %d\n", capacity)
	return nil
}

func GetCapacity() (int, error) {
	if initialised_xeqtr_core == nil {
		return -1, fmt.Errorf("core unintialised")
	}
	return initialised_xeqtr_core.capacity, nil
}
