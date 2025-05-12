# Go Short Variable Declaration (`:=`)

In Go (Golang), the `:=` syntax is called the **short variable declaration**.

🔹 **What it does:**

- It declares and initializes a variable in one step.
- It is shorthand for writing:

```go
var name string = "Bernard"
```
as:

```go
name := "Bernard"
```

🔹 **Rules:**

- Can only be used inside functions (not at package level).
- The variable's type is inferred from the value you assign.
- It cannot be used to re-declare a variable in the same scope unless at least one new variable is being declared.

🔹 **Example:**

```go
package main

import "fmt"

func main() {
    name := "Bernard"     // short variable declaration
    age := 30             // Go infers this as int
    gpa := 3.85           // Go infers this as float64

    fmt.Printf("Name: %s, Age: %d, GPA: %.2f\n", name, age, gpa)
}
```

---

## Rule Examples

### ✅ 1. Can only be used inside functions (not at package level)

❌ **Invalid (Outside any function):**
```go
package main

name := "Bernard"  // ❌ Error: syntax error: non-declaration statement outside function body

func main() {
    fmt.Println(name)
}
```

✅ **Valid (Inside a function):**
```go
package main

import "fmt"

func main() {
    name := "Bernard"  // ✅ OK: short declaration inside a function
    fmt.Println(name)
}
```

### ✅ 2. The variable's type is inferred from the value you assign

```go
package main

import "fmt"

func main() {
    name := "Bernard"   // inferred as string
    age := 30           // inferred as int
    gpa := 3.85         // inferred as float64

    fmt.Printf("%s is %d years old with a GPA of %.2f\n", name, age, gpa)
}
```

Go knows from the number `30` that it’s an `int` — no need to write `var age int = 30`.

### ✅ 3. Cannot be used to re-declare a variable in the same scope (unless at least one new variable)

❌ **Re-declaring only existing variable(s) in same scope:**
```go
package main

func main() {
    name := "Bernard"
    name := "Ama"  // ❌ Error: no new variables on left side of :=
}
```

✅ **At least one new variable on left side:**
```go
package main

import "fmt"

func main() {
    name := "Bernard"
    age := 25

    // OK because 'gpa' is a new variable
    name, gpa := "Ama", 3.75

    fmt.Println(name, age, gpa)  // Output: Ama 25 3.75
}
```

Here, `name` is reused, but `gpa` is new — so Go allows it.

---

## Common Mistakes

| Invalid Syntax | Why it's wrong |
|----------------|----------------|
| `var name := "Bernard"` | ❌ Can't mix `var` and `:=` |
| `var name string := "Bernard"` | ❌ Can't mix type and `:=` |

✅ **Correct Syntax:**

```go
var name string = "Bernard"  // explicit type
var name = "Bernard"         // inferred type
name := "Bernard"            // short declaration, inside function only
```
