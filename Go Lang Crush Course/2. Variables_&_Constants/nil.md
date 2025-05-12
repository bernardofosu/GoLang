## 🔍 Is `nil` in Go the same as zero or undefined?

| Concept       | Go Equivalent     | Description                                                                 |
|---------------|-------------------|-----------------------------------------------------------------------------|
| **nil**       | ✅ `nil`           | Means the variable **has no value**, like `null` or `None` in other languages. |
| **zero value**| ✅ Depends on type | The **default value** for a variable of a certain type (e.g., `0` for `int`, `""` for `string`). |
| **undefined** | ❌ Not a thing     | Go **does not have** "undefined" variables. All variables are **always initialized**, either explicitly or with a **zero value**. |

---

### 📌 Examples

```go
var num int            // zero value is 0
var str string         // zero value is ""
var ptr *int           // zero value is nil
var list []int         // zero value is nil
var m map[string]int   // zero value is nil

// Evaluations:
num == 0      // ✅
str == ""     // ✅
ptr == nil    // ✅
list == nil   // ✅
m == nil      // ✅
```

---

### 🧠 So:

- `nil` is used for **reference types** (pointers, slices, maps, channels, interfaces, functions).
- **Primitive types** like `int`, `string`, `bool`, etc., get a **zero value**, not `nil`.
- Go **never leaves anything undefined** — that's one of the ways it prevents runtime bugs.

