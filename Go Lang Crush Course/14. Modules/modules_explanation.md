## 🧩 What is a Go Module?

A **Go module** is a collection of Go packages defined in a directory with a `go.mod` file at its root.

It tells Go:

- 📛 What your project is called (module)
- 📦 What dependencies you're using
- 🧮 What Go version you're targeting

> 💡 Think of a module as the boundary of your Go project.

---

## 📄 go.mod File: The Heart of a Module

### Example:

```go
module github.com/yourusername/coolproject

go 1.20

require (
    github.com/gorilla/mux v1.8.0
)
```

This means:

- 🔖 My project’s name is `github.com/yourusername/coolproject`
- 🧪 I’m using Go 1.20
- 🔌 I depend on `github.com/gorilla/mux`

---

## 🏗️ Standard Go Project Structure with Modules

```
coolproject/
├── go.mod              ← declares the module name
├── go.sum              ← contains checksums of dependencies
├── main/               ← package main (entry point)
│   └── main.go
├── add/                ← a custom package
│   └── add.go
├── subtract/           ← another custom package
│   └── subtract.go
└── utils/              ← helper utilities
    └── format.go
```

---

## 📦 Package Declarations

- Files in `add/` should say: `package add`
- Files in `main/` must say: `package main`
- Files in `utils/` should say: `package utils`

---

## 🔌 Imports Based on Module Name

If your `go.mod` says:

```go
module github.com/yourusername/coolproject
```

Then in `main.go`, your imports should be:

```go
import (
    "fmt"
    "github.com/yourusername/coolproject/add"
    "github.com/yourusername/coolproject/subtract"
    "github.com/yourusername/coolproject/utils"
)
```

> ✅ Even if your actual folder is just called `coolproject`, Go sees the module name from `go.mod`.

---

## 🔄 go.mod + go.sum: Dependency Management

| File      | Purpose                                                   |
|-----------|-----------------------------------------------------------|
| `go.mod` | Declares your module, Go version, and direct dependencies |
| `go.sum` | Ensures reproducible builds with hashes of dependencies   |

> 🧠 Think of it like `package.json` + `package-lock.json` in Node.js or `requirements.txt` in Python.

---

## 🏁 Running a Go Module

From the root:

```sh
go run ./main
```

Or build it:

```sh
go build -o myapp ./main
./myapp
```

---

## 🧠 Tips for Structuring Larger Go Projects

| Folder     | Purpose                                  |
|------------|------------------------------------------|
| `cmd/`     | Sub-packages for CLI or app binaries     |
| `pkg/`     | Reusable logic packages                  |
| `internal/`| Private packages (not importable externally) |
| `api/`     | API-related structs and routes           |
| `configs/` | Config files                             |

---

## ✅ Mini Checklist for Modules

- [ ] Did you run `go mod init <module-name>`?
- [ ] Are you using correct package declarations in each folder?
- [ ] Are you importing using the module path, not folder name?
- [ ] Is your entry point inside `package main` with `func main()`