# 📘 Understanding Arrays in Go

An array in Go is a collection of elements of the same data type, stored in fixed-size, ordered memory locations.

## 🔹 Full Example

```go
var nums = [3]int{1, 2, 3}
```

Let’s break it down:

### 🔸 var nums
- `var` is the keyword to declare a variable
- `nums` is the name of the array

### 🔸 [3]
- The square brackets hold the size of the array
- `[3]` means this array has 3 elements
- The size is fixed — you cannot add or remove elements

### 🔸 int
- This is the type of data stored in the array
- So `nums` is an array of 3 integers

### 🔸 {1, 2, 3}
- This is the actual list of values in the array
- The values go inside curly braces `{}`
- Each value must match the type (`int` in this case)

## ✅ What Happens in Memory?

```
Index:     0     1     2
Value:     1     2     3
```

- The values are stored in order
- You can access them using their index: `nums[0]`, `nums[1]`, `nums[2]`

## 🔍 Full Example with Output

```go
package main

import "fmt"

func main() {
    var nums = [3]int{1, 2, 3}
    
    fmt.Println(nums[0]) // Output: 1
    fmt.Println(nums[1]) // Output: 2
    fmt.Println(nums[2]) // Output: 3
    fmt.Println(len(nums)) // to print the size of an array
}
```

Let me know if you want the same breakdown for string or float64 arrays!

---

## ❌ You Cannot Mix Types in an Array

In Go, an array can only hold one data type — all elements must be of the same type.

### 🚫 Invalid Example

```go
var mixed = [3]interface{}{1, "Go", 3.14} // looks like multiple types
```

- This works because of the special type `interface{}` (more on that later).
- But the array itself still holds only `interface{}` types, not actual `int`, `string`, `float64` directly.

### ✅ Valid Arrays (One Type Only)

```go
var nums = [3]int{1, 2, 3}           // All int
var names = [2]string{"Go", "Lang"} // All string
var prices = [2]float64{10.5, 20.0} // All float64
```

If you try to mix types without using `interface{}`, Go will give you an error.

## 🧠 Why?

Go is a statically typed language — types are checked at compile time for safety and speed. So every element in an array must match the declared type.

## 💡 Want to store different types?

You can use a slice or array of `interface{}`, which can hold any type:

```go
var mixed = []interface{}{1, "Go", 3.14}
fmt.Println(mixed[0], mixed[1], mixed[2])
```

But then you need type assertion to work with the values safely.

---

# 📘 Arrays in Go – Quick Notes

## 🔹 1. Declaration

```go
var nums [3]int        // Array of 3 integers
var names [2]string    // Array of 2 strings
var prices [4]float64  // Array of 4 float64s
```

## 🔹 2. Initialization

```go
nums[0] = 10
nums[1] = 20
nums[2] = 30

names[0] = "Go"
names[1] = "Lang"

prices[0] = 12.5
prices[1] = 8.99
```

## 🔹 3. Short Initialization

```go
var nums = [3]int{1, 2, 3}
var names = [2]string{"Hello", "World"}
var prices = [2]float64{99.9, 49.5}
```

## 🔹 4. Let Compiler Infer Size

```go
var nums = [...]int{4, 5, 6}
var names = [...]string{"Go", "Gopher"}
```

## 🔹 5. Access Elements

```go
fmt.Println(nums[0])   // First element
fmt.Println(names[1])  // Second element
```

## 🔹 6. Loop Through Array

```go
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}

OR

for index, value := range prices {
    fmt.Println(index, value)
}
```

## 🔹 7. Array is Fixed Size

```go
var nums [2]int
nums[2] = 30  // ❌ Error: index out of range
```

## ✅ Summary Table

| Type     | Example                          |
|----------|----------------------------------|
| int      | var a [3]int = [3]int{1,2,3}     |
| string   | var s [2]string = {"A","B"}      |
| float64  | var f [2]float64 = {1.1, 2.2}    |
