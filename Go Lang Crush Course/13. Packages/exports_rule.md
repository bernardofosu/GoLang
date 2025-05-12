## 📦 Go Export Rules

In Go, exported identifiers (functions, variables, structs, etc.) must start with a **capital letter** to be accessible from other packages.

| Name       | Accessible From Other Packages? |
|------------|-------------------------------|
| `Add()`    | ✅ Yes (exported)              |
| `add()`    | ❌ No (unexported/private)     |
| `Result`   | ✅ Yes                         |
| `result`   | ❌ No                          |

### 📁 Example

Let’s say you have this structure:

```
calculator/
├── add/
│   └── add.go
├── main.go
```

#### 🔹 `add/add.go`
```go
package add

func Add(x int, y int) int {
    return x + y
}

func internalHelper() int {
    return 0 // ❌ Not accessible from outside
}
```

#### 🔹 `main.go`
```go
package main

import (
    "fmt"
    "calculator/add"
)

func main() {
    result := add.Add(5, 3)
    fmt.Println(result) // ✅ Works

    // add.internalHelper() ❌ Compiler error: not exported
}
```

### 💡 Why?
Go uses **uppercase to denote public** (exported) and **lowercase for private** (package-only) visibility. There are no `public`, `private`, or `protected` keywords.

---

## 🔍 What About Python and JavaScript?

### 🐍 Python
Python uses **convention**, not enforcement:

- **Public**: No underscore prefix.
```python
def public_function():
    return "I'm public"

class PublicClass:
    def __init__(self):
        self.value = 10
```

- **Private (by convention)**: Single underscore `_`
```python
def _private_function():
    return "I'm private"

class _PrivateClass:
    def __init__(self):
        self._value = 5
```

- **Strictly private** (Name mangling): Double underscore `__`
```python
class MyClass:
    def __init__(self):
        self.__secret = 100
```

### 🌐 JavaScript
JavaScript uses **scoping, conventions**, and **keywords**:

- **Public**:
```js
function publicFunction() {
  return "I'm public";
}

class MyClass {
  constructor() {
    this.value = 10;
  }
  myMethod() {
    return this.value;
  }
}
```

- **Private (ES6+ with #)**:
```js
class MyClass {
  #privateValue;

  constructor() {
    this.#privateValue = 42;
  }

  #privateMethod() {
    return this.#privateValue;
  }
}
```

- **Private (by convention)**:
```js
class MyClass {
  constructor() {
    this._value = 50; // Conventionally private
  }
}
```

### 📦 Modules in JavaScript

#### ✅ ES6 Modules
```js
// math.js
export function add(x, y) {
  return x + y;
}
export const PI = 3.14;

// app.js
import { add, PI } from './math.js';
```

#### ✅ Default Export
```js
// greet.js
export default function greet(name) {
  return `Hello, ${name}!`;
}

// app.js
import greet from './greet.js';
```

#### ✅ CommonJS (Node.js)
```js
// math.js
function add(x, y) {
  return x + y;
}
const PI = 3.14;
module.exports = { add, PI };

// app.js
const math = require('./math');
```

---

## 🧠 Summary Table
| Language   | Public                          | Private                               |
|------------|----------------------------------|----------------------------------------|
| Python     | No underscore                   | `_` (convention), `__` (mangling)       |
| JavaScript| No symbol / `export` / `module.exports` | `_` (convention), `#` (strict private) |
| Go         | Starts with **uppercase**       | Starts with **lowercase**              |

Let me know if you’d like a cheatsheet or real-world usage scenarios! 🚀

