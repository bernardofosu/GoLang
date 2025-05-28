
# 🐹 Go Built-in Identifiers

Here are examples of **built-in identifiers** in Go that are part of the language itself and do **not require any import** because they’re defined in Go’s **predeclared universe**:

---

## 1️⃣ Built-in Functions

You can use these **anywhere without imports**:

- `panic()`
- `recover()`
- `print()`, `println()`
- `len()`
- `cap()`
- `new()`
- `make()`
- `copy()`
- `append()`
- `close()`
- `delete()`

### 🧪 Example:

```go
func main() {
    s := []int{1, 2, 3}
    println(len(s))    // len is built-in
    println(cap(s))    // capacity of slice

    a := make([]int, 5) // create slice with make

    append(a, 10)       // append to slice

    panic("Something went wrong")  // panic is built-in
}
```

> 🛑 `recover` is typically used in deferred functions to catch panics.

---

## 2️⃣ Predeclared Types

These basic types are built-in:

- `int`, `int8`, `int64`, `uint`, `byte`, `rune`
- `string`
- `bool`
- `error` (interface)

### 🧪 Example:

```go
var x int = 42
var s string = "hello"
var b bool = true
var err error = nil
```

---

## 3️⃣ Predeclared Constants

- `true`, `false`
- `iota` (used in `const` declarations)

### 🧪 Example:

```go
const (
    a = iota  // 0
    b = iota  // 1
)
```

---

## 🚫 What Requires Import?

| Identifier          | Category         | Requires Import?        |
|---------------------|------------------|--------------------------|
| `panic`             | Built-in func    | ❌ No                    |
| `len`               | Built-in func    | ❌ No                    |
| `int`               | Built-in type    | ❌ No                    |
| `error`             | Built-in type    | ❌ No                    |
| `true`              | Built-in const   | ❌ No                    |
| `fmt.Println`       | Package func     | ✅ Yes (`"fmt"`)         |
| `json.Unmarshal`    | Package func     | ✅ Yes (`"encoding/json"`) |

---

## 🔍 Clarifying `print()` vs `fmt.Print()`

| Function         | Source       | Requires Import? | Use Case                |
|------------------|--------------|------------------|--------------------------|
| `print()`        | Built-in     | ❌ No            | Simple debug output      |
| `println()`      | Built-in     | ❌ No            | Debug with newline       |
| `fmt.Print()`    | `fmt` pkg    | ✅ Yes           | Formatted output         |
| `fmt.Println()`  | `fmt` pkg    | ✅ Yes           | Formatted + newline      |

> 💡 Use `print`/`println` for quick debugging only. Use `fmt` for production-grade output.

---

## ✅ Minimal Built-in Example

```go
func main() {
    print("Hello, ")
    println("world!")
}
```

No imports needed 🎉

