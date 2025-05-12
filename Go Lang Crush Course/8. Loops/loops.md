🌀 Loops in Go – Beginner Notes

In Go, a loop helps you repeat a block of code multiple times.
Go uses only one loop keyword: `for`, which makes it simpler than many other languages.

🔁 1. Basic for Loop

```go
package main

import "fmt"

func main() {
    var i int = 0                // Initialization
    for i < 5 {                  // Condition
        fmt.Println(i)           // Loop body
        i = i + 1                // Post statement
    }
}
```

🔹 Explanation:

- `var i int = 0` → This creates a variable `i` of type `int` and sets it to 0.
- `i < 5` → This is the condition. The loop runs as long as `i` is less than 5.
- `i = i + 1` → This increases the value of `i` by 1 after each loop.

🧠 This loop prints:
```
0
1
2
3
4
```

🌀 2. Infinite Loop

```go
package main

import "fmt"

func main() {
    for {
        fmt.Println("Looping forever...")
    }
}
```

🔹 The `for` loop with no condition runs forever. Use `break` to stop it when needed.

🔂 3. Loop Through a Slice using for...range

```go
package main

import "fmt"

func main() {
    var names []string = []string{"Ama", "Kwame", "Kofi"}

    for index, name := range names {
        fmt.Println(index, name)
    }
}
```

🔹 `range` helps you loop through collections like slices.

- `index` → the position (0, 1, 2…)
- `name` → the actual value at that position

🧠 Output:
```
0 Ama
1 Kwame
2 Kofi
```

📝 If you don’t need the index, you can use `_`:

```go
for _, name := range names {
    fmt.Println(name)
}
```

🧭 4. break and continue

🔸 `break` – Stop the loop early

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

🧠 Prints numbers from 0 to 4, then exits.

🔸 `continue` – Skip one round

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
```

🧠 Skips printing 2.

✅ Summary Table

| Loop Type      | Example                          | Use Case                             |
|----------------|----------------------------------|--------------------------------------|
| Basic Loop     | `for i < 5 {}`                   | Repeat code until condition is false |
| Infinite Loop  | `for {}`                         | Keep running until break             |
| Range Loop     | `for i, v := range slice {}`     | Loop through slices or arrays        |
| Break          | `if condition { break }`         | Stop loop early                      |
| Continue       | `if condition { continue }`      | Skip current round                   |

