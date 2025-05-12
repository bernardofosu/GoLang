# 🧑‍💻📅 Beginner-Friendly Guide to Go's `time` Package

Go doesn’t have a separate date package — instead, everything related to dates and time is handled through the time package.

Here’s a beginner-friendly overview with examples and emojis! 🧑‍💻📅

---

## 📦 time Package — Working with Dates 📅

### ✅ Get the Current Date and Time

```go
import "time"

now := time.Now()
fmt.Println("Current time:", now)
```

Example output:

```
Current time: 2025-04-21 18:30:45.123456789 +0000 UTC
```

---

### 🔍 Extract Just the Date or Time

```go
fmt.Println("Year:", now.Year())        // 🗓️
fmt.Println("Month:", now.Month())      // 📅
fmt.Println("Day:", now.Day())          // 🕘

fmt.Println("Hour:", now.Hour())        // 🕐
fmt.Println("Minute:", now.Minute())    // ⏱️
fmt.Println("Second:", now.Second())    // ⏲️
```

---

### 🧮 Formatting Dates and Times

Go uses reference time:

```
Mon Jan 2 15:04:05 MST 2006
```

Yes, it’s weird 😅 — but that’s how you define the format.

```go
fmt.Println(now.Format("2006-01-02"))         // 2025-04-21
fmt.Println(now.Format("02 Jan 2006 15:04"))  // 21 Apr 2025 18:30
```

---

### ⏳ Adding/Subtracting Time

```go
twoDaysLater := now.Add(48 * time.Hour)
fmt.Println("Two days later:", twoDaysLater)

yesterday := now.Add(-24 * time.Hour)
fmt.Println("Yesterday:", yesterday)
```

---

### 📅 Comparing Dates

```go
date1 := time.Now()
date2 := date1.Add(2 * time.Hour)

fmt.Println(date1.Before(date2)) // true ✅
fmt.Println(date1.After(date2))  // false ❌
fmt.Println(date1.Equal(date2))  // false ❌
```

---

### 📆 Creating a Specific Date

```go
customDate := time.Date(2025, time.April, 21, 10, 0, 0, 0, time.UTC)
fmt.Println("My birthday 🎉:", customDate)
```

---

## 📊 Summary Table

| Task                | Function / Method                       |
|---------------------|------------------------------------------|
| Get current time     | `time.Now()` ⏰                          |
| Format date          | `.Format("2006-01-02")` 📄              |
| Add/Subtract duration| `.Add(24 * time.Hour)` ➕➖             |
| Compare times        | `.Before()`, `.After()` 🔍              |
| Extract components   | `.Year()`, `.Month()`, etc. 🧮         |
| Create specific date | `time.Date(...)` 📅                    |