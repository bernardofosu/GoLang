
# 🌱 Environment Variables in Go

In Go, environment variables like `os.Getenv("PORT")` are only loaded **at runtime** — when the line of code is executed.

---

## 🔍 `os.Getenv()` in Go

✅ Fetches the value of an environment variable  
🕒 Only works at **runtime**  
⚠️ Returns an **empty string** (`""`) if the variable is not set — no error

```go
port := os.Getenv("PORT")
if port == "" {
    port = "3000" // fallback default
}
```

---

## 🧠 Your Config Code

```go
AppConfig = Config{
    Port:     os.Getenv("PORT"),
    MongoURI: os.Getenv("MONGODB_URI"),
    DBName:   os.Getenv("MONGODB_DBNAME"),
}
```

🔁 This configuration loads **once**, typically at startup.  
❗ If the environment variables aren't set **before** starting the app, `AppConfig` will contain empty values.

---

## 🛠️ How to See Environment Variables

### 🖥️ Linux/macOS

```bash
printenv
echo $PORT
```

### 🪟 Windows (CMD)

```cmd
set
```

### 🌐 Node.js

```js
console.log(process.env.PORT);
```

### 🐹 Go

```go
fmt.Println(os.Getenv("PORT"))
```

---

## 🧪 How to Set Env Vars for Testing

### 🔧 Linux/macOS

```bash
export PORT=4000
go run main.go
```

### 🔧 Windows (CMD)

```cmd
set PORT=4000
go run main.go
```

---

## 🧬 About `godotenv.Load()`

```go
err := godotenv.Load()
```

✅ Reads the `.env` file and sets vars into memory  
❌ Does **not** permanently set variables on your system  
📦 Variables exist **only during the app's runtime**

### Example:

```go
func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}
```

---

## 🧠 So Can `os.Getenv("PORT")` Read Host Vars?

✅ Yes! If your OS or container has set them, Go can access them at runtime.

| Env Source             | Readable by `os.Getenv()`? | Example                         |
|------------------------|-----------------------------|---------------------------------|
| `.env` via godotenv    | ✅ Yes                       | `.env: PORT=3000`               |
| Shell/Terminal         | ✅ Yes                       | `export PORT=3000`              |
| Docker/Cloud platform  | ✅ Yes                       | `PORT=3000` in container config |
| Not Set Anywhere       | ❌ No                        | Returns `""`                    |

---

## 💡 Recommended Fallback Pattern

```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080" // default fallback
}
```

---

Would you like a ready-to-use snippet or helper function that does this fallback automatically? 🚀
