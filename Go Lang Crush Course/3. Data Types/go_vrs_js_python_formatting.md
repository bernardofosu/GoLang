# 🐍 Python vs 🦫 Go: String Formatting Placeholders

Yep, absolutely — Python supports placeholders for string formatting just like Go! 🔥  
Let's break it down and even match it side-by-side with Go so you see how similar they are 👇

---

## 🐍 Python Placeholders

Python has **multiple** ways to format strings using placeholders:

### 1. Old-style formatting (% operator)

```python
name = "Bernard"
age = 30
print("Name: %s, Age: %d" % (name, age))
```

- `%s` – string  
- `%d` – integer  
- `%f` – float  
- `%%` – literal percent sign  
```sh
"Name: %s, Age: %d"    # The format string, with placeholders
%                     # The operator that tells Python to insert values
(name, age)           # Tuple with values to fill in: name → %s, age → %d
```

### 2. `str.format()` method

```python
print("Name: {}, Age: {}".format(name, age))
```

---

### 3. f-strings (Python 3.6+ — modern & most recommended)

```python
print(f"Name: {name}, Age: {age}")
```

---

## 🦫 Go Placeholders

Go also uses **format verbs** (aka placeholders) — very similar to Python’s `%` style:

```go
name := "Bernard"
age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

- `%s` – string  
- `%d` – decimal (int)  
- `%f` – float  
- `%v` – default value (any type)  
- `%T` – type  
- `%%` – literal percent sign  

---

## ✅ Side-by-Side Comparison

| Concept              | Python                          | Go              |
|----------------------|----------------------------------|-----------------|
| String placeholder   | `%s`, `{}`, `f"{var}"`           | `%s`            |
| Integer placeholder  | `%d`, `{}`, `f"{var}"`           | `%d`            |
| Float placeholder    | `%.2f`, `f"{var:.2f}"`           | `%.2f`          |
| Any value (generic)  | `f"{var}"`, `"{}".format(var)`   | `%v`            |
| Show type            | `type(var)`                      | `%T`            |
| Literal `%`          | `%%`                             | `%%`            |

---

## 💡 Bonus: Python Version of Go’s `%v`

### In Go:
```go
fmt.Printf("Value: %v\n", myVar)
```

### In Python:
```python
print(f"Value: {myVar}")
# or
print("Value: {}".format(myVar))
```

---

Let me know if you’d like a PDF version of this or want to explore formatting for more data types!
---

## ⚙️ JavaScript Placeholders

Great follow-up! 🤓  
JavaScript doesn't have native `printf`-style formatting like Go (`%s`, `%d`, etc.) or even Python's old-style `%` formatting — but it still gives you powerful ways to use placeholders with variables.

---

### 🚀 String Interpolation (Template Literals) — ES6+ Modern Way

This is the most common and cleanest way to format strings in JavaScript today:

```javascript
let name = "Bernard";
let age = 30;
console.log(`Name: ${name}, Age: ${age}`);
```

- Backticks `` ` ``
- Placeholders `${variable}` inside the string
- ✅ Clean, readable, and powerful. Supports expressions too!

```javascript
console.log(`Next year, Bernard will be ${age + 1} years old.`);
```

---

### 🛠️ Classic Concatenation (Pre-ES6)

```javascript
console.log("Name: " + name + ", Age: " + age);
```

Still works, but it's messier with many variables.

---

### 🧠 Simulating `printf`-style Formatting

JavaScript doesn’t have built-in support for `%s`, `%d`, etc., but libraries like `sprintf-js` and `util.format` (Node.js) fill that gap.

#### ✅ Node.js (built-in)

```javascript
const util = require('util');
let formatted = util.format("Name: %s, Age: %d", "Bernard", 30);
console.log(formatted);  // Output: Name: Bernard, Age: 30
```

#### ✅ Using `sprintf-js` (browser or Node)

```javascript
const sprintf = require("sprintf-js").sprintf;
console.log(sprintf("Name: %s, Age: %d", "Bernard", 30));
```

---

## ⚖️ Summary Table

| Feature               | JavaScript                    | Python                     | Go                          |
|-----------------------|-------------------------------|----------------------------|-----------------------------|
| Template placeholders | ``Hello, ${name}``            | `f"Hello, {name}"`         | `fmt.Sprintf("Hello, %s")`  |
| Print with value      | `console.log(name)`           | `print(name)`              | `fmt.Println(name)`         |
| Old-style formatting  | `util.format("%s", val)` (Node) | `"%s" % val`             | `fmt.Printf("%s", val)`     |
| Type checking         | `typeof name`                 | `type(name)`               | `%T` with `fmt.Printf`      |
| Generic print value   | `${value}`                    | `f"{value}"`               | `%v`                        |


ES6, also known as ECMAScript 2015, was officially released in June 2015. 🎉
🗓️ Summary:

    Full name: ECMAScript 2015 (ES6)

    Released: June 2015

    Nickname: ES6

🚀 Major Features Introduced in ES6:

    let and const (block-scoped variables)

    Arrow functions () => {}

    Template literals `Hello, ${name}`

    Default parameters

    Destructuring (const {x, y} = obj)

    Spread/rest operator (...)

    Classes

    Modules (import / export)

    Promises

    Map and Set