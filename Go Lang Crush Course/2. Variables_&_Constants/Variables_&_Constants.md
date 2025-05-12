# 🧮 Variables in Go

## 🐹 Naming Conventions in Go (Golang)

In Go (Golang), the naming convention for variables is **camelCase**. This means that the first word is lowercase, and each subsequent word starts with an uppercase letter, with no spaces or underscores between them.

## Examples:
- `myVariable`
- `userAge`
- `totalAmount`

## Conventions:

### 1. Camel Case (camelCase)
- Used for variable names, function names, and method names.

### 2. Pascal Case (PascalCase)
- Typically used for types, structs, and interfaces. Pascal case is similar to camel case but starts with an uppercase letter.
  
  **Examples**:
  - `PersonDetails`
  - `ProductInfo`

### 3. Snake Case (snake_case)
- Not used in Go, as it is common in languages like Python.

## Quick Reference:
- **Variable and function names**: `camelCase`
- **Type names (structs, interfaces, etc.)**: `PascalCase` (starts with uppercase)

A variable is a named storage location used to hold a value that can change during program execution.

🔹 **Syntax: Declaring Variables**

```go
var name string = "Bernard"
```

🔹 **You can also use type inference:**

```go
var age = 30 // Go infers that age is an int
```

🔹 **Or even shorter with the shorthand (`:=`) inside functions:**

```go
country := "Ghana"
```

✅ **You can reassign a variable:**

```go
country = "Kenya"
```

---

# 🔒 Constants in Go

A constant is a named value that doesn’t change after it's defined.

🔹 **Syntax:**

```go
const pi = 3.14
```

- Must be assigned a value at the time of declaration.
- Cannot be reassigned later.

```go
pi = 3.1415 // ❌ Error: cannot assign to pi
```

---

# 🔑 Key Differences

| Feature      | `var` (Variable)       | `const` (Constant)             |
|--------------|------------------------|--------------------------------|
| Value        | Can change             | Cannot change                  |
| Declaration  | With `var` or `:=`     | With `const` only              |
| Usage        | General-purpose storage| Fixed values (like pi, app name, etc.) |
| Inference    | Yes (`:=` or `var`)    | No (must set a value)          |

---

# ✨ Example: Both Together

```go
package main

import "fmt"

func main() {
    var name = "Bernard"
    const appName = "GoLangApp"

    fmt.Println("Name:", name)
    fmt.Println("App:", appName)

    name = "Kwasi"      // ✅ Allowed
    // appName = "Test" // ❌ Error

    fmt.Println("New Name:", name)
}
```


# 🔧 Variable Behavior in Go

Yes, exactly! When you declare a variable with `var` in Go, it means the variable's value can change, but the **type** of the variable cannot change once it's set. The type is determined either **explicitly** or through **type inference** (if you don't specify the type, Go infers it based on the initial value you assign).

## 📝 Key Points:

- `var` is used to declare a variable, and it allows you to **change the value** of the variable during the program's execution.
- However, the **type of the variable is fixed** after the initial assignment and **cannot be changed** later.

## 💡 Example:

```go
var age = 30  // Go infers the type as int
age = 35      // ✅ This is fine because age is still of type int

// If you try to change the type of the variable:
age = "Hello" // ❌ Compile-time error: cannot assign string to int
```

## 🔍 Explanation:

- **Changing Value**: You can update the value of the variable (e.g., `age = 35`), but it will always remain of the same type (in this case, `int`).
- **Type is Fixed**: Once `age` is declared as an `int` (through `var age = 30`), you cannot reassign it to a value of a different type (e.g., a `string`).

## 🎯 Why is this important?

This feature helps **prevent bugs** related to unexpected type changes during program execution, providing more **type safety** than dynamically typed languages.
