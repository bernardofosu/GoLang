# 🔄 Go's strconv Package Cheatsheet

The `strconv` package in Go is used to convert between strings and other basic data types — like `int`, `float64`, `bool`, etc. It's super useful when you're dealing with input/output, files, or data from APIs that come in as strings.

## 🔧 Common Uses of strconv

### 1. 🔢 String to Integer

```go
import "strconv"

str := "42"
num, err := strconv.Atoi(str) // Atoi = ASCII to Integer
if err != nil {
    fmt.Println("Conversion failed:", err)
}
fmt.Println(num) // 42
```

### 2. 🔁 Integer to String

```go
num := 42
str := strconv.Itoa(num) // Itoa = Integer to ASCII
fmt.Println(str) // "42"
```

### 3. 📏 String to Float

```go
str := "3.14"
f, err := strconv.ParseFloat(str, 64)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(f) // 3.14
```

### 4. 🔄 Float to String

```go
f := 3.14159
str := strconv.FormatFloat(f, 'f', 2, 64)
// 'f' = decimal format, 2 = precision, 64 = float64
fmt.Println(str) // "3.14"
```

### 5. ✅ String to Boolean

```go
str := "true"
b, err := strconv.ParseBool(str)
fmt.Println(b) // true
```

### 6. 🔁 Boolean to String

```go
b := true
str := strconv.FormatBool(b)
fmt.Println(str) // "true"
```

## 🧠 Summary Table

| Conversion         | Function                 |
|--------------------|--------------------------|
| String → Int       | `strconv.Atoi()`         |
| Int → String       | `strconv.Itoa()`         |
| String → Float     | `strconv.ParseFloat()`   |
| Float → String     | `strconv.FormatFloat()`  |
| String → Bool      | `strconv.ParseBool()`    |
| Bool → String      | `strconv.FormatBool()`   |

---

### 🔢 Convert String → Int

Use `strconv.Atoi()` or `strconv.ParseInt()`:

```go
import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Conversion error:", err)
    }
    fmt.Println(num) // 123 as int
}
```

### 🔢 Convert Float → Int

You can't use `strconv` for this — just type cast:

```go
func main() {
    f := 45.67
    i := int(f) // Truncates: removes decimal part
    fmt.Println(i) // Output: 45
}
```

🧠 **Note:** This truncates, not rounds. So `99.99` becomes `99`, not `100`.

---

### 🔢 Convert String → Float → Int (combo!)

When string has a float value, convert it in two steps:

```go
func main() {
    str := "98.6"

    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("Float parse error:", err)
        return
    }

    i := int(f)
    fmt.Println(i) // Output: 98
}
```

---

## ✅ Summary

| From      | To     | How                                |
|-----------|--------|-------------------------------------|
| `string`  | `int`  | `strconv.Atoi()`                   |
| `float64` | `int`  | `int(f)`                           |
| `string`  | `float`| `strconv.ParseFloat()`             |
| `string` → `float` → `int` | First `ParseFloat()`, then `int(f)` |
