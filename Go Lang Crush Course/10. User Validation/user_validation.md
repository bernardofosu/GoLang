
# 🧑‍💻 Go (Golang) Input Validation Examples

In Go, you can perform user input validation using `if`, `else if`, and `else`, just like in other C-style languages. Here’s how you can validate different types of user inputs using these structures.

---

## ✅ Example 1: Validate User Age Input

```go
package main

import (
    "fmt"
)

func main() {
    var age int

    fmt.Print("Enter your age: ")
    _, err := fmt.Scan(&age)

    if err != nil {
        fmt.Println("Invalid input! Please enter a number.")
    } else if age < 0 {
        fmt.Println("Age cannot be negative.")
    } else if age < 18 {
        fmt.Println("You are a minor.")
    } else if age < 60 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a senior citizen.")
    }
}
```

🧠 **Explanation:**

- `fmt.Scan(&age)` takes input from the user.
- The `err` is checked to ensure the user didn't enter something invalid (like text instead of a number).
- Multiple conditions are checked using `else if`.
- The `else` block catches everything not covered by the previous conditions.

✅ **Sample Output:**

```
Enter your age: twenty
Invalid input! Please enter a number.

Enter your age: -5
Age cannot be negative.

Enter your age: 25
You are an adult.
```

---

## ✅ Example 2: Validate Username Length Using `len()`

```go
package main

import (
    "fmt"
)

func main() {
    var username string

    fmt.Print("Enter a username: ")
    fmt.Scan(&username)

    if len(username) == 0 {
        fmt.Println("Username cannot be empty.")
    } else if len(username) < 3 {
        fmt.Println("Username must be at least 3 characters long.")
    } else if len(username) > 15 {
        fmt.Println("Username is too long. Max 15 characters allowed.")
    } else {
        fmt.Println("Username is valid!")
    }
}
```

🧠 **How it works:**

- `len(username)` gives the number of characters in the input.
- The conditions check:
    - If it’s empty.
    - If it’s too short.
    - If it’s too long.
    - Or if it’s valid.

✅ **Sample Output:**

```
Enter a username: 
Username cannot be empty.

Enter a username: ab
Username must be at least 3 characters long.

Enter a username: thisisaverylongusername
Username is too long. Max 15 characters allowed.

Enter a username: cool_user
Username is valid!
```

---

## ✅ Example 3: Validate a Password Using Logical Operators

```go
package main

import (
    "fmt"
)

func main() {
    var password string

    fmt.Print("Enter your password: ")
    fmt.Scan(&password)

    if len(password) == 0 {
        fmt.Println("Password cannot be empty.")
    } else if len(password) < 6 || len(password) > 12 {
        fmt.Println("Password must be between 6 and 12 characters.")
    } else if !(len(password) >= 6 && len(password) <= 12) {
        // This is equivalent to the condition above, shown with NOT and AND
        fmt.Println("Invalid password length.")
    } else {
        fmt.Println("Password is valid.")
    }
}
```

✅ **Sample Output:**

```
Enter your password: 
Password cannot be empty.

Enter your password: abc
Password must be between 6 and 12 characters.

Enter your password: abc12345678910
Password must be between 6 and 12 characters.

Enter your password: abc123
Password is valid.
```

💡 **Logical Operators Recap:**

- `A && B` → true if both A and B are true.
- `A || B` → true if at least one of A or B is true.
- `!A` → true if A is false (negation).

---

## ✅ Example 4: Validate a Positive Whole Number for Age

To validate that the user enters a positive integer (i.e., not negative and not a float), you need to:
1. Accept input as a string.
2. Check if it contains only digits (no dots, minus signs, or letters).
3. Convert to `int` safely.
4. Check if it’s non-negative.

```go
package main

import (
    "fmt"
    "strconv"
    "unicode"
)

func isDigitsOnly(s string) bool {
    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}

func main() {
    var input string

    fmt.Print("Enter your age: ")
    fmt.Scan(&input)

    if !isDigitsOnly(input) {
        fmt.Println("Invalid input. Please enter a whole number (no letters, no negative signs, no decimals).")
        return
    }

    age, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Conversion error:", err)
        return
    }

    if age < 0 {
        fmt.Println("Age cannot be negative.")
    } else {
        fmt.Printf("Valid age: %d
", age)
    }
}
```

🧪 **Sample Runs:**

```
Enter your age: -23
Invalid input. Please enter a whole number (no letters, no negative signs, no decimals).

Enter your age: 18.5
Invalid input. Please enter a whole number (no letters, no negative signs, no decimals).

Enter your age: abc
Invalid input. Please enter a whole number (no letters, no negative signs, no decimals).

Enter your age: 23
Valid age: 23
```

---

## 📚 Additional Examples

Want examples for validating strings, email formats, passwords, or adding a loop to keep asking until the input is valid? Let me know! 👨‍💻
