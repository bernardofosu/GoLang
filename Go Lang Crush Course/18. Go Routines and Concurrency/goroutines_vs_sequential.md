# Comparing Sequential Execution vs Goroutines in Go

## 🧪 Objective

Demonstrate the difference between:
- Sequential CPU-bound task execution (no goroutines)
- Concurrent execution using goroutines and `sync.WaitGroup`

## 💡 System Info

- Logical CPUs: 8
- GOMAXPROCS: Set to match CPU count (`runtime.GOMAXPROCS(runtime.NumCPU())`)

---

## 🧱 Version 1: **Sequential Execution (No Goroutines)**

### ✅ Code

```go
package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	numCPUs := runtime.NumCPU()
	log.Printf("Number of CPUs available: %d\n", numCPUs)

	runtime.GOMAXPROCS(numCPUs)
	log.Printf("GOMAXPROCS set to: %d\n", numCPUs)

	start := time.Now()

	numTasks := 10
	for i := 0; i < numTasks; i++ {
		log.Printf("Task %d started\n", i)
		for j := 0; j < 10000000; j++ {
		}
		log.Printf("Task %d finished\n", i)
	}

	duration := time.Since(start)
	log.Printf("Total time taken: %s\n", duration)
}
```

### 🖨️ Sample Output

```
2025/05/08 09:48:10 Number of CPUs available: 8
2025/05/08 09:48:10 GOMAXPROCS set to: 8
2025/05/08 09:48:10 Task 0 started
...
2025/05/08 09:48:10 Task 9 finished
2025/05/08 09:48:10 Total time taken: 110.452ms
```

---

## ⚙️ Version 2: **Concurrent Execution with Goroutines + WaitGroup**

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
	numCPUs := runtime.NumCPU()
	log.Printf("Number of CPUs available: %d\n", numCPUs)

	runtime.GOMAXPROCS(numCPUs)
	log.Printf("GOMAXPROCS set to: %d\n", numCPUs)

	start := time.Now()

	var wg sync.WaitGroup
	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			log.Printf("Goroutine %d started\n", id)
			for j := 0; j < 10000000; j++ {
			}
			log.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()

	duration := time.Since(start)
	log.Printf("Total time taken: %s\n", duration)
}
```

### 🖨️ Sample Output

```
2025/05/08 09:51:13 Number of CPUs available: 8
2025/05/08 09:51:13 GOMAXPROCS set to: 8
2025/05/08 09:51:13 Goroutine 3 started
...
2025/05/08 09:51:13 Goroutine 0 finished
2025/05/08 09:51:13 Total time taken: 14.4282ms
```

---

## 📊 Summary Table

| Feature                     | Sequential              | Goroutines + WaitGroup     |
|----------------------------|-------------------------|----------------------------|
| Execution model            | Single-threaded         | Multi-threaded             |
| Concurrency                | ❌ No                   | ✅ Yes                     |
| Parallelism                | ❌ Only 1 core used      | ✅ Up to 8 cores used       |
| Simulated task duration    | ~110ms                  | ~14ms                      |
| Synchronization            | Not needed              | Required (`sync.WaitGroup`)|
| CPU utilization            | Low                     | High                       |

---

## 🧠 Conclusion

- Goroutines significantly reduce execution time when multiple CPU cores are available.
- Parallelism improves throughput for CPU-bound operations.
- `sync.WaitGroup` is essential for coordinated completion.

---
