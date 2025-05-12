# 📦 Understanding Go Package Structure

## 🟢 ✅ What's Required

| Thing                          | Must match?       |
|-------------------------------|-------------------|
| Package name ↔️ Directory name| ✅ Usually, yes (by convention, not rule) |
| File name ↔️ Package name     | ❌ No, they can be completely different |

## 🧠 Example: File names ≠ Package name

```go
// File: math.go
package utils

func Add(a, b int) int {
	return a + b
}

// File: calculator.go
package utils

func Subtract(a, b int) int {
	return a - b
}
```

✅ Both files are in the `utils` package, so they can share code and be imported like:

```go
import "yourproject/utils"
```

## 🟡 But… Keep This in Mind

- All files in the same folder **must** declare the **same package name**.
- It’s convention (and smart practice) for the **folder name = package name**.

## 🔁 Example That Breaks

```go
// File: add.go
package add

func Sum(a, b int) int {
	return a + b
}

// File: subtract.go
package subtract // ❌ in the same folder

func Difference(a, b int) int {
	return a - b
}
```

❌ This will **fail to compile**, unless `add` and `subtract` are in separate folders.

---

## ✅ Summary Table

| Concept                              | Rule/Convention         |
|-------------------------------------|--------------------------|
| File name = Package name            | ❌ Not required          |
| Folder name = Package name          | ✅ Strongly recommended |
| All files in a folder: same package | ✅ Required             |

---

## 🧱 Example: Go Math App Structure

```
mathapp/
├── go.mod
├── main.go                → package main
├── helpers.go             → package main ✅
├── add/
│   ├── sum.go             → package add ✅
│   └── extra_add.go       → package add ✅ (different file name, same package)
├── subtract/
│   └── subtract.go        → package subtract
```

### 🔍 main.go & helpers.go

```go
// main.go
package main

import (
	"fmt"
	"mathapp/add"
	"mathapp/subtract"
)

func main() {
	fmt.Println("Sum:", add.Sum(5, 2))
	fmt.Println("Difference:", subtract.Difference(5, 2))

	printWelcomeMessage() // From helpers.go
}
```

```go
// helpers.go
package main // ✅ Must match the package in main.go since it's in the same folder

func printWelcomeMessage() {
	fmt.Println("Welcome to the math app!")
}
```

💡 **Key Point**: File names (`main.go`, `helpers.go`) don’t need to match `package main`, but they **must all declare the same package** if they’re in the same folder.

---

### 🔍 add/sum.go & add/extra_add.go

```go
// add/sum.go
package add

func Sum(a, b int) int {
	return a + b
}
```

```go
// add/extra_add.go
package add // ✅ Still part of the "add" package, just a different file

func AddThree(a, b, c int) int {
	return a + b + c
}
```

💡 These are two separate files, but since they’re in the same folder, they must use the same package name: `add`.

---

### 🔍 subtract/subtract.go

```go
// subtract/subtract.go
package subtract

func Difference(a, b int) int {
	return a - b
}
```

✅ This is in a different folder (`subtract/`), so it can have its own package name.

---

## ⚠️ What Would Break?

If you accidentally did this:

```go
// add/sum.go
package add

// add/extra_add.go
package mathoperations // ❌ Different package name in the same folder
```

⛔ Go will give a compile-time error, because **all files in a folder must belong to the same package**.




# 🧱 Go Project Structure Guide

A well-structured Go (Golang) project makes your codebase clean, scalable, and easy to maintain.  
This layout is inspired by conventions from the Go community (e.g., [golang-standards/project-layout](https://github.com/golang-standards/project-layout)).

---

## 🗂️ Recommended Project Structure

```
your-project/
├── cmd/                  # 🏁 Main applications for this project
│   └── yourapp/          # ▶️ Example: `main.go` for CLI or server
│       └── main.go
├── pkg/                  # 📦 Library code ok to use by external apps
│   └── yourlib/          # 🔁 Shared, reusable code
├── internal/             # 🚫 Private application and library code
│   └── yourpackage/      # 🔒 Code for internal use only
├── api/                  # 🔗 Protocol definitions (OpenAPI, gRPC)
│   └── proto/            # 📜 gRPC or REST API specs
├── web/                  # 🌐 Frontend files or static assets
├── configs/              # ⚙️ Configuration files (YAML, JSON, etc.)
├── scripts/              # 🛠️ Automation scripts (Bash, Python, Go)
├── deployments/          # 🚀 Deployment configs (Docker, K8s, etc.)
├── test/                 # 🧪 Integration & test data
├── docs/                 # 📚 Project documentation
├── go.mod                # 📄 Go module definition
├── go.sum                # 🔒 Go dependency checksums
└── README.md             # 📝 Project overview
```

---

## 🔍 Folder Breakdown

| 📁 Folder      | 🧾 Purpose                                                |
|----------------|-----------------------------------------------------------|
| `cmd/`         | Entry point of your application (e.g., `cmd/server/main.go`) |
| `pkg/`         | Code intended to be imported by other projects            |
| `internal/`    | Code that must not be imported by others (Go enforces this) |
| `api/`         | Definitions and schemas for APIs (OpenAPI, gRPC)          |
| `configs/`     | Static configuration files                                |
| `scripts/`     | Useful scripts for CI, tooling, and automation            |
| `deployments/` | Kubernetes manifests, Docker Compose files, etc.         |
| `test/`        | Test-specific helpers or data                             |
| `web/`         | Static or frontend web resources                          |
| `docs/`        | Documentation, diagrams, and guides                       |


✅ Example: Simple CLI App
```sh
awesome-cli/
├── cmd/
│   └── awesomecli/
│       └── main.go
├── internal/
│   └── commands/
│       └── hello.go
├── pkg/
│   └── utils/
│       └── string_utils.go
├── go.mod
└── README.md
```
🔧 Tips

    Keep main.go short—just parse config and call functions from internal/ or pkg/.

    Prefer internal/ for your business logic to avoid accidental use in other projects.

    Use pkg/ for code you expect to share (like custom logger or auth helpers).

    Separate concerns: API specs ≠ business logic ≠ config ≠ scripts.

