# 💛 Dockerizing Your Gin App Step-by-Step

You already have this structure:

```
your_project/
├── main.go
├── templates/
│   └── hello.html
├── go.mod
├── go.sum
```

---

# 🐋 1. Create a Dockerfile
Inside your project folder, create a new file called **Dockerfile**:

```Dockerfile
# 1. Use the official Golang image as a builder
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your code
COPY . .

# Build the Go app
RUN go build -o myapp main.go

# 2. Create a small final image
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from builder
COPY --from=builder /app/myapp .

# Copy templates folder too (needed for HTML rendering)
COPY --from=builder /app/templates ./templates

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
```

---

# 📜 2. Create a `.dockerignore`
Also create `.dockerignore` to avoid copying unnecessary files:

```
# .dockerignore
go.mod
go.sum
Dockerfile
*.md
*.exe
*.out
vendor/
.git/
```

---

# 🚰 3. Build the Docker Image
Now, open your terminal inside your project folder and run:

```bash
docker build -t gin-app .
```

✅ This will create an image named **gin-app**.

---

# 🚀 4. Run your Docker Container
After building, you can run:

```bash
docker run -p 8080:8080 gin-app
```

✅ Now your Gin app will be accessible at:

```
http://localhost:8080
```

---

# 🧐 Quick Summary

| Step            | Command                         | Meaning               |
|-----------------|----------------------------------|-----------------------|
| Build image     | `docker build -t gin-app .`      | Create Docker image   |
| Run container   | `docker run -p 8080:8080 gin-app`| Map port 8080          |
| App URL         | Open `http://localhost:8080`     | Access your Gin server |

---

# 📦 Final Structure inside container

```
/root/
├── myapp           # compiled app
└── templates/
    └── hello.html
```

---

# ✨ Bonus Tip
If you want **live code reload** during development (like editing Go code and automatic restart inside Docker), I can also show you how to set that up! 🚀

---

