# 🧠 Understanding `var` in Go — Declaration, Initialization, and Inference

## 🧠 What is `var` in Go?

In Go, the `var` keyword is used to **declare variables**.  
A variable holds data that your program can use and modify.

---

## 1️⃣ Declaration Only

This means you introduce the variable and optionally give it a type, but you don't assign it a value yet.

```go
var name string  // 📦 name is declared as a string, but has no value yet
```

➡️ Go gives it the **zero value** automatically:

- `string` → `""` (empty string)  
- `int` → `0`  
- `bool` → `false`

### ✅ Example:

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Println("Age:", age) // ➡️ Output: Age: 0
}
```

---

## 2️⃣ Declaration + Initialization

This is when you declare a variable and assign a value at the same time.

```go
var age int = 25
```

✅ Type is clear, and value is assigned.

### ✅ Example:

```go
package main

import "fmt"

func main() {
    var city string = "Accra"
    fmt.Println("City:", city) // ➡️ Output: City: Accra
}
```

---

## 3️⃣ Type Inference with `var`

Go can **guess the type** from the value you assign. You can skip writing the type:

```go
var country = "Ghana"  // Go infers it's a string
```

### ✅ Example:

```go
package main

import "fmt"

func main() {
    var height = 1.75
    fmt.Printf("Height: %.2f meters\n", height) // ➡️ Output: Height: 1.75 meters
}
```

---

## 4️⃣ Short Declaration (`:=`)

This is the most common way **inside functions**.  
It declares and initializes a variable without using `var`.

```go
name := "Kwame"
```

### ✅ Example:

```go
package main

import "fmt"

func main() {
    message := "Hello, Go!"
    fmt.Println(message) // ➡️ Output: Hello, Go!
}
```

⚠️ `:=` can **only be used inside a function** — not at the package level!

---

## 🧾 Summary Table

| Pattern                     | Example              | Notes                          |
|----------------------------|----------------------|--------------------------------|
| Declaration only           | `var x int`          | `x = 0` (zero value)           |
| Declaration + Initialization | `var x int = 10`     | Explicit type                  |
| Type Inference             | `var x = "hello"`    | Go figures out the type        |
| Short Declaration          | `x := 20`            | Most common inside functions   |
