# 📝 Comments in Go

In Go, comments are used to add explanatory notes in the code. They are ignored by the compiler and don't affect the execution of the program. Go supports both single-line and multi-line comments.

## 1. 💬 Single-Line Comments

A single-line comment starts with `//` and continues to the end of the line.  
**Example:**

```go
// This is a single-line comment
fmt.Println("Hello, World!")
```

## 2. 🧱 Multi-Line Comments (Block Comments)

Multi-line comments begin with `/*` and end with `*/`. They can span multiple lines.  
**Example:**

```go
/*
This is a multi-line comment.
It can span multiple lines.
*/
fmt.Println("Hello, World!")
```

## 3. 📚 Documentation Comments

In Go, comments that are placed immediately before a function, method, or type declaration are treated as documentation comments. These comments can be used by tools like `godoc` to generate documentation for the code.  
**Example:**

```go
// Greet function prints a greeting message.
func Greet(name string) {
    fmt.Println("Hello, " + name)
}
```

This style of commenting is used for documenting the purpose of functions, variables, or packages in Go.
