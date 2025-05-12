# 🧾 Go Input Functions: Scan, Scanln, and Scanf

Great question! 🙌 Here's a clear breakdown of the differences between `Scan()`, `Scanln()`, and `Scanf()` in Go — all of which are used to receive user input, but behave slightly differently:

## 📌 `fmt.Scan()`

### 🔹 Description:
- Reads input space-separated
- Stops scanning on whitespace (space or newline)
- Continues scanning until all arguments are filled

### ✅ Example:
```go
var name string
var age int

fmt.Print("Enter your name and age: ")
fmt.Scan(&name, &age)

fmt.Println("Value of num:", num)    // 42
fmt.Println("Address of num:", &num) // Memory address of num
```
**Input:** `Bernard 30`
**Output:** `name = "Bernard", age = 30`

## 📌 `fmt.Scanln()`

### 🔹 Description:
- Like `Scan()`, but stops scanning at newline
- Useful when input is on one line only
- It reads values until Enter is pressed

### ✅ Example:
```go
var name string
var age int

fmt.Print("Enter your name and age: ")
fmt.Scanln(&name, &age)
```
**Input:** `Bernard 30`
**Output:** `name = "Bernard", age = 30`

❗ If you enter more values than variables, it causes an error because it expects the full line to be consumed.

---

## 📌 `fmt.Scanf()` — formatted input

### 🔹 Description:
- Uses a format string like `"%s %d"`
- Allows more control over how input is parsed
- Must match the format exactly, or it may fail

### ✅ Example:
```go
var name string
var age int

fmt.Print("Enter your name and age: ")
fmt.Scanf("%s %d", &name, &age)
```
**Input:** `Bernard 30`
**Output:** `name = "Bernard", age = 30`

---

## 🧠 TL;DR: Quick Comparison
| Function   | Stops at               | Uses format string? | Good for           |
|------------|------------------------|----------------------|--------------------|
| `Scan()`   | Whitespace (space or newline) | ❌ No              | Simple input       |
| `Scanln()` | Newline (Enter key)    | ❌ No              | One-line input     |
| `Scanf()`  | Follows specified format | ✅ Yes             | Precise control    |

---

## ✅ Using `fmt.Scan()` – space-separated input
```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Enter your name and age (e.g. Bernard 30): ")
    fmt.Scan(&name, &age)

    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
```
**Input format:** `Bernard 30`

---

## ✅ Using `fmt.Scanln()` – enter-separated input on same line
```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Enter your name and age: ")
    fmt.Scanln(&name, &age)

    fmt.Printf("Welcome, %s! Age: %d\n", name, age)
}
```
**Input format:** `Bernard 30`

---

## ✅ Using `fmt.Scanf()` – with format string
```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Enter your name and age: ")
    fmt.Scanf("%s %d", &name, &age)

    fmt.Printf("Nice to meet you, %s! You're %d years old.\n", name, age)
}
```
**Input format:** `Bernard 30`

---

## 📝 Notes:
- `%s` is for a string
- `%d` is for an integer
- `&variable` is used to pass the memory address where the input will be stored

