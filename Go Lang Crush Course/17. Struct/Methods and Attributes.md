# 🚀 Go: Methods and Attributes Explained

In Go, **attributes** are fields inside a `struct`, and **methods** are functions associated with that struct (via a receiver).

---

## ✅ Example: Methods and Attributes in Go

```go
package main

import "fmt"

// Define a struct with attributes (fields)
type User struct {
	Name  string
	Email string
	Age   int
}

// Method to greet the user (reads attributes)
func (u User) Greet() string {
	return fmt.Sprintf("Hello, %s!", u.Name)
}

// Method to update the email (modifies an attribute)
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

// Method to check if the user is an adult
func (u User) IsAdult() bool {
	return u.Age >= 18
}

func main() {
	// Create a User instance
	user := User{Name: "Alice", Email: "alice@example.com", Age: 22}

	// Access attributes
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)

	// Call methods
	fmt.Println(user.Greet())           // → Hello, Alice!
	fmt.Println("Is adult?", user.IsAdult()) // → true

	// Update email using method
	user.UpdateEmail("alice@newdomain.com")
	fmt.Println("Updated Email:", user.Email)
}
```

---

## 🔍 Breakdown

| Concept              | Go Example             | Meaning                           |
|----------------------|-------------------------|-----------------------------------|
| Attribute            | `Name`, `Email`, `Age` | Struct fields (data)              |
| Method               | `Greet()`, `UpdateEmail()` | Functions tied to the struct     |
| Receiver             | `(u User)` or `(u *User)` | Binds method to `User` type      |
| Pointer receiver     | `*User`                | Allows modifying struct fields    |

---

## 📥 Input Parameter vs Receiver

In this method:

```go
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}
```
- `u *User` is the **receiver** – it's not an input like a parameter, but it tells Go that this method belongs to `User`.
- `newEmail string` is the actual **input**.
- When calling it like `user.UpdateEmail("new@domain.com")`, Go implicitly passes a pointer to `user`.

---

## 🤔 Why `Greet()` Uses a Copy

```go
func (u User) Greet() string {
	return fmt.Sprintf("Hello, %s!", u.Name)
}
```

This uses a **value receiver** because:

✅ It's just reading data (no mutation).  
✅ Small structs are cheap to copy.  
✅ It's semantically clear: "Greet" shouldn't change anything.

---

## 🧠 When to Use What

| Use Case                   | Value Receiver `(u User)` | Pointer Receiver `(u *User)` |
|----------------------------|----------------------------|-------------------------------|
| Just reading data          | ✅                          | 🟡 Optional                   |
| Modifying struct fields    | ❌                          | ✅                            |
| Avoiding large copies      | ❌                          | ✅                            |
| Consistent method set      | 🟡 Recommended if others use pointers |

---

## ✅ Summary

- Attributes = struct fields. 🏷️
- Methods = functions bound to structs. 🛠️
- Pointer receivers allow changes. 🔁
- Value receivers are simple and safe for read-only. 📖

Would you like to explore method chaining or embedding structs next? 😊
