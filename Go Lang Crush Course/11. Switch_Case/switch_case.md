# 🌀 Go (Golang) `switch` Statement Guide

## 🎟️ What is a `switch` Statement in Go?

A `switch` statement in Go is a clean alternative to multiple `if-else` conditions. It checks cases one by one against a given expression and executes the first matching one.

### 📚 Basic Syntax

```go
switch expression {
case value1:
    // code block
case value2:
    // code block
default:
    // code block if no cases match
}
```

- Go’s switch automatically breaks after the first match (no need to write `break` like in C/C++).
- You can evaluate expressions, types, and even use switch without a condition (acting like `if-else`).

## 🎯 Example: Booking Application

Let’s say we’re building a booking app that processes different booking types like:

- "flight"
- "hotel"
- "car"
- anything else → unknown

### ✍️ Example Code

```go
package main

import (
    "fmt"
)

func main() {
    bookingType := "hotel"

    switch bookingType {
    case "flight":
        fmt.Println("Processing flight booking... ✈️")
    case "hotel":
        fmt.Println("Processing hotel reservation... 🏨")
    case "car":
        fmt.Println("Processing car rental... 🚗")
    default:
        fmt.Println("Unknown booking type. Please try again.")
    }
}
```

### 🧾 Output

```
Processing hotel reservation... 🏨
```

## 🔁 Advanced Examples

### Multiple Values in One Case

```go
switch bookingType {
case "flight", "train", "bus":
    fmt.Println("Processing a travel booking... 🚆🚌✈️")
}
```

### ⚙️ Switch Without Condition (like `if-else`)

```go
price := 1200

switch {
case price < 500:
    fmt.Println("Budget booking.")
case price >= 500 && price < 1000:
    fmt.Println("Standard booking.")
default:
    fmt.Println("Premium booking.")
}
```

### 🔒 Type Switch (Bonus: for interfaces)

```go
var booking interface{} = 42

switch v := booking.(type) {
case string:
    fmt.Println("It's a string:", v)
case int:
    fmt.Println("It's an int:", v)
default:
    fmt.Println("Unknown type.")
}
```

---

## 🔄 What is `default` in a `switch` Statement?

The `default` keyword works like the `else` block in `if-else`.

👉 It runs when none of the `case` values match.

### 💥 `switch` + `default`

```go
package main

import "fmt"

func main() {
    status := "pending"

    switch status {
    case "confirmed":
        fmt.Println("Booking confirmed ✅")
    case "cancelled":
        fmt.Println("Booking cancelled ❌")
    default:
        fmt.Println("Unknown booking status ⚠️")
    }
}
```

🧾 Output:

```
Unknown booking status ⚠️
```

### 🧱 Equivalent `if-else`

```go
package main

import "fmt"

func main() {
    status := "pending"

    if status == "confirmed" {
        fmt.Println("Booking confirmed ✅")
    } else if status == "cancelled" {
        fmt.Println("Booking cancelled ❌")
    } else {
        fmt.Println("Unknown booking status ⚠️")
    }
}
```

🧾 Output:

```
Unknown booking status ⚠️
```

---

## ✅ When to Use Which?

| Situation                     | Use `switch` ✅       | Use `if-else` ❗         |
|------------------------------|------------------------|-------------------------|
| Multiple discrete values     | ✅ Cleaner             | 👎 Can be bulky         |
| Complex logical conditions   | 👎 Limited             | ✅ Better               |
| Type checking                | ✅ Type switch         | 👎 Not supported        |
| Ranges or compound logic     | 👎 Not ideal           | ✅ Best suited          |

---

## 🔁 One More Comparison: Booking Prices

### 🎯 With `if-else`

```go
price := 800

if price < 500 {
    fmt.Println("Budget booking 💸")
} else if price <= 1000 {
    fmt.Println("Standard booking 🧳")
} else {
    fmt.Println("Premium booking 💼")
}
```

### 🌀 With `switch`

```go
switch {
case price < 500:
    fmt.Println("Budget booking 💸")
case price <= 1000:
    fmt.Println("Standard booking 🧳")
default:
    fmt.Println("Premium booking 💼")
}
```

---

## ✨ Summary

- `default` in `switch` = `else` in `if-else`
- Use `switch` when comparing one value to many known options.
- Use `if-else` for more complex logic, compound conditions, or range checks.
