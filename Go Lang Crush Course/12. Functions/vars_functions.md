# 🌍 Package-Level Variables in Go

## 🧠 What Are They?

Package-level variables are variables declared outside of any function, at the top level of a `.go` file, and they are accessible to all functions in that package — including `main()`.

They are also sometimes called **global variables** (but only within that package).

---

## 🔤 Syntax

```go
package main

import "fmt"

// 🌍 This is a package-level variable
var appName = "GoApp"

func greet() {
    fmt.Println("Welcome to", appName)
}

func main() {
    greet()
    fmt.Println("Running", appName)
}
```

### ✅ Output:
```
Welcome to GoApp
Running GoApp
```

---

## 👵️ When Should You Use Package-Level Variables?

- ✅ When a value needs to be shared across multiple functions
- ✅ When a variable should retain its value throughout the execution
- ✅ For configuration, constants, or caches

---

## ⚠️ Be Careful!

- 🔴 Overusing package-level variables can lead to tight coupling, hard-to-test code, and unexpected side effects — especially in concurrent programs.

Prefer function parameters and return values when possible.

---

## 🔐 Constant Example

You can also declare package-level constants:

```go
const pi = 3.14

func area(radius float64) float64 {
    return pi * radius * radius
}
```

---

## 🤖 Multiple Package-Level Variables

```go
var (
    username = "admin"
    port     = 8080
    isLive   = true
)
```

---

## ⚠️ := Cannot Be Used at Package Level

You **cannot** use the `:=` short variable declaration at the package level.

```go
// ❌ This will NOT compile:
// msg := "Hello, Go"

// ✅ Correct way:
var msg = "Hello, Go"
```

**Why?**

- `:=` is only valid inside function bodies.
- Package-level declarations require `var` or `const`.

---

## 🧠 Summary Table

| Feature              | Description                                     |
|----------------------|-------------------------------------------------|
| Scope                | Visible to all functions in the package         |
| Declared             | Outside functions, at file top level            |
| Lifetime             | Lasts the entire runtime                        |
| Good for             | Constants, shared config, caching               |
| Avoid for            | Temporary data, function-local behavior         |
| := allowed?          | ❌ No, only inside functions                   |
| Declared with var    | ✅ Yes                                          |

---

Want a mini project example using package-level variables or a breakdown of how they differ from local variables? I can do that too! 🚰

