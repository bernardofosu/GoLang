## 🏗️ Go Build — What Is It?

`go build` is a command that compiles your Go code into a binary executable.

### 🔧 Usage:
```bash
go build
```
If you're in a folder with a `main` package, this creates an executable named after the folder.

### 📌 What It Does:
| Task                     | Description                                   |
|--------------------------|-----------------------------------------------|
| ✅ Compiles code         | Turns `.go` files into machine code.           |
| 🧹 Resolves dependencies | Automatically downloads required modules.    |
| 📦 Skips unused packages | Builds only what's needed.                  |
| ⚠️ Does not run code     | Unlike `go run`, this only builds.          |

---

### 🔍 Example 1: Building a CLI app

**main.go:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

```bash
go build
```

**💡 Output:**
- On Linux/macOS: `./your-folder-name`
- On Windows: `your-folder-name.exe`

**Run it:**
```bash
./your-folder-name
```

---

### 🔍 Example 2: Building a specific file
```bash
go build main.go
```

---

### ⚙️ Advanced: Custom Output Name
```bash
go build -o hello-world
```
This creates a binary named `hello-world`.

---

## ✅ Summary
| Command         | Action                              |
|----------------|-------------------------------------|
| `go build`     | Compiles your Go code.              |
| `go run`       | Builds and runs the code.           |
| `go install`   | Builds + installs binary in `$GOBIN` |

---

## 📦 Project Structure (Use Case)
```
math-app/
├── go.mod                         👈 module name: github.com/yourusername/math-app
├── main.go                        👈 entry point, package main
├── add/
│   └── add.go                     👈 package add
├── subtract/
│   └── subtract.go                👈 package subtract
└── utils/
    └── utils.go                   👈 package utils
```

### 🧠 Content Example
**add/add.go:**
```go
package add

func Add(a, b int) int {
    return a + b
}
```

**subtract/subtract.go:**
```go
package subtract

func Subtract(a, b int) int {
    return a - b
}
```

**utils/utils.go:**
```go
package utils

func Square(n int) int {
    return n * n
}
```

**main.go:**
```go
package main

import (
    "fmt"
    "github.com/yourusername/math-app/add"
    "github.com/yourusername/math-app/subtract"
    "github.com/yourusername/math-app/utils"
)

func main() {
    fmt.Println("Add:", add.Add(2, 3))
    fmt.Println("Subtract:", subtract.Subtract(5, 1))
    fmt.Println("Square:", utils.Square(4))
}
```

### 🏗️ Using `go build`
```bash
cd math-app
go build
```

### 🧾 Output:
- If your folder is named `math-app`:
  - On macOS/Linux: `./math-app`
  - On Windows: `math-app.exe`

### 🧠 Key Points
- ✅ Only run `go build` once from the root, not per package.
- ✅ Binary name = folder name by default.
- ✅ Only imported packages are compiled.

---

## 🏗️ Custom Build Name
```bash
go build -o my-custom-name
```
This creates a binary named `my-custom-name`.

**🧪 Example:**
```bash
cd math-app
go build -o math-runner
```
**Run it:**
```bash
./math-runner
```

You can also specify a path:
```bash
go build -o bin/math-runner
```

---

## ✅ After `go build -o custom-name`

🎉 Anyone can run the binary if:
- ✅ Same OS (e.g., Linux → Linux)
- ✅ Same architecture (e.g., amd64)
- ✅ No recompilation is needed

**🧠 You don't need Go to *run* the binary, only to *build* it.**

---

## 🛠️ Cross-Compiling Examples

Build for different OS/architecture:
```bash
GOOS=windows GOARCH=amd64 go build -o math-runner.exe
GOOS=linux GOARCH=amd64 go build -o math-runner-linux
GOOS=darwin GOARCH=amd64 go build -o math-runner-mac
```

---

## 🧪 Ubuntu → Other Linux
Yes! ✅

| Source OS | Target OS | Works?          |
|-----------|-----------|-----------------|
| Ubuntu    | Fedora    | ✅ Yes           |
| Ubuntu    | Alpine    | ⚠️ Use static build |

To build a static binary:
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mytool
```

---

Want a script or Makefile for all targets? Just say the word! 💬

