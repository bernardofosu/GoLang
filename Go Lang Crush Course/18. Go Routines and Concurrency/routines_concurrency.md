# 🧵 What Is a Goroutine?

A **goroutine** is a lightweight thread managed by the Go runtime.

✅ In simple terms:
> A goroutine is a way to run a function at the same time as other functions — **concurrently**.

---

## 🚀 How to Start a Goroutine

You just put the keyword `go` in front of a function call:

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	go sayHello() // 🔄 runs in the background

	fmt.Println("Main function done.")

	time.Sleep(1 * time.Second) // ⏳ give goroutine time to finish
}
```

🧠 **Key points**:
- `go sayHello()` runs `sayHello()` in a separate lightweight thread.
- `time.Sleep(...)` is used so the program doesn’t end before the goroutine runs.
- If you remove `time.Sleep`, the program might exit before the goroutine has a chance to execute.

---

## 🔄 Concurrency vs Parallelism

| Concept      | Description |
|--------------|-------------|
| **Concurrency** | Doing multiple tasks at once by switching between them. Goroutines use this model. |
| **Parallelism** | Doing multiple tasks at once on multiple CPUs. Go can do this too, but concurrency ≠ parallelism. |

💡 Think of **concurrency** as a good multitasker.  
**Parallelism** is like many people doing tasks at the same time.

---

## 🎯 Real Example: Multiple Goroutines

```go
package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello,", name)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go greet("Alice")
	go greet("Bob")

	time.Sleep(2 * time.Second) // Wait for goroutines to finish
	fmt.Println("Main done!")
}
```

---

## 💌 Communicating Between Goroutines: Channels

```go
func main() {
	ch := make(chan string) // 📨 channel of strings

	go func() {
		ch <- "Hello from Goroutine!" // send message into channel
	}()

	msg := <-ch // 🛑 wait and receive message from channel
	fmt.Println(msg)
}
```

---

## 📌 Summary

| Term        | What it means |
|-------------|----------------|
| `go`        | Starts a goroutine (concurrent task) |
| **Goroutine** | Lightweight thread managed by Go runtime |
| **Channel**   | Tool to communicate between goroutines |
| **Concurrency** | Tasks that run overlapping in time |
| **Parallelism** | Tasks that run exactly at the same time |