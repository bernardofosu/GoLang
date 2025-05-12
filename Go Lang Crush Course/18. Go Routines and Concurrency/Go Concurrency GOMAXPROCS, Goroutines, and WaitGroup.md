# Go Concurrency: GOMAXPROCS, Goroutines, and WaitGroup

## 📌 Overview

Go’s concurrency model uses goroutines — lightweight threads managed by the Go runtime. `GOMAXPROCS` determines how many OS threads can execute Go code simultaneously.

This document explores performance comparisons and behavior with and without goroutines, and with `sync.WaitGroup`.

---

## ⚙️ Example 1: Using Goroutines with WaitGroup

### ✅ Code

```go
package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Log the number of CPUs available
	numCPUs := runtime.NumCPU()
	log.Printf("Number of CPUs available: %d\n", numCPUs)

	// Set GOMAXPROCS based on the number of CPUs
	runtime.GOMAXPROCS(numCPUs)
	log.Printf("GOMAXPROCS set to: %d\n", numCPUs)

	// Start the timer
	start := time.Now()

	// Use WaitGroup to wait for all goroutines
	var wg sync.WaitGroup
	numGoroutines := 10
	wg.Add(numGoroutines)

	// Launch multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			log.Printf("Goroutine %d started\n", id)
			for j := 0; j < 10000000; j++ {
			}
			log.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Measure and log time taken
	duration := time.Since(start)
	log.Printf("Total time taken: %s\n", duration)
}
```

### 🧾 Output Example

```
2025/05/08 09:51:13 Goroutine 3 finished
2025/05/08 09:51:13 Goroutine 0 started
2025/05/08 09:51:13 Goroutine 4 finished
2025/05/08 09:51:13 Goroutine 1 finished
2025/05/08 09:51:13 Goroutine 2 finished
2025/05/08 09:51:13 Goroutine 7 finished
2025/05/08 09:51:13 Goroutine 8 finished
2025/05/08 09:51:13 Goroutine 5 finished
2025/05/08 09:51:13 Goroutine 0 finished
2025/05/08 09:51:13 Total time taken: 14.4282ms
```

---

## 🔄 How `sync.WaitGroup` Works

### 🧠 Conceptual Diagram

```text
WaitGroup Counter: 0
      |
wg.Add(3) --> Counter: 3  ──┬── Goroutine A
                            ├── Goroutine B
                            └── Goroutine C
                                   |
                          Each calls wg.Done() (↓ counter)
                                   |
           wg.Done()  -> Counter: 2
           wg.Done()  -> Counter: 1
           wg.Done()  -> Counter: 0
                                   |
                              wg.Wait() unblocks
```

### 📌 Summary:

* ➕ `wg.Add(n)` increases the counter by `n`.
* ➖ `wg.Done()` decreases the counter by 1.
* ⏳ `wg.Wait()` blocks until the counter reaches 0.

---

Would you like to add a comparison with the version that doesn’t use goroutines next?
