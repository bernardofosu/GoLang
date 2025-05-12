# Go Packages Overview

Go’s power really shines through its standard library and community packages. Let’s break down the most used Go packages — both from the standard library and popular external ones — along with real-world use cases.

## 🧰 Top Go Standard Library Packages (Built-in)

| Package            | Purpose / Use Case                               | Example Use                      |
|--------------------|-------------------------------------------------|----------------------------------|
| `fmt`              | Formatting strings, printing to console         | `fmt.Println("Hello")`          |
| `net/http`         | Build HTTP servers and make HTTP requests       | REST APIs, Web servers           |
| `os`               | OS-level operations (files, env vars, args)    | Reading files, env config        |
| `io / io/ioutil`   | Reading/writing streams (e.g., files, requests) | Copy files, stream uploads/downloads |
| `strings`          | String manipulation (split, join, replace, etc.)| Text parsing                     |
| `strconv`          | Convert between strings and numbers              | `strconv.Itoa(10) -> "10"`      |
| `encoding/json`    | Parse and generate JSON                          | Working with APIs, config files  |
| `time`             | Working with time, durations, scheduling        | Timestamps, delays, timers       |
| `math/rand`        | Generate random numbers                          | Simulations, random IDs, games   |
| `context`          | Handle timeouts, deadlines, cancellations       | API calls, goroutines            |

## 📦 Top Community Packages (External)

These are installed using `go get` and widely used in production.

| Package                             | Use Case                                      | Why It’s Popular                 |
|-------------------------------------|-----------------------------------------------|-----------------------------------|
| `github.com/gin-gonic/gin`         | Web framework / REST API                      | Lightweight, fast, and easy to use |
| `github.com/gorilla/mux`           | HTTP request router                           | Flexible URL routing              |
| `github.com/spf13/cobra`           | Build CLI apps                               | Used by Docker, Kubernetes, etc.  |
| `github.com/spf13/viper`           | Manage config files (YAML, JSON, env vars)  | Perfect for config-heavy apps     |
| `go.uber.org/zap`                  | Logging (structured, fast)                   | High-performance logs             |
| `github.com/joho/godotenv`         | Load .env files                              | For local dev environment config  |
| `gorm.io/gorm`                      | ORM (Object Relational Mapping for DBs)     | Easy DB operations with PostgreSQL, MySQL |
| `github.com/dgrijalva/jwt-go`      | Handle JWT tokens                            | Auth system in web apps          |
| `github.com/go-redis/redis`        | Redis client                                 | Caching, pub/sub, session storage |
| `github.com/aws/aws-sdk-go`        | AWS integration                              | Upload to S3, use Lambda, etc.   |

## 💡 Real-World Use Case Examples

| Project                          | Packages Used                             |
|----------------------------------|------------------------------------------|
| 🛒 E-commerce backend API         | `net/http`, `gorm`, `viper`, `jwt-go`   |
| 🔐 Authentication system          | `net/http`, `jwt-go`, `bcrypt`, `gorm`  |
| 🗃 CLI tool for backups           | `os`, `cobra`, `viper`                  |
| 🌐 Static site server             | `net/http`, `fs`, `html/template`       |
| ⚙️ Microservice with Redis        | `go-redis/redis`, `context`, `json`     |
| 📦 File uploader to S3           | `os`, `aws-sdk-go`, `io`, `mime/multipart` |

## 🧠 Summary
- Start with standard packages like `fmt`, `net/http`, `os`, and `time`.
- Then explore external packages when you need something specialized.
- Many Go projects are web servers, CLIs, and microservices — the package ecosystem supports that perfectly.