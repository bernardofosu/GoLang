# 🧵 Go Concurrency Model: M:N Scheduler Explained

This guide visualizes how **Go's concurrency model** works using an M:N scheduler — where **many goroutines (G)** are mapped to **a smaller set of OS threads (M)** via **logical processors (P)**.

---

## 🟨 Goroutines (G)

- Lightweight, user-space threads created using `go funcName()`.
- Managed by **Go's runtime**, **not the OS**.
- Can scale to **millions** — efficient and low overhead.
- Queued and scheduled onto **Logical Processors (P)**.

---

## 🟩 Logical Processors (P)

- Internal schedulers within Go.
- Each `P` maintains a **run queue of goroutines**.
- Count is determined by `GOMAXPROCS`.
- A `P` is **bound to an OS thread (M)** to execute goroutines.

---

## 🟦 OS Threads (M)

- Real operating system threads.
- `P` uses `M` to run goroutines.
- If a goroutine blocks (e.g., on I/O), **the `M` may be parked**.
- **New `M`s are created** to keep `P`s working without delay.

---

## 🔁 Execution Flow Diagram

```text
    G (Goroutines)
     ↓   ↓   ↓   ↓   ↓
    +----+----+----+----+
    | Run Queue of P₁   |       P = Logical Processor
    +-------------------+       M = OS Thread
             ↓
         +--------+
         |   P₁   |────────┐
         +--------+        │
             ↓             ▼
         +--------+     +--------+
         |   M₁   | --> |  CPU 1 |
         +--------+     +--------+

If M₁ is blocked, P₁ is reassigned to a new M (e.g., M₂) to continue execution.
```

📌 **Key Insight**  
GOMAXPROCS defines how many goroutines run concurrently.  
Many more goroutines can exist but are queued for execution.  
The Go scheduler handles this transparently and efficiently.

---

## 🧠 Analogy: A Restaurant Kitchen

| Concept    | Analogy                  |
|------------|---------------------------|
| Goroutines | Orders from customers     |
| Ps         | Cooks handling task queues|
| Ms         | Stove burners             |
| Scheduler  | Head chef coordinating work|

Each cook (P) works on a queue of orders (G) and uses a burner (M) to prepare dishes.  
If a burner gets occupied (e.g., I/O wait), the chef finds another burner to keep cooking going.

✅ Efficient. Scalable. Non-blocking. That’s Go's concurrency power.

---

## 🧩 Clarifications About CPUs and Threads

🔍 **System Example**  
Your system has:

- 4 physical CPU cores
- 8 logical processors (via Hyper-Threading: 2 threads/core)

🧠 **In Go:**

| Component            | Role                                 |
|----------------------|--------------------------------------|
| Goroutine (G)        | The task/function to run             |
| Logical Processor (P)| Manages and queues goroutines        |
| OS Thread (M)        | Executes the work assigned by P      |
| CPU Core / Thread    | Hardware unit running the OS thread M|

A goroutine (G) runs on a P, which runs on an M, which the OS schedules on a logical CPU.

---

## ❓ Does GOMAXPROCS(n) Use Logical or Physical Cores?

✅ **Answer:**  
GOMAXPROCS(n) tells Go how many Ps to create. Each P runs on an M (OS thread), and the OS schedules M on any available hardware thread (logical processor) — not limited to physical cores.

🔁 **Behavior:**

- `runtime.GOMAXPROCS(8)` = Go uses 8 Ps
- Each P gets an M
- OS schedules M on 8 logical processors (hyperthreads)

🛠 You can determine your available logical processors in Go using:

```go
runtime.NumCPU() // returns 8 on your system
```

---

## ✅ Summary

- Goroutines (G) are cheap and many  
- Logical Processors (P) are Go's internal schedulers (set by GOMAXPROCS)  
- OS Threads (M) are mapped to CPU logical cores by the OS  
- Go handles blocking and scaling efficiently  
