# 🧠 Go Conditional Statements: `if`, `else if`, `else` + Operators Explained 💡

## 🧠 What is an `if` statement?

An `if` statement lets your program make decisions based on conditions.

Think of it like this:

> “If something is true, then do this.”

Go evaluates a condition (which results in `true` or `false`) and then runs a block of code **only if the condition is true**.

### 🔍 Syntax of an `if` statement:

```go
if condition {
    // code to execute if condition is true
}
```

### ✅ Example:

```go
package main

import "fmt"

func main() {
    // age := 25
    var age = 25
    if age >= 18 {
        fmt.Println("You are an adult.")
    }
}
```

🧾 **Output**:  
`You are an adult.`  
Because `age` is greater than or equal to 18.

---

## 💡 What about `else`?

An `else` block tells the program **what to do if the `if` condition is false**.

### 🧱 Syntax:

```go
if condition {
    // true block
} else {
    // false block
}
```

### ✅ Example:

```go
package main

import "fmt"

func main() {
    age := 15

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }
}
```

🧾 **Output**:  
`You are a minor.`  
Because `age` is not greater than or equal to 18.

---

## 🔄 What is `else if`?

`else if` is used when you have **multiple conditions to check** — not just true or false, but several possibilities.

### 🧱 Syntax:

```go
if condition1 {
    // do this if condition1 is true
} else if condition2 {
    // do this if condition2 is true
} else {
    // do this if none are true
}
```

### ✅ Example:

```go
package main

import "fmt"

func main() {
    score := 85

    if score >= 90 {
        fmt.Println("Grade A")
    } else if score >= 80 {
        fmt.Println("Grade B")
    } else if score >= 70 {
        fmt.Println("Grade C")
    } else {
        fmt.Println("Fail")
    }
}
```

🧾 **Output**:  
`Grade B`  
Because score is 85, which matches the second condition.

---

## 🎁 Bonus: Short Statement in `if`

Go allows you to **declare a variable inside the `if` statement**:

```go
if num := 10; num%2 == 0 {
    fmt.Println("Even number")
}
```

Here, `num := 10` is only accessible **inside** the `if` block.

---

# 🤔 Assignment Operator vs Comparison Operators

Let’s dive into the difference between the **assignment operator** (`=`) and **comparison operators** (`==`, `!=`, `<`, `>`, etc.) in Go 🧐

---

## 🔧 1. Assignment Operator `=`

The assignment operator is used to **assign a value** to a variable.

### 🧱 Syntax:

```go
variableName = value
```

### ✅ Example:

```go
package main

import "fmt"

func main() {
    age := 25 // This is short for: var age int = 25
    fmt.Println("Your age is:", age)
}
```

`=` assigns the value 25 to the variable `age`.

---

## ⚖️ 2. Comparison Operators

Comparison operators are used to **compare values**.  
They always return a **boolean result**: `true` or `false`.

### 🧮 Common comparison operators in Go:

| Operator | Meaning              | Example (`x = 10, y = 5`) |
|----------|----------------------|----------------------------|
| `==`     | Equal to             | `x == y` → `false`         |
| `!=`     | Not equal to         | `x != y` → `true`          |
| `<`      | Less than            | `x < y` → `false`          |
| `>`      | Greater than         | `x > y` → `true`           |
| `<=`     | Less than or equal   | `x <= 10` → `true`         |
| `>=`     | Greater than or equal| `x >= y` → `true`          |

---

### ✅ Example using comparison:

```go
package main

import "fmt"

func main() {
    age := 20

    if age == 18 {
        fmt.Println("You are exactly 18 years old.")
    } else if age > 18 {
        fmt.Println("You are older than 18.")
    } else {
        fmt.Println("You are younger than 18.")
    }
}
```

🧾 **Output**:  
`You are older than 18.`

---

## ❗ Important:

- `=` is **not** the same as `==`
- `=` sets a value  
- `==` checks a value

### ❌ Common mistake:

```go
if age = 18 {  // ❌ Compilation error!
    fmt.Println("Wrong!")
}
```

### ✅ Correct:

```go
if age == 18 {
    fmt.Println("Correct!")
}
```

✅ Key Notes for Go:
- Parentheses are NOT required around the condition (e.g., if age > 18 is fine).
- Curly braces {} are mandatory even for single-line statements.
- Go also enforces formatting rules — tools like gofmt will reformat your code to follow Go standards.
