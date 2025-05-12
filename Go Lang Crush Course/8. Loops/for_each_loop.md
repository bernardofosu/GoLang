# 🧭 Go "For-Each" Loop Notes

In Go, a **for-each loop** is used to iterate over collections like slices, arrays, maps, or strings. Go doesn't have a `foreach` keyword like some other languages (e.g., PHP or C#), but you can achieve the same behavior using the `for` loop with the `range` keyword.

---

## 🔹 Concept: What is a "for-each" loop?

A "for-each" loop is a special kind of loop that goes through each element in a collection (like a list or array) **one at a time**. You don't have to manage index values manually unless you need them.

In Go, we use:

```go
for index, value := range collection {
    // use index and value here
}
```

But for beginners, we’ll avoid shorthand (:=) and use full variable declarations instead.

## 🧪 Basic Example: Iterating Over a Slice

```go
package main

import "fmt"

func main() {
    var names = []string{"Ama", "Kojo", "Esi"}

    var index int
    var value string

    for index, value = range names {
        fmt.Println("Index:", index, "Value:", value)
    }
}
```

### 🔍 Explanation:

- `range names` — tells Go to loop over the names slice.
- `index` — stores the current position (starting from 0).
- `value` — stores the actual element at that index.

## 🔁 Ignoring the Index

If you only care about the values and not the position:

```go
package main

import "fmt"

func main() {
    var names = []string{"Ama", "Kojo", "Esi"}

    var value string

    for _, value = range names {
        fmt.Println("Value:", value)
    }
}
```

- `_` is used to ignore the index.

## 🚩 Using break to Stop Early

```go
package main

import "fmt"

func main() {
    var names = []string{"Ama", "Kojo", "Esi"}

    var index int
    var value string

    for index, value = range names {
        if value == "Kojo" {
            fmt.Println("Found Kojo! Stopping early.")
            break
        }
        fmt.Println("Checking:", value)
    }
}
```

- `break` exits the loop completely when a condition is met.

## 🔄 Using continue to Skip an Item

```go
package main

import "fmt"

func main() {
    var names = []string{"Ama", "Kojo", "Esi"}

    var index int
    var value string

    for index, value = range names {
        if value == "Kojo" {
            fmt.Println("Skipping Kojo...")
            continue
        }
        fmt.Println("Processing:", value)
    }
}
```

- `continue` skips the rest of the loop and goes to the next item.

## 📌 Key Points Summary

| Feature           | Go Syntax                             |
|-------------------|----------------------------------------|
| Basic loop        | `for index, value = range collection`  |
| Ignore index      | `for _, value = range collection`      |
| Stop loop         | `break`                                |
| Skip iteration    | `continue`                             |

### ✅ Great for:

- Looping through lists (slices/arrays)
- Reading values from maps or strings
- Writing clean, readable code without manual index management



# 💡 Go For-Each Loop: Splitting Names Example

Great question! Let’s break it down 👇

You want to:

- Take the string `["Nana Kwasi"]`
- Split the full name "Nana Kwasi" into separate parts: "Nana" and "Kwasi"
- Use a for-each loop in Go to loop over the parts

---

## ✅ Code Example: Splitting and Looping Over Names

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Step 1: Full name as a string
    var fullName string = "Nana Kwasi"

    // Step 2: Split the name into parts using strings.Split
    var names []string = strings.Split(fullName, " ")

    // Step 3: Use a for-each loop to go through each part
    var index int
    var value string

    for index, value = range names {
        fmt.Println("Part", index, ":", value)
    }
}
```

### 🔍 Explanation:

- `strings.Split(fullName, " ")` splits the string into a slice like `["Nana", "Kwasi"]`
- `range names` loops over each part of the name
- `index` is the position (e.g., 0, 1), `value` is the word itself (e.g., "Nana", "Kwasi")
- `fmt.Println(...)` prints each word with its index

### 🧠 Output:

```
Part 0 : Nana
Part 1 : Kwasi
```

---

## ✅ Full Example: Using an Array

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Step 1: Create an array with one full name
    var namesArray [1]string = [1]string{"Nana Kwasi"}

    // Step 2: Get the first (and only) name in the array
    var fullName string = namesArray[0]

    // Step 3: Split the name into parts
    var parts []string = strings.Split(fullName, " ")

    // Step 4: Loop through each part using for-each (range)
    var index int
    var value string

    for index, value = range parts {
        fmt.Println("Part", index, ":", value)
    }
}
```

### 🧠 Output:
```
Part 0 : Nana
Part 1 : Kwasi
```

---

## 🔍 Explanation Table

| Step       | Description                                      |
|------------|--------------------------------------------------|
| Step 1     | Create an array with one full name               |
| Step 2     | Access that name using `namesArray[0]`           |
| Step 3     | Use `strings.Split(fullName, " ")` to split      |
| Step 4     | Loop through the resulting slice with `range`    |

---

## 🔄 strings.Fields vs strings.Split

### 🔍 Difference Between `strings.Fields()` and `strings.Split()`

| Function         | Behavior                                                                 |
|------------------|--------------------------------------------------------------------------|
| `strings.Split()` | Splits based on a specific separator you provide (e.g., a space " ")     |
| `strings.Fields()`| Splits based on **any** whitespace and automatically ignores extra spaces |

### ✅ When to Use `strings.Fields()`
Use when:

- You want to break a sentence into words
- You don’t want to worry about multiple or mixed whitespaces
- You want a cleaner and simpler way to tokenize a string

### 🧪 Example with `strings.Fields()`

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    var fullName string = "   Nana    Kwasi  Ofosu   "
    var parts []string = strings.Fields(fullName)

    for index, name := range parts {
        fmt.Println("Word", index, ":", name)
    }
}
```

### 🧾 Output:
```
Word 0 : Nana
Word 1 : Kwasi
Word 2 : Ofosu
```

✅ **Notice**: it automatically removed all extra spaces.

### 🔁 If You Used `strings.Split(fullName, " ")` Instead
You’d get some empty strings for the extra spaces:
```
["", "", "", "Nana", "", "", "", "Kwasi", "", "Ofosu", "", "", ""]
```

That’s why `strings.Fields()` is your go-to tool for splitting strings by words or natural whitespace 🧼

