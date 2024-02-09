package hard

import (
	"fmt"
	"runtime"
)

func Memory() float64 {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Memory Usage: %.2f MB\n", float64(mem.Alloc)/1024/1024)
	return float64(mem.Alloc) / 1024 / 1024
}

func MemoryFormat() string {
	return fmt.Sprintf("%.2f Mb", Memory())
}

func ActiveGoroutines() int {
	return runtime.NumGoroutine()
}
