# 📤 Go: fmt.Printf vs fmt.Sprintf

Absolutely, here's the cleaned-up and consistent version with full `var` declarations instead of `:=`, just like you asked 👇

## 🖨️ fmt.Printf

- **Purpose**: Prints formatted output directly to the console (stdout).
- **Does not return anything**.
- Think of it like `print()` in Python or `printf` in C.

```go
package main

import "fmt"

func main() {
    var name string = "Bernard"
    fmt.Printf("Hello, %s!\n", name)  // Output: Hello, Bernard!
}
```

---

## 📝 fmt.Sprintf

- **Purpose**: Returns the formatted string as a value, instead of printing it.
- **Useful when you want to store, log, or reuse the string.**

```go
package main

import "fmt"

func main() {
    var name string = "Bernard"
    var greeting string = fmt.Sprintf("Hello, %s!", name)
    fmt.Println(greeting)  // Output: Hello, Bernard!
}
```

---

## 🧠 TL;DR

| Function     | Outputs to        | Returns a value? | Use Case                  |
|--------------|-------------------|------------------|---------------------------|
| `fmt.Printf` | Console (stdout)  | ❌ No             | Just display something    |
| `fmt.Sprintf`| A string          | ✅ Yes            | Store or manipulate output|

---

## 🔁 Python vs Go Analogy

| Go                 | Python Equivalent         | Meaning                          |
|--------------------|---------------------------|----------------------------------|
| `fmt.Printf(...)`  | `print(...)`              | Prints directly                  |
| `fmt.Sprintf(...)` | `message = f"...{}..."`   | Returns a formatted string to reuse |

---

## ✅ Side-by-side comparison

### Go:

```go
package main

import "fmt"

func main() {
    var name string = "Bernard"
    fmt.Printf("Hello, %s!\n", name)                    // Prints directly
    var message string = fmt.Sprintf("Hello, %s!", name)
    fmt.Println(message)                                // Prints the stored string
}
```

### Python:

```python
name = "Bernard"
print(f"Hello, {name}!")            # Direct print
message = f"Hello, {name}!"         # Store string
print(message)                      # Print stored string
```
