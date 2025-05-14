# 🧾 Case Study: Understanding Pointers in Go

## 📌 Concept: Pass by Value vs. Pass by Reference

In Go, values are passed to functions by value (i.e., a copy). To update the original value, you need to pass a pointer (memory address) using `&`, and receive it using `*` (a pointer type).

---

## ❌ Case Study 1: Passing by Value (No Effect on Original)

### ✅ Code:
```go
package main

import (
	"fmt"
)

func updateAge(agePtr int) {
	agePtr = 99  // 👈 modifies only the copy
}

func main() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	fmt.Println("You are", age, "years old")

	updateAge(age)  // 👈 Pass by value
	fmt.Println("You are now", age, "years old after update")
}
```

### 🧠 What’s Happening?
- `updateAge(age)` passes a copy of `age`.
- Inside `updateAge`, `agePtr = 99` updates the copy, not the original.
- `age` in `main()` remains unchanged.

### 🧪 Walkthrough (If user enters 25):
| Step              | Value         |
|-------------------|---------------|
| age in main()     | 25            |
| updateAge(age)    | agePtr = 25   |
| agePtr = 99       | Changes copy only |
| Final print in main() | age still 25 |

### ✅ Output:
```sql
Enter your age:
25
You are 25 years old
You are now 25 years old after update
```

---

## ✅ Case Study 2: Passing by Reference (Effect on Original)

### ✅ Code:
```go
package main

import (
	"fmt"
)

func updateAge(agePtr *int) {
	*agePtr = 99  // 👈 Dereference and update the original
}

func main() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	fmt.Println("You are", age, "years old")

	updateAge(&age)  // 👈 Pass pointer (reference)
	fmt.Println("You are now", age, "years old after update")
}
```

### 🧠 What’s Happening?
- `updateAge(&age)` passes a pointer to `age` (memory address).
- `*agePtr = 99` updates the value at that address.
- `age` in `main()` gets updated to 99.

### 🧪 Walkthrough (If user enters 25):
| Step              | Value         |
|-------------------|---------------|
| age in main()     | 25            |
| updateAge(&age)   | agePtr → &age |
| *agePtr = 99      | Changes original |
| Final print in main() | age now 99 |

### ✅ Output:
```sql
Enter your age:
25
You are 25 years old
You are now 99 years old after update
```

---

## 🎯 Key Pointer Concepts

| Symbol / Concept | Meaning                         |
|------------------|----------------------------------|
| `int`            | Regular value type – passed by copy |
| `*int`           | Pointer to int – can modify original |
| `&age`           | "Address-of" operator – get pointer to age |
| `*agePtr = 99`   | Dereference – update the value at the address |

---

## 🧠 Pointer Syntax Breakdown

### ✅ Function Declaration
```go
func updateAge(agePtr *int)  // receives a pointer to int
```
- `agePtr` is the name, `*int` is the type

### ❌ Incorrect Go Syntax (like C/C++)
```go
func updateAge(*agePtr int)  // ❌ Syntax error in Go
```

Go follows the **name type** order:
- ✅ `ptr *int`
- ✅ `x []string`
- ✅ `a map[string]int`

Unlike C/C++:
- ✅ `int* ptr` is okay
- ✅ `int *ptr, b;` → only `ptr` is a pointer

---

## 🧠 Two Uses of * — Don’t Confuse Them!

| Context       | Example    | Meaning                    |
|---------------|------------|----------------------------|
| Declaration   | `var p *int` | `p` is a pointer to int     |
| Dereferencing | `*p = 100`   | Set value at pointer's address |

- ✅ In declaration: `*` is part of the type
- ✅ In assignment: `*` is the dereference operator

---

## 🧪 Bonus Tip: Common Pointer Mistakes in Go

| Mistake                                   | Error                                       |
|-------------------------------------------|---------------------------------------------|
| Passing `&x` to a function expecting `int` | cannot use `&x` (type `*int`) as type `int` |
| Using `*varName` in a declaration          | syntax error: unexpected `*varName`         |

---

## 🧩 Summary Table: Go vs. C/C++

| Language | Type Style  | Example         | Notes                                      |
|----------|-------------|-----------------|--------------------------------------------|
| Go       | name type   | `ptr *int`      | Clear: each variable has a full type       |
| C/C++    | type name   | `int* ptr, b;`  | `ptr` is pointer, `b` is not!              |
