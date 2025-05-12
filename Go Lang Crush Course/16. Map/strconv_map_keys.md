# Converting Non-String Keys to Strings in Go 🧳🔄

If you need to use non-string keys like integers or floats as keys in a `map[string]string`, you can convert them using the `strconv` package.

Here’s a quick breakdown with examples:

## 🔄 Converting Other Data Types to String using strconv

### ✅ Example: int to string
```go
import "strconv"

id := 42
idStr := strconv.Itoa(id) // int to string

countries := map[string]string{
    idStr: "Ghana", // now we can use it as a key
}
fmt.Println(countries["42"]) // Output: Ghana
```

### ✅ Example: float to string
```go
import "strconv"

price := 12.99
priceStr := strconv.FormatFloat(price, 'f', 2, 64) // float64 to string

products := map[string]string{
    priceStr: "Laptop",
}
fmt.Println(products["12.99"]) // Output: Laptop
```

### ✅ Example: bool to string
```go
import "strconv"

status := true
statusStr := strconv.FormatBool(status)

flags := map[string]string{
    statusStr: "Active",
}
fmt.Println(flags["true"]) // Output: Active
```

## 🧠 This is especially useful when:
- You're getting dynamic data (e.g., from user input or external APIs)
- You need to serialize map keys
- You're storing results for lookup based on values that aren’t naturally strings

---

## 🧩 Breaking Down `strconv.FormatFloat`

### 📌 Function:
```go
strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

This formats a float (`f`) as a string according to the format you specify.

### 🧩 Parameters Explained:
| Argument         | Value in Your Code | Meaning                           |
|------------------|--------------------|-----------------------------------|
| `price`          | `12.99`            | The float you want to convert     |
| `'f'`            | `f`                | Format — `f` means decimal format |
| `2`              | `2`                | Precision — Keep 2 digits after the decimal |
| `64`             | `64`               | Bit size — Treat as `float64`     |

### ✅ Example in Action:
```go
price := 12.9876
priceStr := strconv.FormatFloat(price, 'f', 2, 64)
fmt.Println(priceStr)
```

🟢 Output: `12.99`

### 🎯 Format Options (`fmt byte`):
| Format | Description                          |
|--------|--------------------------------------|
| `f`    | Decimal point (e.g., `12.34`)        |
| `e` or `E` | Scientific notation (e.g., `1.23e+01`) |
| `g` or `G` | Compact representation (either `f` or `e` depending on size) |

### 🧠 Why specify `64` for `bitSize`?
- Use `64` for `float64` values (most common in Go).
- Use `32` for `float32`.

This ensures Go handles the number's precision correctly.

---

## 🔁 Convert uint to string in Go

If you're using a `uint` (unsigned integer), and want to convert it to a string, you'd use a different function from the `strconv` package:

### ✅ Example with `uint`:
```go
var myUint uint = 42
str := strconv.FormatUint(uint64(myUint), 10)
fmt.Println(str) // Output: "42"
```

### 🧠 Why `uint64(myUint)`?
The function signature is:
```go
func FormatUint(i uint64, base int) string
```
So, you must explicitly cast your `uint` to `uint64`.

### 📌 Parameters:
| Argument         | Meaning                           |
|------------------|-----------------------------------|
| `uint64(myUint)` | The unsigned integer you want to convert |
| `10`             | Base for formatting (e.g., `10` for decimal, `2` for binary, `16` for hex) |

### ✅ Example with base `16` (hex):
```go
id := uint(255)
hexStr := strconv.FormatUint(uint64(id), 16)
fmt.Println(hexStr) // Output: "ff"
```

---

## 🟨 Alternative: Use `fmt.Sprintf` (Less Strict)
```go
str := fmt.Sprintf("%d", myUint)
```
It works with `uint`, `int`, etc., but it's less type-safe and less performant than `strconv`.

---

## 🚀 Summary
| Type   | Function               | Example                            |
|--------|------------------------|------------------------------------|
| `float`| `strconv.FormatFloat()` | `strconv.FormatFloat(12.99, 'f', 2, 64)` |
| `int`  | `strconv.Itoa()`        | `strconv.Itoa(123)`                |
| `uint` | `strconv.FormatUint()`  | `strconv.FormatUint(42, 10)`       |
