# 📂 Folder Structure for the Given Code
```sh
go get github.com/gin-gonic/gin
```

```
your_project/
|
|├── main.go        # <-- Your main Go file
|
└── templates/     # <-- Folder for HTML templates
    └── hello.html # <-- HTML template file
```

---

# 📝 What Each File/Folder is For:

| Name          | Purpose                                         |
|:--------------|:-------------------------------------------------|
| `main.go`     | Your Go application where Gin server is created |
| `templates/`  | Folder to store your `.html` template files      |
| `hello.html`  | HTML file loaded by `LoadHTMLGlob("templates/*")`|

---

# 📜 Example Content:

## `main.go`
```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Load all HTML templates inside templates/ folder
    router.LoadHTMLGlob("templates/*")

    // router.LoadHTMLFiles("templates/index.html", "templates/about.html")


    router.GET("/welcome/:name", func(c *gin.Context) {
        // Get the "name" from the URL path
        name := c.Param("name")

        // Pass it to the HTML template
        c.HTML(200, "hello.html", gin.H{
            "name": name,
        })
    })

    router.Run(":8080")
    // // Start the server on a specific IP and port (e.g., 192.168.1.100:8080)
    // router.Run("192.168.1.100:8080")  // This binds to the specific IP address
}
```

---

## `templates/hello.html`
```html
<!DOCTYPE html>
<html>
<head>
    <title>Welcome</title>
</head>
<body>
    <h1>Hello, {{ .name }}!</h1>
</body>
</html>
```

---

# 🚀 Now Test:

| URL | Output |
|:---|:---|
| `http://localhost:8080/welcome/Alice` | Hello, Alice! |
| `http://localhost:8080/welcome/Bob`   | Hello, Bob!   |

---

# 🔥 Key Points:

- `:name` → path param in URL.
- `c.Param("name")` → extract from URL inside handler.
- `gin.H{"name": name}` → send value to template.
- `{{ .name }}` → Go template syntax to use the value inside HTML.

---

# 📚 Difference Between `c.Data()` and `c.JSON()`:

| Method   | Purpose                       | Content-Type         | Example Content |
|:---------|:-------------------------------|:---------------------|:----------------|
| `c.Data()` | Send any raw data manually    | You must specify it   | HTML, text, files |
| `c.JSON()` | Send JSON data automatically | Sets `application/json` | JSON API response |

---

# 🏯 Super Simple Way to Remember:

| Want to Send | Use |
|:------------|:----|
| JSON (API)  | `c.JSON()` |
| HTML/Text/Files | `c.Data()` |

---

# 🚰 Also, Special Shortcuts in Gin:

| Method | Shortcut For |
|:-------|:-------------|
| `c.String()` | Send plain text |
| `c.HTML()`   | Render HTML template |
| `c.File()`   | Send a file for download |

---

# 📂 Final Line:
- ✅ `c.JSON()` → automatic JSON response.
- ✅ `c.Data()` → flexible manual control.
- ✅ `c.HTML()` → for rendering HTML templates.

---