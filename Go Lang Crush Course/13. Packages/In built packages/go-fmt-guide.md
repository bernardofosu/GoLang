
# 📦 fmt Package in Go

The `fmt` package in Go is one of the most used packages for input/output operations, especially printing to the terminal. 🖨️✨

## ✨ Overview

`fmt` stands for **format**. It helps you print output, take input, and format data beautifully!

```go
import "fmt"
```

---

## 🖨️ Printing Output

### ✅ `fmt.Println(...)`

Prints and moves to a new line.

```go
fmt.Println("Hello, World!")  // ➡️ Hello, World!\n
```

### ✅ `fmt.Print(...)`

Prints without adding a newline.

```go
fmt.Print("Hello ")
fmt.Print("World")  // ➡️ Hello World
```

### ✅ `fmt.Printf(...)`

Prints with formatting using placeholders.

```go
name := "Gopher"
age := 5
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

| Verb | Meaning              | Example      |
|------|----------------------|--------------|
| %s   | string               | "hello"      |
| %d   | decimal (int)        | 42           |
| %f   | float                | 3.14         |
| %v   | default format (any) | true, 123    |
| %T   | type of the variable | int, string  |

---

## 📋 Formatting Instead of Printing

Sometimes you just want to create a string, not print it.

### ✅ `fmt.Sprintf(...)`

Returns a formatted string.

```go
result := fmt.Sprintf("Pi is about %.2f", 3.14159)
fmt.Println(result) // Pi is about 3.14
```

---

## 🧠 Reading Input

### ✅ `fmt.Scan(...)`

```go
var name string
fmt.Print("Enter your name: ")
fmt.Scan(&name)
fmt.Println("Hello,", name)
```

### ✅ `fmt.Scanf(...)` — like Printf but for input

```go
var age int
fmt.Scanf("Your age is %d\n", &age)
```

---

## 🔁 Summary Table

| Function   | Use                    |
|------------|------------------------|
| Print()    | Print with no newline  |
| Println()  | Print with newline     |
| Printf()   | Print with format      |
| Sprintf()  | Format and return as string |
| Scan()     | Read user input        |
| Scanf()    | Read formatted input   |

---

## 🧪 Bonus Example

```go
package main

import "fmt"

func main() {
    name := "Golang"
    version := 1.21
    fmt.Printf("Welcome to %s! 🚀 Version: %.2f\n", name, version)
}
```

**Output:**

```
Welcome to Golang! 🚀 Version: 1.21
```

---

## 🧠 No `f` = No formatting

| Function       | Placeholders allowed? | Adds newline? |
|----------------|------------------------|----------------|
| fmt.Print()    | ❌ No                   | ❌ No          |
| fmt.Println()  | ❌ No                   | ✅ Yes         |
| fmt.Printf()   | ✅ Yes (`%s`, `%d`, ...) | ❌ No (unless manually) |

### ❌ Wrong usage:

```go
name := "Go"
fmt.Println("My name is %s", name) // Output: My name is %s name
```

### ✅ Correct usage:

```go
fmt.Printf("My name is %s\n", name) // Output: My name is Go
```

Let me know if you want a cheat sheet on all the placeholders (`%s`, `%v`, `%T`, etc.)! 🧰
