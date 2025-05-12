# GOMAXPROCS in Go (Golang)

GOMAXPROCS is an environment variable or runtime setting in Go (Golang) that controls how many operating system threads can execute Go code simultaneously — in other words, it sets the maximum number of CPU cores that can be used by the Go scheduler for running goroutines in parallel.

## 🔍 Why GOMAXPROCS Matters:

Go uses a concurrent scheduler that maps lightweight goroutines onto real OS threads. GOMAXPROCS sets the upper limit on how many OS threads can be executing Go code at once. It directly affects performance, especially in CPU-bound programs.

## 🧠 Default Behavior:

```bash
In Go 1.5 and later, GOMAXPROCS defaults to the number of logical CPUs on the machine.
```

You can override it via environment variable or in code.

### ✅ How to Use

**Environment Variable:**

```bash
GOMAXPROCS=4 go run myprogram.go
```

**In Code:**

```go
import "runtime"

func main() {
    runtime.GOMAXPROCS(4) // Use 4 OS threads
}
```

## 🧪 Example:

If your machine has 8 cores, but you only want to use 2 cores to limit CPU usage:

```bash
GOMAXPROCS=2 ./myapp
```

The Go runtime will schedule goroutines across only 2 OS threads at a time.

## ⚠️ Notes:

- It does not limit the number of goroutines — only how many can run in parallel.
- It affects CPU-bound tasks more than I/O-bound tasks.

## 📊 Visual Example of Goroutines with Different GOMAXPROCS Values

| GOMAXPROCS | Expected Time (↓ faster)       |
|------------|--------------------------------|
| 1          | ⏱️ Slowest (1 core used)       |
| 2          | 🚀 Faster (2 cores)            |
| 4          | ⚡ Even faster                  |
| 8          | 🧨 Fastest (if CPU has 8+ cores)|

### 🧠 Interpretation:

The more CPU cores (via GOMAXPROCS) available, the more goroutines can run in parallel, reducing execution time.

After a certain point (e.g. GOMAXPROCS > physical cores), you’ll see diminishing returns or even overhead.

## 🧠 How Go Uses Threads and Cores:

Go has its own runtime scheduler, which is distinct from the operating system's thread scheduler. It doesn't run goroutines directly on OS threads in a 1:1 fashion. Instead, it uses a M:N scheduler:

- **M**: Number of goroutines (millions, potentially).
- **N**: Number of OS threads actually executing Go code, controlled by GOMAXPROCS.

## 🧩 Key Components:

| Component           | Description                                               |
|---------------------|-----------------------------------------------------------|
| Goroutines (G)      | Lightweight functions managed by Go, not the OS.          |
| OS Threads (M)      | Real threads managed by the OS.                           |
| Logical Processors (P) | A Go runtime abstraction. Each P runs 1 OS thread.       |

Go uses a work-stealing scheduler where goroutines are queued per-P and can be stolen by other Ps when idle.

## 🔁 The Execution Model:

1. You create thousands of goroutines.
2. Go assigns them to P logical processors.
3. Each P has its own run queue of goroutines.
4. Go maps each P to a real OS thread (M), up to GOMAXPROCS.
5. Goroutines are executed by those threads; blocked ones yield and others can take their place.

## 📌 Main Thread:

Even the `main()` function runs in a goroutine. It's just the first goroutine that kicks everything off.

## 🧠 Analogy:

Think of it like a restaurant:

- **Goroutines** = Orders 🍽️
- **P (Logical Processors)** = Cooks 👨‍🍳
- **M (OS Threads)** = Stove burners 🔥
- **Go scheduler** = Kitchen manager assigning cooks to burners 👩‍🍳

## 📊 Breakdown of the Diagram:

### 🟨 Goroutines (G)

These are lightweight user-space threads managed entirely by the Go runtime.

You can think of them as tasks or functions that want CPU time.

### 🟦 OS Threads (M)

These are real operating system threads (what C/C++ threads map to).

Go uses a small pool of these.

### 🟩 Logical Processors (P)

These are Go’s internal scheduler units.

Each P maintains a run queue of goroutines.

Go schedules execution by pairing each P with an M (OS thread).

The number of P is defined by GOMAXPROCS, and determines how many goroutines can be executing in parallel.

## 🔄 How Scheduling Works:

- Each P has a queue of goroutines.
- A P is bound to an M.
- The M picks goroutines from P’s queue and executes them.
- If a goroutine blocks (e.g., waiting on I/O), the M may be parked and a new M is spawned to keep the P working.
- Work stealing allows idle Ps to steal goroutines from others.

## 🧠 Summary of the Flow:

- Goroutines → scheduled on → Logical Processors (P)
- Logical Processors → run on → OS Threads (M)
- OS Threads → executed by → CPU cores

## ⚙️ GOMAXPROCS:

If `GOMAXPROCS=3`, Go creates 3 Ps → meaning 3 goroutines can run in parallel across 3 OS threads, assuming enough cores.