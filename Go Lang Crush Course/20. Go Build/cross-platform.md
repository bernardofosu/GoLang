# 🚀 Build Go Binaries for Multiple OSes + 🐳 Docker Packaging

---

## ✅ PART 1: Build Go for Multiple OSes

### 🧱 Basic Cross-Compilation

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o build/myapp-linux

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o build/myapp-mac

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o build/myapp-mac-arm

# Windows
GOOS=windows GOARCH=amd64 go build -o build/myapp-win.exe
```

---

### 🛠️ Bash Script: `build-all.sh`

```bash
#!/bin/bash

mkdir -p build

GOOS=linux   GOARCH=amd64 go build -o build/myapp-linux
GOOS=darwin  GOARCH=amd64 go build -o build/myapp-mac
GOOS=darwin  GOARCH=arm64 go build -o build/myapp-mac-arm
GOOS=windows GOARCH=amd64 go build -o build/myapp-win.exe

echo "✅ All builds done!"
```

### Make it Executable and Run

```bash
chmod +x build-all.sh
./build-all.sh
```

---

## 🐳 PART 2: Build Inside a Docker Container

### 1⃣ Create a `Dockerfile`

> Let’s say your Go app has a `main.go`.

```dockerfile
# Use Go image to build your app
FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN go build -o myapp

# Use a minimal image for running
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/myapp .

CMD ["./myapp"]
```

---

### 2⃣ Build the Docker Image

```bash
docker build -t my-go-app .
```

---

### 3⃣ Run It 🚀

```bash
docker run --rm my-go-app
```

---

## 🤊 BONUS: Static Build (For Tiny Docker Images)

If you're targeting **Alpine** or want a **fully static binary**:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp
```

Then your Dockerfile can be ultra-minimal:

```dockerfile
FROM scratch
COPY myapp /myapp
CMD ["/myapp"]
```

💡 `FROM scratch` creates a zero-layer, no-shell image — **smallest possible container**.

---

## 🧠 Summary

| 🏁 Goal                    | ✅ Command / Tool                          |
|---------------------------|--------------------------------------------|
| Cross-platform binaries   | `GOOS=... GOARCH=... go build`             |
| One-click script          | `build-all.sh`                             |
| Dockerized app            | `docker build -t my-go-app .`              |
| Smallest container        | `FROM scratch` or `FROM alpine` + static   |

---

Need help creating a downloadable release ZIP or GitHub Actions workflow to automate this? Just say the word ✨

