# 🏁 Go Program Entry Point: `main()`

## What Is package main Then?
- main is a special package name used only for programs that are meant to be executed (like CLI tools, web servers, etc).
- The Go compiler looks for a main package and its main() function to know where to start running the code.

Exactly — you’ve got the idea! In Go, execution starts from the `main` function in the `main` package. So:

✅ **If You Have Helper Functions:**  
They don’t run automatically — they must be called directly or indirectly from `main()`.

## 🔁 How It Works

### 🔹 Example:

```go
package main

import "fmt"

func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("Bernard")
}
```

- Here, `main()` is the entry point.  
- `greet("Bernard")` is called inside `main()`, so it runs.

## 🧠 Indirect Calls

Even if a function is not called directly in `main()`, it can still be executed if it's called by another function that `main()` calls.

### Example:

```go
func sayHello() {
    greet("Kwasi")
}

func main() {
    sayHello() // This calls greet indirectly
}
```

## 🔒 If You Don’t Call It — It Won’t Run

If you define a function and never call it (directly or indirectly from `main`), Go just compiles it — but it doesn’t do anything.
