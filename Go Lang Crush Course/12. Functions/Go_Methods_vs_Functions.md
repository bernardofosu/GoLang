
# 🔍 Methods vs. Functions in Go

## ✅ Regular Function with Multiple Inputs and Outputs

```go
func Calculate(a int, b int) (int, int) {
    sum := a + b
    product := a * b
    return sum, product
}

func main() {
    s, p := Calculate(3, 4)
    fmt.Println("Sum:", s, "Product:", p)
}
```

- 🧾 **Inputs**: `a`, `b` (both `int`)  
- 📤 **Outputs**: `sum` and `product` (both `int`)  
- 🚫 **No receiver** → standalone function

---

## ✅ Method with Receiver, Multiple Inputs and Outputs

```go
type Calculator struct {
    Name string
}

func (c *Calculator) Calculate(a int, b int) (int, int) {
    sum := a + b
    product := a * b
    return sum, product
}

func main() {
    calc := &Calculator{Name: "SimpleCalc"}
    s, p := calc.Calculate(3, 4)
    fmt.Println("Sum:", s, "Product:", p)
}
```

- Same inputs/outputs as the function  
- 🧍‍♂️ **Receiver**: `(c *Calculator)` binds method to `Calculator` type  
- 🔐 Can access struct data like `c.Name`

---

## 🔍 Comparison Table

| Feature                | Regular Function                                      | Method with Receiver                          |
|------------------------|-------------------------------------------------------|------------------------------------------------|
| 📝 **Syntax**           | `func Name(a, b int) (int, int)`                     | `func (r *Type) Name(a, b int) (int, int)`     |
| 🙅‍♂️ **Receiver**         | ❌ None                                              | ✅ Yes — `(r *Type)`                            |
| 🔐 **Access struct fields** | ❌ Cannot (unless passed manually)                  | ✅ Yes — has access to Type fields              |
| 📞 **Call Syntax**       | `Calculate(3, 4)`                                   | `myStruct.Calculate(3, 4)`                      |
| 🧩 **Belongs to a type**  | ❌ No                                                | ✅ Yes                                           |

---

## 🧠 Key Takeaway

- 📌 A **function** is independent.  
- 🧠 A **method** is tied to a type and can use that type's fields.

---

## ✅ Example: Function Accepting Mixed Types

```go
package main

import (
    "fmt"
)

func RepeatWord(word string, times int) string {
    result := ""
    for i := 0; i < times; i++ {
        result += word
    }
    return result
}

func main() {
    repeated := RepeatWord("Go", 3)
    fmt.Println(repeated) // Output: GoGoGo
}
```

### 🔍 Breakdown

- `func RepeatWord(word string, times int) string`:
  - Takes a `string` and an `int`
  - Returns a `string`
  - ✅ Mixing types is totally fine in Go
