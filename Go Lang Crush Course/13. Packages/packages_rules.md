# 🐹 Go Package Rules: Structuring Multi-Package Projects

---

## ✅ Overview
In Go, when you want multiple packages:

- ✅ You create **separate folders** for each package.
- 📁 Each folder can have any name (e.g., `add/`, `subtract/`).
- 📄 Inside each folder, `.go` files must declare the **same package name**.
- 🚫 These packages **do not run on their own** — they must be imported into a **main** package.

---

## 📁 Example Project Structure

```
calculator/
├── add/
│   └── add.go          → package add
├── subtract/
│   └── subtract.go     → package subtract
├── multiply/
│   └── multiply.go     → package multiply
├── main/
│   └── start.go        → package main
└── go.mod
```

---

## 📦 Package Files

### `add/add.go`
```go
package add

func Sum(a, b int) int {
    return a + b
}
```

### `subtract/subtract.go`
```go
package subtract

func Diff(a, b int) int {
    return a - b
}
```

---

## 🚀 `main/start.go` (Entry Point)
```go
package main

import (
    "fmt"
    "calculator/add"
    "calculator/subtract"
)

func main() {
    fmt.Println("Addition:", add.Sum(5, 3))
    fmt.Println("Subtraction:", subtract.Diff(9, 4))
}
```

---

## 🔥 Key Rules to Remember

| Rule                                | Explanation                                           |
|-------------------------------------|-------------------------------------------------------|
| ✅ One folder = One package          | All `.go` files in a folder must use the same package |
| ✅ Only `package main` gets executed | Go runs only the file with `func main()`              |
| ✅ You must import your packages     | Other code is only used if imported in main           |
| ❌ No automatic execution            | Packages won’t run unless called from main            |

---

## 🧠 Think of it this way
> “In Go, all roads lead to `package main` — it’s the only package that runs, and everything else is a helper until you invite it to the party.” 🎉

---

## 🛠️ Sample Project: `golang-app`

```
golang-app/
├── go.mod
├── add/
│   └── add.go
├── subtract/
│   └── subtract.go
└── main/
    └── start.go
```

### 📄 go.mod
```go
module golang-app

go 1.20
```

### ➕ `add/add.go`
```go
package add

func Sum(a, b int) int {
    return a + b
}
```

### ➖ `subtract/subtract.go`
```go
package subtract

func Diff(a, b int) int {
    return a - b
}
```

### 🚀 `main/start.go`
```go
package main

import (
    "fmt"
    "golang-app/add"
    "golang-app/subtract"
)

func main() {
    fmt.Println("5 + 3 =", add.Sum(5, 3))
    fmt.Println("5 - 2 =", subtract.Diff(5, 2))
}
```

---

## ✅ How to Run
```bash
cd golang-app

# Run directly
go run ./main

# OR build it
go build -o calc ./main
./calc
```

---

## 🧠 Reminders
- `start.go` can have any name, but it must be in `package main`
- All imports use the format: `golang-app/<folder>` based on `go.mod`
- **Folder name = package name** (common Go convention)

---

## 🔍 Why You Must Use Full Import Paths
Go won’t recognize:
```go
import "add" // ❌ won’t work
```
Because:
- It looks for a standard package (like `fmt`, `net/http`) or
- A third-party module from the Go proxy

✅ Instead, always use:
```go
import "golang-app/add"
```

---

## 📄 Example: go.mod with GitHub module
```go
module github.com/bernard/calculator
```

### Updated Structure:
```
calculator/
├── go.mod                         ← module github.com/bernard/calculator
├── add/add.go                     ← package add
├── subtract/subtract.go          ← package subtract
└── main/start.go                  ← package main
```

### ✅ Correct imports in `main/start.go`
```go
import (
    "fmt"
    "github.com/bernard/calculator/add"
    "github.com/bernard/calculator/subtract"
)
```

---

## 🔑 Summary Rules

| Thing                               | Correct                                               |
|------------------------------------|-------------------------------------------------------|
| ✅ Use module path + subfolder     | `github.com/you/projectname/foldername`               |
| ❌ Don’t just use folder name      | `import "add"` → ❌ won’t work                         |
| ✅ Check your go.mod               | Always reference the module name there                |

---

