# 🐹 Go: Module Name ≠ Folder Name

✅ You can name the folder anything and the module anything else.

> What matters in imports is what’s declared in `go.mod`, not the actual folder name.  
> Go uses the module name in `go.mod` + subfolder for imports.

---

## 🧱 Example:

**Folder Path:**  
`~/Projects/calc-app` ← real folder

**📄 go.mod:**
```
module golang-app
```

**📁 Project structure:**
```
calc-app/
├── go.mod                 → module = golang-app
├── add/add.go            → package add
└── main/start.go         → package main
```

**📄 In start.go:**
```go
import "golang-app/add"
```

✅ Even though the real folder is `calc-app`, Go doesn’t care — it only looks at `golang-app` from `go.mod`.

---

## 🐍 Python and 💻 JavaScript: Folder name does matter

### 📦 Python:
```
project/
├── utils/
│   └── math.py
└── main.py
```

**In main.py:**
```python
from utils import math
```
🧠 You must use the actual folder name (`utils`). No magic module renaming like Go.

---

### 📦 JavaScript (Node.js):
```
project/
├── helpers/
│   └── math.js
└── index.js
```

**In index.js:**
```javascript
const math = require('./helpers/math');
```

🔗 Again, JS uses real folder and file paths.

---

## 🔑 Summary Table

| Feature                          | Go                          | Python / JS         |
|----------------------------------|------------------------------|----------------------|
| Import based on actual folder?   | ❌ Not required               | ✅ Yes               |
| Import based on module name?     | ✅ Uses go.mod module name    | ❌ No                |
| Can folder name ≠ module name?   | ✅ Yes                        | ❌ No                |
| Can run any file directly?       | ❌ Only from package main     | ✅ Yes (with CLI or main guard) |

---

## 🧠 Final Take:

- In Go, the `go.mod` module name defines your import base — not your actual folder name.
- But in Python and JS, you import based on real filesystem paths.
