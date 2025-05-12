# 🔁 Go `for` Loop – Beginner Notes

In Go, the **`for` loop** is used to repeat a block of code multiple times. It's the **only loop keyword** in the language, which makes Go simple and clean! 💡

---

## 🧱 Structure of a `for` Loop

```go
package main

import "fmt"

func main() {
    var i int = 0           // Initialization
    for i < 5 {             // Condition
        fmt.Println(i)      // Body
        i = i + 1           // Post Statement
    }
}
```

🔹 Explanation:

| Part           | What it does                              |
|----------------|-------------------------------------------|
| var i int = 0  | Initializes a variable i with the value 0 |
| i < 5          | Loop continues as long as this condition is true |
| i = i + 1      | Increases i by 1 after each loop          |
| fmt.Println(i) | Prints the current value of i             |

🧠 This will output:

```sh
0
1
2
3
4
```

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    var fullName string = "   Nana    Kwasi  Ofosu   "
    var parts []string = strings.Fields(fullName)

    // Create an array to hold the parts
    var nameArray []string

    // Append each part to the array
    for index, name := range parts {
        nameArray = append(nameArray, name)
        fmt.Println("Word", index, ":", name)
    }

    // Print the final array
    fmt.Println("Final name array:", nameArray)
}
```

📌 Key Points

- ✅ `for` is the only looping keyword in Go
- 📏 The loop stops when the condition becomes false
- 🔢 You can use `i = i + 1` or `i++` to increment the counter
- ✍️ Always declare variables fully as a beginner (`var i int = 0`)

>[!NOTE]
>1. Internally, strings in arrays/slices are still strings, which are enclosed in quotes in memory.
>2. Printed Output: Go does not display the quotes around string elements when you print the array or slice itself. It will show the elements without extra quotes.

✅ Summary Table

| Concept       | Explanation                                 |
|---------------|---------------------------------------------|
| for loop      | Repeats code while a condition is true      |
| Full declare  | Use `var i int = 0` for clarity              |
| Stop condition| Controlled using `i < some_number`          |
| Step/increment| Use `i = i + 1` to increase each round       |



# 🔁 Go `for` Loops – Beginner Notes

In Go, the **`for` loop** is the only loop structure you'll need. It can handle traditional loops, infinite loops, and range-based loops.

---

## 1️⃣ Basic `for` Loop

```go
package main

import "fmt"

func main() {
    var i int = 0
    for i < 5 {
        fmt.Println(i)
        i = i + 1
    }
}
```

🔹 Explanation  

| Part           | What it does                           |
|----------------|----------------------------------------|
| var i int = 0  | Start value                            |
| i < 5          | Loop continues while this condition is true |
| i = i + 1      | Increment after each loop iteration    |
| fmt.Println(i) | Output current value                   |

🧾 Output:

```
0
1
2
3
4
```

---

## 2️⃣ for with break 🔚

Use `break` to exit the loop early when a condition is met.

```go
package main

import "fmt"

func main() {
    var i int = 0
    for i < 10 {
        if i == 5 {
            break // Exit loop when i is 5
        }
        fmt.Println(i)
        i = i + 1
    }
}
```

🧾 Output:

```
0
1
2
3
4
```

---

## 3️⃣ for with continue 🔁

Use `continue` to skip the current loop iteration.

```go
package main

import "fmt"

func main() {
    var i int = 0
    for i < 5 {
        i = i + 1
        if i == 3 {
            continue // Skip printing 3
        }
        fmt.Println(i)
    }
}
```

🧾 Output:

```
1
2
4
5
```

(3 is skipped because of `continue`)

---

## 4️⃣ for with range 🔂

Use `range` to loop over elements in a collection (like a slice or array).

```go
package main

import "fmt"

func main() {
    var names = []string{"Ama", "Kojo", "Esi"}
    for index, value := range names {
        fmt.Println("Index:", index, "Value:", value)
    }
}
```

🔹 Explanation  

| Term   | Meaning                         |
|--------|----------------------------------|
| range  | Iterates over each element       |
| index  | The current position (0, 1, 2...) |
| value  | The element at that position     |

🧾 Output:

```
Index: 0 Value: Ama
Index: 1 Value: Kojo
Index: 2 Value: Esi
```

---

## 📌 Summary Table

| Feature    | Purpose                           |
|------------|-----------------------------------|
| `for`      | Main loop keyword in Go           |
| `break`    | Exits the loop early              |
| `continue` | Skips current loop iteration      |
| `range`    | Iterates over collections         |

---

## 🔁 Looping Over a List/Array in Different Languages

### ✅ Python 🐍

```python
names = ["Ama", "Kojo", "Esi"]
for name in names:
    print(name)
```

🔹 Python uses `for item in collection` syntax — very readable and beginner-friendly!

---

### ✅ JavaScript 🟨

```javascript
let names = ["Ama", "Kojo", "Esi"];
for (let name of names) {
    console.log(name);
}
```

🔹 JavaScript uses `for...of` to loop through array values.

---

### ✅ Go 🦫

```go
package main

import "fmt"

func main() {
    var names = []string{"Ama", "Kojo", "Esi"}
    for index, value := range names {
        fmt.Println("Index:", index, "Value:", value)
    }
}
```

🔹 `range` gives both the index and the value.  
You can ignore the index using `_`:

```go
for _, value := range names {
    fmt.Println(value)
}
```

---

## 📌 Summary Table

| Language   | Loop Syntax                     | Keyword Used |
|------------|----------------------------------|---------------|
| Python     | `for item in list:`             | `in`          |
| JavaScript | `for (let item of array)`       | `of`          |
| Go         | `for _, value := range slice`   | `range`       |
