# 📘 Common Go Types with Examples

Let’s look at the most common Go types: `int`, `string`, and `float64` — along with simple examples.

---

## 🔢 `int` Example

```go
package main

import "fmt"

func main() {
    var age int = 25
    fmt.Println("Age:", age)
}
```

📝 `int` is used for whole numbers (no decimals).

---

## 🔤 `string` Example

```go
package main

import "fmt"

func main() {
    var name string = "Bernard"
    fmt.Println("Name:", name)
}
```

📝 `string` stores text (UTF-8 encoded characters).

---

## 🌊 `float64` Example

```go
package main

import "fmt"

func main() {
    var temperature float64 = 36.6
    fmt.Println("Temperature:", temperature)
}
```

📝 `float64` is used for decimal (floating point) numbers. It’s more precise than `float32`.

---

## 🎯 Bonus: All Three in One

```go
package main

import "fmt"

func main() {
    var name string = "Bernard"
    var age int = 30
    var gpa float64 = 3.85

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("GPA:", gpa)
}
```

---

## 🧠 Why Explicit Types Matter

Exactly! 💡 One key reason to explicitly declare a variable’s type in Go is when you define the variable **without immediately assigning a value**. Go is statically typed, so it needs to know the type at compile time.

### 🔧 Example:

```go
package main

import "fmt"

func main() {
    var score int       // Declared, but not assigned
    score = 100         // Assigned later
    fmt.Println("Score:", score)
}
```

### ✅ But when you assign right away, Go can infer the type:

```go
package main

import "fmt"

func main() {
    var score = 100  // Go infers 'int' from the value 
    fmt.Println("Score:", score)
}
```

📝 **In short:**
- If you declare now, assign later → you must specify the type.
- If you declare and assign together → Go can infer the type.

---

# 📦 All Go Data Types (Quick Breakdown)

## 🔢 Numeric Types

### Integers:
- Signed: `int`, `int8`, `int16`, `int32`, `int64`
- Unsigned: `uint`, `uint8` (alias for `byte`), `uint16`, `uint32`, `uint64`, `uintptr`
- `int` and `uint` are platform-dependent (32 or 64-bit)

### Floating-Point:
- `float32`, `float64`

### Complex Numbers:
- `complex64` (real and imaginary parts are `float32`)
- `complex128` (real and imaginary parts are `float64`)

---

## 🔠 Text Type
- `string` (immutable UTF-8 text)

---

## 🧮 Boolean Type
- `bool` (values: `true` or `false`)

---

## 🕵️‍♂️ Alias Types
- `byte` → alias for `uint8`
- `rune` → alias for `int32` (used for Unicode code points)

---

## 📦 Composite Types
- **Array** – fixed-size: `[5]int`
- **Slice** – dynamic array: `[]int`
- **Struct** – custom data types:

```go
type Person struct {
    name string
    age int
}
```

- **Map** – key-value pairs: `map[string]int`
- **Pointer** – e.g., `*int`
- **Function** – e.g., `func(int, int) int`
- **Interface** – abstraction for multiple types
- **Channel** – used for goroutine communication: `chan int`

---

Let me know if you want to explore any of these further or build a project with them! 🚀
