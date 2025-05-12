# 🧮 Project Structure

```
mathapp/
├── go.mod
├── main.go
├── add/
│   └── add.go
├── subtract/
│   └── subtract.go
├── multiply/
│   └── multiply.go
├── divide/
│   └── divide.go
```

---

## 1️⃣ go.mod — Go module file

Run this in terminal inside `mathapp/` to initialize:

```bash
go mod init mathapp
```

---

## 2️⃣ main.go

```go
package main

import (
    "fmt"
    "mathapp/add"
    "mathapp/subtract"
    "mathapp/multiply"
    "mathapp/divide"
)

func main() {
    a := 20.0
    b := 5.0

    fmt.Println("Add:", add.Sum(a, b))
    fmt.Println("Subtract:", subtract.Difference(a, b))
    fmt.Println("Multiply:", multiply.Product(a, b))

    result, err := divide.Quotient(a, b)
    if err != nil {
        fmt.Println("Divide Error:", err)
    } else {
        fmt.Println("Divide:", result)
    }
}
```

---

## 3️⃣ add/add.go

```go
package add

func Sum(a, b float64) float64 {
    return a + b
}
```

---

## 4️⃣ subtract/subtract.go

```go
package subtract

func Difference(a, b float64) float64 {
    return a - b
}
```

---

## 5️⃣ multiply/multiply.go

```go
package multiply

func Product(a, b float64) float64 {
    return a * b
}
```

---

## 6️⃣ divide/divide.go

```go
package divide

import "errors"

func Quotient(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

---

## ✅ Run the App

In your terminal, from the root of `mathapp`, run:

```bash
go run main.go
```

You should see:

```
Add: 25
Subtract: 15
Multiply: 100
Divide: 4
```

---

## 📦 `package main` ≠ `func main()`

### `package main` tells Go:

> “This code is meant to be compiled as a standalone executable (not a shared library).”

### `func main()` is:

> The entry point Go looks for when running that executable.

---

### So:
| Element       | Purpose                                               |
|---------------|-------------------------------------------------------|
| `package main`| Defines this file as part of an executable program    |
| `func main()` | Tells Go where to start execution                     |

---

## 🔄 Why `main.go` has `package main`

It's convention, not a requirement.  
You can name the file anything — `start.go`, `app.go`, `hello.go`, whatever — but if it contains `func main()`, it must be in `package main`.

And Go will still look for `func main()` when you run:

```bash
go run .
```

or

```bash
go build
./yourbinary
```

---

## 🧠 Bonus: Other Packages

Your custom packages like `add`, `subtract`, etc., **don't** need `package main` —  
They use their own package name so they can be modular and imported into any app, including `main`.

---

## TL;DR

✅ `package main`: required to build an executable.  
✅ `func main()`: required to define the entry point.  
🗂️ Filename (`main.go`) is just a convention — not required.
