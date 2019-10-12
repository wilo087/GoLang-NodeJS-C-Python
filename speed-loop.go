package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i < 1000000000; i++ {
		// Empty lopp
	}
	end := time.Now()

	elapsed := end.Sub(start)

	output := fmt.Sprintf("Go looped 1,000,000,000 time in %s ", elapsed)
	fmt.Println(output)
}
