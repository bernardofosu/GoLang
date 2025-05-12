🧠 **Understanding Concurrency in Go (Beginner Explanation)**

✅ **1. Go runs sequentially by default**  
When you write a Go program, it normally runs line by line — one thing at a time.

```go
fmt.Println("First")
fmt.Println("Second")
```
The second line runs only after the first one finishes. That’s sequential.

---

✅ **2. Go supports concurrency using the `go` keyword**  
When you write:

```go
go someFunction()
```

You’re telling Go:  
“Run this function in the background as a goroutine (lightweight thread), and continue with the next lines.”

**BUT** — and this is the key part — the main program doesn’t wait for these background goroutines to finish unless you tell it to wait.

🚨 **Problem: Goroutines might still be running when `main()` finishes**

```go
func main() {
    go fmt.Println("Hello from goroutine!")
    fmt.Println("Main done!")
}
```

You might not see the goroutine's message, because the program could exit before the goroutine even gets a chance to run.

---

✅ **3. Solution: Use the `sync` package to make the main thread wait**  
The most common tool is `sync.WaitGroup`. It lets the main function say:  
“Hey, I’m going to wait here until all the goroutines are done.”

✅ **Example:**

```go
package main

import (
    "fmt"
    "sync"
)

func printMessage(wg *sync.WaitGroup) {
    defer wg.Done() // Mark as done when finished
    fmt.Println("Hello from goroutine!")
}

func main() {
    var wg sync.WaitGroup

    wg.Add(1) // We're waiting for 1 goroutine

    go printMessage(&wg)

    wg.Wait() // Wait here until the goroutine is done
    fmt.Println("Main function done!")
}
```

---

🧠 **Final Summary for Beginners**

✅ Go runs code line-by-line by default.  
✅ You can run things at the same time using `go someFunction()`.  
⚠️ But Go won’t wait for background work to finish unless you tell it to.  
✅ Use `sync.WaitGroup` to tell Go: “Wait for all the goroutines before exiting.”

---

🧰 **What is `sync` in Go?**  
The `sync` package provides basic synchronization tools to help manage concurrent access to shared resources, avoid race conditions, and coordinate multiple goroutines.

---

🔧 **Most Common `sync` Tools (with Examples)**

### 1. `sync.WaitGroup` – 🔄 Wait for Goroutines to Finish  
Used to wait for multiple goroutines to complete their tasks.

✅ **Example:**

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d is done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers finished.")
}
```

🧠 **Key Methods:**  
- `Add(n int)` – Increase the number of goroutines to wait for.  
- `Done()` – Decrease the count by one.  
- `Wait()` – Block until the count becomes zero.

---

### 2. `sync.Mutex` – 🔐 Locking to Protect Shared Data  
Used to make sure only one goroutine accesses a shared resource at a time.

✅ **Example:**

```go
package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
    defer wg.Done()

    mutex.Lock()          // Lock before accessing shared resource
    counter++
    mutex.Unlock()        // Always unlock after you're done
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
```

🧠 **Key Methods:**  
- `Lock()` – Locks the mutex so no one else can enter the critical section.  
- `Unlock()` – Releases the lock so other goroutines can continue.

---

### 3. `sync.RWMutex` – 🧩 Read-Write Lock  
Like Mutex, but allows multiple readers at the same time (read locks), or only one writer (write lock).

✅ **Example:**

```go
package main

import (
    "fmt"
    "sync"
)

var (
    data  = make(map[string]string)
    rwMux sync.RWMutex
)

func write(key, value string, wg *sync.WaitGroup) {
    defer wg.Done()
    rwMux.Lock()
    data[key] = value
    rwMux.Unlock()
}

func read(key string, wg *sync.WaitGroup) {
    defer wg.Done()
    rwMux.RLock()
    fmt.Println("Read:", key, "=", data[key])
    rwMux.RUnlock()
}

func main() {
    var wg sync.WaitGroup

    wg.Add(2)
    go write("language", "Go", &wg)
    go read("language", &wg)

    wg.Wait()
}
```

🧠 **Key Methods:**  
- `Lock()` and `Unlock()` for write operations.  
- `RLock()` and `RUnlock()` for read operations (multiple reads allowed).

---

### 4. `sync.Once` – ☝️ Run Something Only Once  
Ensures a piece of code runs only once, no matter how many times it's called.

✅ **Example:**

```go
package main

import (
    "fmt"
    "sync"
)

var once sync.Once

func initApp() {
    fmt.Println("App initialized")
}

func main() {
    for i := 0; i < 3; i++ {
        go func() {
            once.Do(initApp) // Only this will run once
        }()
    }

    // Wait so goroutines finish (not best practice, just quick demo)
    var input string
    fmt.Scanln(&input)
}
```

🧠 **Key Method:**  
- `Do(func)` – Executes the function exactly once, no matter how many goroutines call it.

---

### 5. `sync.Map` – 🗺️ Concurrent-Safe Map  
Unlike normal maps, `sync.Map` is safe for concurrent use and doesn’t need you to use a mutex.

✅ **Example:**

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map

    m.Store("language", "Go")           // Add or update
    m.Store("version", "1.20")

    lang, _ := m.Load("language")       // Read
    fmt.Println("Language:", lang)

    m.Range(func(key, value interface{}) bool {
        fmt.Println(key, "=", value)
        return true // continue looping
    })

    m.Delete("version")                 // Remove key
}
```

🧠 **Key Methods:**  
- `Store(key, value)` – Add or update  
- `Load(key)` – Get a value  
- `Delete(key)` – Remove a value  
- `Range(func)` – Loop through the map

---

🧠 **Summary Table**

| Tool        | Purpose                         | Key Methods                          |
|-------------|----------------------------------|--------------------------------------|
| WaitGroup   | Wait for goroutines to finish   | `Add()`, `Done()`, `Wait()`          |
| Mutex       | Lock/unlock critical section    | `Lock()`, `Unlock()`                 |
| RWMutex     | Read/write locking              | `RLock()`, `RUnlock()`, `Lock()`     |
| Once        | Run code only once              | `Do(func)`                           |
| Map         | Thread-safe map                 | `Store()`, `Load()`, `Delete()`, `Range()` |
