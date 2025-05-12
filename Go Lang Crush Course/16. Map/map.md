# 📌 Maps in Go 🦫

A **map** in Go is an **unordered collection of key-value pairs**. It's similar to dictionaries in Python 🐍 or objects in JavaScript 🌐.

---

## 🔹 Declaring a Map

```go
var myMap map[string]int // map from string to int
```

This creates a **nil map**. You must initialize it before using it:

```go
myMap = make(map[string]int)
```

Or use shorthand:

```go
myMap := map[string]int{
    "apple":  5,
    "banana": 3,
}
```

---

## 🔹 Accessing and Modifying Maps

```go
myMap["orange"] = 7           // Add or update
value := myMap["apple"]       // Retrieve
delete(myMap, "banana")       // Delete key
```

---

## 🔹 Check if a Key Exists

```go
val, exists := myMap["apple"]
if exists {
    fmt.Println("Apple count:", val)
} else {
    fmt.Println("Apple not found")
}
```

---

## 🔹 Iterating Through a Map

```go
for key, value := range myMap {
    fmt.Println(key, value)
}
```

> ⚠️ Note: Map iteration order is **random**.

---

## 💡 Use Cases for Maps in Go

### ✅ 1. Frequency Counter

Count occurrences of words or numbers:

```go
words := []string{"go", "go", "lang", "code"}
freq := make(map[string]int)
for _, word := range words {
    freq[word]++
}
fmt.Println(freq)
```

### ✅ 2. Lookup Tables

Quickly access values based on keys:

```go
countries := map[string]string{
    "GH": "Ghana",
    "NG": "Nigeria",
}
fmt.Println(countries["GH"]) // Output: Ghana
```

### ✅ 3. Caching Results

Avoid recalculations by caching:

```go
cache := make(map[int]int)
func fib(n int) int {
    if n <= 1 {
        return n
    }
    if val, ok := cache[n]; ok {
        return val
    }
    cache[n] = fib(n-1) + fib(n-2)
    return cache[n]
}
```

---

## 🧠 Mixed-Type Maps Using `interface{}`

To store values of different types:

```go
myMap := map[string]interface{}{
    "count":  10,
    "name":   "Bernard",
    "active": true,
}

fmt.Println(myMap["name"]) // Bernard

// Type assertion
name := myMap["name"].(string)
count := myMap["count"].(int)

// Safer with check
if val, ok := myMap["count"].(int); ok {
    fmt.Println("Count is:", val)
} else {
    fmt.Println("Type mismatch!")
}
```

---

## 🧠 Things to Remember

| Feature       | Go Maps                                  |
|---------------|-------------------------------------------|
| Key types     | Any **comparable** type (string, int, etc) |
| Value types   | Any type                                  |
| Order         | **Random**                                |
| Thread-safe?  | ❌ No (use `sync.Map` for concurrency)   |
| Nil map       | Will panic on write                       |

---

Would you like diagrams 🎟️ showing memory layout or access patterns next?

