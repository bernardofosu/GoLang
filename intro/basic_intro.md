# Go (Golang) Overview

Go (or Golang) was created by Google in 2007 and officially released in 2009. Here's the breakdown:

## 👨‍💻 Creators of Go
Go was designed by three legendary engineers:

| Name                | Notable Work                                       |
|---------------------|----------------------------------------------------|
| Robert Griesemer    | Worked on the Java HotSpot VM & V8 engine         |
| Rob Pike            | Helped build Unix and Plan 9 at Bell Labs         |
| Ken Thompson        | Co-created Unix and the C programming language     |

These guys are OGs of computer science! 🧠💻

## 🎯 Why Was Go Created?
Google built Go to solve real-world problems they were facing, like:

| Problem Google Faced                   | How Go Solved It                          |
|----------------------------------------|-------------------------------------------|
| ✅ Slow compilation in C++             | ⚡ Go compiles very fast                  |
| ✅ Cumbersome concurrency               | 🧵 Built-in goroutines & channels         |
| ✅ Complex dependency management        | 📦 Simple imports & go.mod system        |
| ✅ Messy codebases                     | ✨ Enforces clean, readable code          |
| ✅ Poor tooling                        | 🛠️ Ships with great tools (e.g. go fmt, go test, go vet) |

## 💡 Go’s Goals
- **Simplicity** – Minimalistic but powerful
- **Speed** – Both in compilation & runtime
- **Concurrency-first** – Easily write scalable systems
- **Readability** – Easy to maintain by large teams
- **Static typing** – Catch bugs early

## 🧑‍💼 Who Uses Go Today?
Go is widely used by:
- **Google** – For many internal tools
- **Docker** – Entire Docker engine is written in Go
- **Kubernetes** – Written in Go
- **Terraform** – Written in Go
- **Cloudflare**, **Uber**, **Twitch**, **Dropbox** – All use Go

# 🧪 Go Variables & Data Types Example

```go
package main

import "fmt"

func main() {
    // 🔢 Integers
    var age int = 30
    var year int64 = 2025

    // 🔤 Strings
    var name string = "Bernard"

    // 🔁 Short declaration (Go idiom)
    country := "Ghana"

    // ✅ Boolean
    var isCloudEngineer bool = true

    // 🔢 Float
    var temperature float64 = 36.6

    // 🧮 Constants
    const pi = 3.14159

    // 🕳 Zero values
    var uninitializedInt int       // defaults to 0
    var uninitializedString string // defaults to ""

    // 🖨 Print values
    fmt.Println("Name:", name)
    fmt.Println("Country:", country)
    fmt.Println("Age:", age)
    fmt.Println("Year:", year)
    fmt.Println("Is Cloud Engineer?", isCloudEngineer)
    fmt.Println("Temperature:", temperature)
    fmt.Println("PI:", pi)
    fmt.Println("Uninitialized Int:", uninitializedInt)
    fmt.Println("Uninitialized String:", uninitializedString)
}
```

# 🧠 Breakdown of Common Go Types

| Type                    | Example           | Description                              |
|-------------------------|-------------------|------------------------------------------|
| `int`, `int64`, `uint` | `42`, `2025`      | Whole numbers (signed/unsigned)         |
| `float32`, `float64`   | `3.14`, `36.6`    | Decimal numbers                          |
| `string`                | `"hello"`         | Text                                     |
| `bool`                  | `true`, `false`   | Boolean values                           |
| `const`                 | Immutable value    | Used for fixed constants                 |