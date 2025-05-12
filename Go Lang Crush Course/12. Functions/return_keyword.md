## 🔙 What is `return` in Go?

In Go, the `return` keyword is used to exit a function and optionally send a value or values back to the caller.

---

## 📦 Basic Syntax

```go
func add(a int, b int) int {
    return a + b
}
```

✅ In this example:
- `add` is a function that takes two integers.
- It uses `return` to send back their sum.

---

## 🧠 Why Use `return`?
Because you often want your function to give back a result, like:
- Calculations ✅
- Error messages 🧨
- Modified data 📦

---

## 🧮 Example: Return One Value

```go
func greet(name string) string {
    return "Hello, " + name
}
```

---

## 🔁 Return Multiple Values
One of Go’s coolest features is multiple return values.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

✔️ This function returns:
- A result (`float64`)
- An error (`error`)

---

## 🆗 Named Return Values
Go lets you name your return values, so you can return without arguments:

```go
func getUser() (name string, age int) {
    name = "Alice"
    age = 30
    return // same as return name, age
}
```

This is called a **naked return** — useful for short functions.

---

## 🚫 Exiting Early
You can use `return` to exit a function early:

```go
func checkAge(age int) string {
    if age < 18 {
        return "Too young"
    }
    return "Welcome"
}
```

---

## 🧪 Example With `main()`

```go
package main

import "fmt"

func square(n int) int {
    return n * n
}

func main() {
    result := square(5)
    fmt.Println("Square:", result)
}
```

🖨️ Output: `Square: 25`

---

## ✔️ Returning Arrays

```go
func getArray() [3]int {
    return [3]int{1, 2, 3}
}
```

Note: Arrays are fixed in size and **copied**, not passed by reference.

---

## ✅ Returning Slices (Preferred)

```go
func getSlice() []int {
    return []int{10, 20, 30, 40}
}
```

✅ Slices are:
- Dynamic
- Efficient (lightweight internal structure)
- Passed by reference

---

## 🔁 Returning Multiple Slices

```go
func splitEvenOdd(nums []int) ([]int, []int) {
    var evens, odds []int
    for _, num := range nums {
        if num%2 == 0 {
            evens = append(evens, num)
        } else {
            odds = append(odds, num)
        }
    }
    return evens, odds
}
```

---

## 🧠 Summary

| Return Type | Mutable? | Efficient? | Preferred? |
|-------------|----------|------------|------------|
| Array       | ❌ No     | ❌ No       | 🚫 Rare     |
| Slice       | ✅ Yes    | ✅ Yes      | ✅ ✅ ✅     |

---

Want a visual Markdown chart comparing Go, JS, Python, and Java for multiple returns? Just say the word. 😄

