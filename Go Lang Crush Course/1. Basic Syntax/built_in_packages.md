# ✅ Built-in Functions and Packages in Python, Node.js, and Go

In languages like Python, Node.js, and Go, there are built-in libraries or modules that provide commonly used functionality. However, the way they work in each language differs:

---

## 1. 🐍 Python

In Python, built-in functions like `print()` are readily available without needing an explicit import because they are part of Python’s core language features. These are not tied to a specific module you need to import, as Python provides a set of built-in functions for immediate use.

### Example:
```python
print("Hello, World!")  # No need to import anything
```

The core functions, such as `print()`, are available globally in the Python interpreter, so you don’t need to worry about importing them explicitly.

---

## 2. 🟢 Node.js

In Node.js, functions like `console.log()` are part of the global object, which means you can use them directly without importing any specific package or module. The `console` module is automatically available in all Node.js applications.

### Example:
```javascript
console.log("Hello, World!");  // No need to import anything
```

Since `console` is part of the global environment in Node.js, there’s no need to import it separately like you would need to with other modules (e.g., `fs` for file system operations).

---

## 3. 🐹 Go

In Go, everything is organized into packages, and you have to explicitly import the packages that provide the functionality you need.

Functions like `fmt.Println()` are part of the `fmt` package, so you must import that package first before using it.

### Example:
```go
package main

import "fmt"  // Importing the fmt package

func main() {
    fmt.Println("Hello, World!")  // Using fmt.Println() from the fmt package
}
```

In Go, you can’t directly use `Println()` without importing `fmt` first because `fmt` is an external package that you have to access by importing it.

---

## 📌 Summary

- **Python**: The `print()` function is part of the core language, so it is available globally without an import.
- **Node.js**: Functions like `console.log()` are part of the global object in Node.js, so they don’t require an import.
- **Go**: You need to explicitly import packages (e.g., `fmt`) to access specific functionality like `fmt.Println()`.

In short, **Python** and **Node.js** provide certain functionalities (like `print()` and `console.log()`) as global features or built-ins, while **Go** requires you to import specific packages to access features.

