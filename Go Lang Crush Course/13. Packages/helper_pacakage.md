# 🛠️ Using Multiple `.go` Files in the Same Package (Go)

In Go, if you have multiple `.go` files in the same package (e.g., `main`), and they're in the same folder, they **share the same scope** — meaning you **don’t need to import anything** to use functions across those files.

So yes, you can create a `helper.go` with `package main`, and just run:

```bash
go run main.go helper.go
```

Or even simpler:

```bash
go run .
```

(as long as you’re in the same directory).

---

## 📁 Example Structure

```
mathapp/
├── go.mod
├── main.go
└── helper.go
```

---

## `main.go`

```go
package main

import "fmt"

func main() {
    a := 10.0
    b := 2.0

    fmt.Println("Add:", add(a, b))
    fmt.Println("Subtract:", subtract(a, b))
    fmt.Println("Multiply:", multiply(a, b))

    result, err := divide(a, b)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Divide:", result)
    }
}
```

---

## `helper.go`

```go
package main

import "errors"

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

---

## ✅ Running It

From within the `mathapp/` folder:

```bash
go run . # . means current dir
# or
go run main.go helper.go
```

### 🖥️ Output:

```
Add: 12
Subtract: 8
Multiply: 20
Divide: 5
```

---

✅ This approach is **great for small projects or quick demos**!
