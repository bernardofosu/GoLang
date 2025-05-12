
# expected 'package', found Printsyntax
```sh
expected 'package', found Printsyntax
```
## 🚀 Entry Point in Go
// package declaration (every Go file starts with this)
```go
package main
```
## What Is package main Then?
- main is a special package name used only for programs that are meant to be executed (like CLI tools, web servers, etc).
- The Go compiler looks for a main package and its main() function to know where to start running the code.
  
In Go, the "expected declaration" error typically occurs when the Go compiler expects a declaration, such as a function, variable, or type definition, but encounters something else (like syntax errors or misplaced code).

## What is an Entry Point in Go?

The entry point of a Go program is the function where the program begins execution. For Go programs, the entry point is always the `main` function inside a `main` package.

### Entry Point in Go:
- **Main Function**: The designated entry point.
- **Main Package**: Where the program begins.

## Key Concepts:

### 1. Main Package
- The `main` package tells Go that this file is where the program execution begins. Without this, Go won't know where to start.

### 2. Main Function
- The `main()` function inside the `main` package is the starting point of a Go program. When the Go program is executed, it will look for this function to begin processing.

### 3. Expected Declaration
- When you see an error like "expected declaration," Go is looking for something it expects, such as the start of a function or variable, but it finds something out of place, like misplaced syntax. In Go, the `main` function should only exist once in the `main` package in any executable program. This is because Go programs are structured around the idea of having a single entry point, which is the `main` function in the `main` package.

## Explanation:

### Main Package:
- A Go application, when executed, must have one `main` package and one `main` function within it. This is the entry point of the program.
- Only one `main` function can exist in the `main` package.
- The `main` package is specifically used to indicate the entry point of the program.**
In Go, everything is organized under packages. A package is a way to group related functions, types, and variables together, making code modular and reusable.



# Error message undefined: PrintIncompilerUndeclaredName
```sh
 error message undefined: PrintIncompilerUndeclaredName
```
## 📦 Importing the fmt Package in Go

```go
package main // main package is where the entry point is located

import "fmt" // importing the fmt package for printing

func main() { // entry point function
    fmt.Print("Hello, Go!") // this is where execution starts
}
```
### 📦 fmt Package

**fmt stands for formatting.**

It is a standard Go package used for:
- Printing output to the console (like Println, Printf, etc.).
- Scanning input from the user (like Scan, Scanf).
- Formatting strings and values.

## where functions come from and which packages provide which functions
- [Packages Docs](https://pkg.go.dev/std)
- [Get started with Go](https://go.dev/doc/tutorial/getting-started)


## No packages found for open file D:\GoLang\main.go.go list, package main
```sh
No packages found for open file D:\GoLang\main.go.go list
package main
```
### 🧾 Example `go.mod`:
```go
module my-project
```

## Why You Need to Import fmt:

### 1. Accessing Library Functions
Go is a statically typed language, and each package in Go is isolated from others. This means that functions, variables, and types defined in one package cannot be used in another unless you explicitly import that package.

- **Example**: The `fmt.Println()` function is defined in the `fmt` package, so to use `fmt.Println()` in your code (to print to the console), you need to import the `fmt` package first.

### 2. Package Organization
Go follows the principle of package-based organization. Each package has its own namespace, so you can't just use any function or type from another package unless it's explicitly imported.

## Common Functions in fmt:
- `fmt.Println()`: Prints a line of text followed by a newline.
- `fmt.Print()`: Prints text without a newline.
- `fmt.Printf()`: Allows formatted printing (like `printf` in C).
- `fmt.Scan()`: Reads input from the user.

## How Importing Works:
In Go, imports allow you to access external functionality provided by Go's standard library or third-party libraries. When you write `import "fmt"`, you're telling the Go compiler to link the `fmt` package with your code, which gives you access to the functions defined in that package.

If you don't import the `fmt` package, the compiler doesn't know where to find `fmt.Println()`, and you’ll get an error saying the function is undefined.

## Conclusion:
- `import "fmt"` is necessary in Go to access functions like `fmt.Println()`, which is used for printing output to the console.
- Packages in Go are self-contained, and you need to import them explicitly to use their functionality.


# 🚀 Running `main.go` in Go: How It Works

When you run:

```bash
go run main.go
```

Go looks for the **entry point** of your program. Here's how it works 👇

---

## 🧠 What Go Looks For

✅ Your file **must** contain:

```go
package main

func main() {
    // your code here
}
```

---

## 📌 Behind the Scenes

When executing `go run main.go`, Go:

1. 🔍 **Checks for the `main` package**
   - Only `package main` can be compiled into an **executable**.

2. 👀 **Looks for `func main()`**
   - This is the **starting point** of execution.
   - Just like `int main()` in C or `def main()` in Python (when manually called).

---

## ❌ What Happens if It's Missing?

If either `package main` or `func main()` is missing, you'll get an error:

- ❗ `go: cannot run non-main package`
- ❗ `main.go:7:1: expected 'package'`

---

## ✅ Example That Works

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
    // fmt.Print("Hello, Go! ") you have to give space 
}
```
Println is a function from the fmt package used to print output to the console. The **ln** part stands for **"line"** and signifies that the function will print the provided text followed by a newline character (\n)

This works because:

- ✅ It has `package main`
- ✅ It includes `func main()`

---

## 💡 Tip

Always remember:

- `package main` = you want to build an **executable**
- `func main()` = Go's **entry point** to run your code

That's it! Now go run your Go code like a pro 💪🐹

