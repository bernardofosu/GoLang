# 🚀 Go Concurrency Model: GOMAXPROCS and CPU Management

## ✅ **Automatically Set GOMAXPROCS Based on Available CPUs**
To automatically set `GOMAXPROCS` to match the number of CPUs available on your machine, you can use the `runtime` package in Go like this:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU) // Automatically set to match CPU count
	fmt.Printf("GOMAXPROCS set to %d\n", numCPU)
}
```

🛠️ Explanation:

- `runtime.NumCPU()` returns the number of logical CPUs usable by the current process.
- `runtime.GOMAXPROCS(n)` sets the maximum number of CPUs that can execute Go code simultaneously.

⚠️ Note: Since Go 1.5, GOMAXPROCS is automatically set to the number of available CPUs by default. So in most cases, you don’t need to set it manually unless you want to tune performance or for testing.

## ⚙️ Manually Set GOMAXPROCS

If you want to explicitly set GOMAXPROCS, you can do it in two ways:

### Option 1: Set in Go Code

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4) // 🧠 Manually set to 4 logical processors
	fmt.Printf("GOMAXPROCS is manually set to %d\n", runtime.GOMAXPROCS(0))
}
```

🎯 This forces the Go scheduler to use exactly 4 logical processors, regardless of how many CPUs your machine has.

### Option 2: Set via Environment Variable

```bash
GOMAXPROCS=4 go run main.go
```

🧪 This is useful for testing how your program behaves with different levels of concurrency without modifying the code.

## 🔄 Dynamically Set GOMAXPROCS Based on CPU Count

You can also dynamically set GOMAXPROCS to match the number of available CPUs at runtime:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
	fmt.Printf("GOMAXPROCS set to match CPU count: %d\n", n)
}
```

## 🧠 What Happens When You Set GOMAXPROCS Higher Than Available Logical Cores?

If you set GOMAXPROCS to a number higher than the available logical CPUs (hardware threads), the excess goroutines will queue and wait for an available logical processor (P) to execute them.

📊 Example:
Machine with 4 Logical CPUs (hyper-threading enabled):

- Available logical processors (hardware threads): 8.
- Set GOMAXPROCS(10):

  - 8 logical processors will be used immediately.
  - The remaining 2 goroutines will be queued and wait until one of the logical processors becomes free.

### 🔁 Execution Flow:
- The Go scheduler tries to execute up to 8 concurrent goroutines (since there are 8 logical processors).
- The excess goroutines (in this case, 2) will be put into the run queue.
- Once a logical processor becomes free (e.g., after a goroutine finishes), it picks up the next goroutine from the queue.

## ⚠️ Important to Remember:
- Increased GOMAXPROCS without enough CPUs means underutilization of available cores, which could lead to less efficient scheduling.
- Setting it too high could also lead to performance bottlenecks if the system is continually waiting for processors to become available.
