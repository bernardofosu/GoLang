# Formatting Decimal Points in Go (Golang)

Great question! 🙌 In Go, you can format a float to 1 or 2 decimal places using `fmt.Sprintf` or `fmt.Printf`.

---

## ✅ Example using `fmt.Printf` (just prints):

```go
package main

import "fmt"

func main() {
    gpa := 3.8567

    fmt.Printf("GPA (1 decimal): %.1f\n", gpa)  // GPA (1 decimal): 3.9
    fmt.Printf("GPA (2 decimals): %.2f\n", gpa) // GPA (2 decimals): 3.86
}
```

---

## ✅ Example using `fmt.Sprintf` (returns string):

```go
package main

import "fmt"

func main() {
    gpa := 3.8567

    formatted := fmt.Sprintf("%.2f", gpa)
    fmt.Println("Formatted GPA:", formatted) // Formatted GPA: 3.86
}
```

---

## 🔁 If you need to round and keep it as a `float64` (not string):

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    gpa := 3.8567

    // Round to 2 decimal places
    rounded := math.Round(gpa*100) / 100
    fmt.Println("Rounded float:", rounded) // Rounded float: 3.86
}
```

---

## 🔍 What does `%.1f` mean?

Here's what `%.1f` means in Go (also similar in C, Python, etc.):

- `%` → Start of a format verb.
- `.1` → Precision: Show 1 digit after the decimal point.
- `f` → Format as a floating-point number (decimal number).

| Format | Meaning                    | Example Output (for 3.8567) |
|--------|----------------------------|-----------------------------|
| %.1f   | Float with 1 decimal place | 3.9                         |
| %.2f   | Float with 2 decimals      | 3.86                        |
| %.3f   | Float with 3 decimals      | 3.857                       |

---

## ✅ More Examples:

```go
value := 12.34567

fmt.Printf("%.1f\n", value)  // Output: 12.3
fmt.Printf("%.2f\n", value)  // Output: 12.35
fmt.Printf("%.3f\n", value)  // Output: 12.346
```

The `f` means float format, and the `.N` controls how many digits after the dot.

Let me know if you're displaying, rounding for math, or saving to DB — each case can use a slightly different method!
