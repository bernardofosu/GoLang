# Variable Scope Explained 🌐

Great! Let's break down the concepts of variable scope—especially focusing on local, package/module, and global scope—across Go 🪻, Python 🐍, and JavaScript 🌐.

## 📚 Scope Overview

Scope determines where a variable is accessible in your program. Let’s define the three major scopes you'll encounter:

### 🔸 1. Local Scope

🧠 A variable declared inside a function or block; only available there.

#### **Python Example**
```python
def greet():
    name = "Bernard"  # local to this function
    print(name)

greet()
# print(name) ❌ Error
```

#### **JavaScript Example**
```javascript
function greet() {
  let name = "Bernard"; // local to this function
  console.log(name);
}

greet();
// console.log(name); ❌ Error
```

#### **Go Example**
```go
package main

import "fmt"

func greet() {
    name := "Bernard" // local to greet
    fmt.Println(name)
}
```

### 🔸 2. Package / Module Scope
🧠 A variable declared at the top level of a file; accessible within the same file or package/module.

#### **Python Example (Module-level)**
```python
# module1.py
greeting = "Hello"  # module scope

def greet():
    print(greeting)
```

#### **JavaScript Example (File-level / Module)**
```javascript
// greetings.js
const greeting = "Hello"; // module scope

function greet() {
  console.log(greeting);
}
```

#### **Go Example (Package-level)**
```go
package main

import "fmt"

var greeting = "Hello" // visible throughout the package

func greet() {
    fmt.Println(greeting)
}
```
> ✅ In Go, anything defined outside a function but inside a package is package-scoped.

### 🔸 3. Global Scope
🧠 Accessible everywhere, including across files (if properly imported or declared).

#### **Python Notes**
No true global variables across files unless explicitly imported.
```python
# global.py
value = 42

# main.py
from global import value
print(value)
```

#### **JavaScript Notes**
Browser: Variables declared with `var` at the top level go to the window object.
Node.js: Each file is a module, so you export and import for global-like behavior.
```javascript
// global.js
export const value = 42;

// main.js
import { value } from './global.js';
console.log(value);
```

#### **Go Notes**
Variables are only global to the package unless you export them by capitalizing their names and importing the package elsewhere.
```go
// utils/math.go
package utils

var AddConstant = 5  // exported (global across packages)

// main.go
import "yourmodule/utils"

func main() {
    fmt.Println(utils.AddConstant)
}
```

## 🧠 Summary Table

| Scope          | Python                   | JavaScript                   | Go                                  |
|----------------|--------------------------|------------------------------|-------------------------------------|
| Local          | Inside function or class | Inside function/block        | Inside function                     |
| Module         | File-level in .py file   | File-level (CommonJS/ES6)    | File-level in package               |
| Global         | Across files via import  | Browser: window, Node: export| Capitalized identifiers (exported) |

## ✨ Bonus Notes
- Go is stricter with scope—no “true global” variables across packages unless explicitly exported.
- Python and JS let you bend rules with some tricks, but it’s better to stay modular.
- In all 3, global state is discouraged in large applications—use modules and imports wisely! ✅

---

## 💡 Go Global Variables

In Go, there’s no true “global across the whole app” variable—everything is scoped at the **package level**, and if you want it to be accessible **outside that package**, then:

- Declare it outside of any function (so it's package-level).
- Start its name with a capital letter to make it exported.

### 📆 Example: Package-Level (Global) Variable in Go

#### `mathutils/math.go`
```go
package mathutils

// Exported variable (capitalized)
var AddConstant = 10

// Exported function
func Add(x int, y int) int {
    return x + y + AddConstant
}
```

#### `main.go`
```go
package main

import (
    "fmt"
    "yourmodule/mathutils"  // replace with actual module name
)

func main() {
    sum := mathutils.Add(5, 3)
    fmt.Println("Sum:", sum)

    // Access the package-level variable
    fmt.Println("Constant:", mathutils.AddConstant)
}
```

## 🔁 Recap:

| Visibility     | Name style     | Scope     | Accessible from other packages? |
|----------------|----------------|-----------|---------------------------------|
| Local          | anyCase        | Function  | ❌ No                         |
| Package-local  | lowerCamel     | File/pkg  | ❌ No                         |
| Exported       | Capitalized    | File/pkg  | ✅ Yes                        |

> ✅ So yes, if you want global-like behavior in Go, you:
> - Define at package level
> - Use a capital letter
> - Import that package wherever needed