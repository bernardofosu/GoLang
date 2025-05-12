# 🧱 Structs in Go — Beginner Notes

## 📌 What is a struct?

In Go, a struct (short for structure) is a custom data type that lets you group multiple pieces of data together. Each piece of data in a struct is called a field. These fields can have different data types.

> Think of a struct like an object in JavaScript or a class in Python, but simpler.

---

## ✅ Defining a Struct

To define a struct, you use the `type` and `struct` keywords:

```go
type Person struct {
    Name string
    Age  int
}
```

### 🧠 What this means:

- `type Person struct` → You are creating a new data type called `Person`.
- Inside the curly braces `{}`, you list the fields of the struct.
- `Name string` → A field called `Name` that holds text.
- `Age int` → A field called `Age` that holds a whole number.

---

## 🧪 Creating a Struct (Full form)

Let’s create a `Person` and store it in a variable:

✅ Method 1: Step-by-Step (Long Form)
```go
var person1 Person
person1.Name = "Alice"
person1.Age = 30
```

✅ Method 2: Compact (Struct Literal)
```go
person1 := Person{
    Name: "Alice",
    Age:  30,
}
```

### 🧠 What this does:

- `var person1 Person` → This creates a new variable `person1` of type `Person`.
- Then we set each field individually.

---

## 📝 Shorter way (with explanation)

You can also create and set fields in one line:

```go
person2 := Person{
    Name: "Bob",
    Age:  25,
}
```

- `:=` is short for "create and assign"
- Fields are listed with their names and values
- This is called **composite literal syntax**

---

## 🔁 Comparing with JavaScript and Python

**JavaScript equivalent:**
```js
let person = {
    name: "Alice",
    age: 30
};
```

**Python equivalent:**
```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age

person = Person("Alice", 30)
```

> In Go, you don’t need to write a constructor like `__init__`. Just define the fields!

---

## 📬 Accessing Fields

You can get or change a field using the dot (`.`) operator:

```go
fmt.Println(person2.Name) // prints "Bob"
person2.Age = 26
fmt.Println(person2.Age)  // prints 26
```

---

## 🔁 Pointer to Struct (Advanced but useful)

Sometimes you want to change a struct from inside a function. Then, you pass a pointer:

```go
func haveBirthday(p *Person) {
    p.Age++
}
```

- `*Person` → means you're passing a reference, not a copy.
- `p.Age++` → increases the person's age by 1.

---

## 📦 Summary

| Concept         | Go                          | JavaScript              | Python                 |
|----------------|-----------------------------|--------------------------|------------------------|
| Define          | `type Person struct`        | Object literal           | `class Person:`        |
| Create instance | `Person{Name: ..., Age: ...}` | `{ name: ..., age: ... }` | `Person("Name", age)`  |
| Access fields   | `p.Name`                    | `person.name`            | `person.name`          |

---

## 🧠 Does Go have OOP like Python and JavaScript?

Yes... but it's different.

Go is not a traditional object-oriented programming (OOP) language like Python, Java, or JavaScript.

But Go has some OOP-like features:

| Feature                 | Go                          | Python/JavaScript           |
|------------------------|-----------------------------|-----------------------------|
| Classes                | ❌ No                        | ✅ Yes (`class`)            |
| Inheritance            | ❌ No                        | ✅ Yes (`class A(B)`)       |
| Methods                | ✅ Yes                       | ✅ Yes                      |
| Encapsulation          | ✅ Yes (capitalized fields)  | ✅ Yes (`_name`, `#name`)   |
| Interfaces             | ✅ Yes                       | ✅ (via `abc`, duck typing) |
| Polymorphism           | ✅ Yes                       | ✅ Yes                      |

---

## ✅ What Go *does* have:

### 1. Structs

```go
type Person struct {
    Name string
    Age  int
}
```

### 2. Methods on Structs

You can attach functions (methods) to structs:

```go
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```

Very similar to methods in classes in Python/JS.

### 3. Interfaces

Go uses **interfaces** instead of inheritance.

```go
type Greeter interface {
    Greet()
}
```

Any type that has a `Greet()` method automatically implements the `Greeter` interface — no need to say `"implements Greeter"`.

---

## 🚫 What Go *does not* have:

- ❌ No classes
- ❌ No inheritance
- ❌ No constructors (`__init__`)
- ❌ No `super()` or base class calls
- ❌ No method overloading

Go favors **composition over inheritance**.

```go
type Employee struct {
    Person  // embeds the Person struct
    Role string
}
```

This is how Go "reuses" behavior.

---

## 🧪 Summary

Go supports object-like behavior, but it's not class-based OOP. Instead, it's focused on:

- Simplicity
- Composition over inheritance
- Interfaces for polymorphism
- No hidden magic — everything is explicit
