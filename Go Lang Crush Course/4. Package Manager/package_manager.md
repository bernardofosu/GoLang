# 📦 Go Standard Library & Package Management

When you install the Go compiler, a standard library comes built-in with it — no need to install fmt, math, time, or others separately.

## ✅ Standard Library Includes

Here are some commonly used packages that come with Go out of the box:

| Package         | Purpose                                             |
|----------------|-----------------------------------------------------|
| fmt            | Formatted I/O (printing, etc.)                      |
| math           | Basic math operations like Round, Sqrt, Pow, etc.   |
| time           | Date/time handling                                  |
| strings        | String manipulation (e.g., Contains, ToUpper)       |
| strconv        | Convert between strings and numbers                 |
| os             | Interact with the operating system (files, env vars)|
| io             | Input/output interfaces                             |
| bufio          | Buffered I/O                                        |
| net/http       | Web server or HTTP client                           |
| encoding/json  | Working with JSON                                   |

### ✅ Example: No Need to Install These
```go
import (
    "fmt"
    "math"
    "strings"
    "time"
)
```
Go handles these internally and quickly — so as long as you're using the official Go compiler, they're always available without extra downloads.

---

## 📦 Go's Package Management System

Go has its own built-in package manager — and it's super streamlined compared to tools like pip or npm.

### 🔧 `go get` (Classic)
Used to fetch and install external packages:
```bash
go get github.com/gorilla/mux
```
This:
- Downloads the module from the internet (like GitHub)
- Adds it to your project's `go.mod` file
- Places it in your Go workspace (`$GOPATH/pkg/mod`)

### 📘 `go.mod` (Go Modules)
This is Go's version of `package.json` or `requirements.txt`.

You create it using:
```bash
go mod init your-module-name
```
Then install packages with:
```bash
go get example.com/your/dependency
```
Or just import the package in your code and run:
```bash
go build
```
Go will auto-fetch what's missing based on `go.mod`.

### 🧾 Example `go.mod`:
```go
module myapp

go 1.21

require (
    github.com/gorilla/mux v1.8.0
)
```

---

## 🗂️ Summary

| Language     | Package Manager     | Dependency File     |
|--------------|----------------------|----------------------|
| Python       | pip                  | requirements.txt     |
| JavaScript   | npm / yarn           | package.json         |
| Go           | go get, go mod       | go.mod               |

