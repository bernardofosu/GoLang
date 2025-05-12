# 🕰️ The `time` Package in Go

The `time` package in Go provides functionality for working with dates, times, and durations. It's used for:
- 📅 Getting the current date and time
- ⏳ Measuring how long something takes
- 💤 Pausing execution
- 🧮 Calculating time differences

---

## 📦 Importing the Package

```go
import "time"
```

## ⏱️ Getting the Current Time

```go
now := time.Now()
fmt.Println("Current Time:", now)
```

📝 This prints the current date and time like:  
`2025-04-21 14:33:07.123456 +0000 UTC`

## 🕰️ Formatting Time

Go uses a reference date `Mon Jan 2 15:04:05 MST 2006` to format time.

```go
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

🧠 Remember: those numbers are not random — they're the reference layout!

## 🧮 Time Difference

```go
start := time.Now()
time.Sleep(2 * time.Second)
elapsed := time.Since(start)

fmt.Println("Elapsed time:", elapsed)
```

✅ This shows the time difference:  
`Elapsed time: 2.002134s`

## ⏰ Delaying Execution

```go
fmt.Println("Waiting for 3 seconds...")
time.Sleep(3 * time.Second)
fmt.Println("Done waiting!")
```

🛑 `time.Sleep()` pauses the program for a specific duration.

## 📊 Duration and Adding Time

```go
duration := 2 * time.Hour
future := now.Add(duration)

fmt.Println("2 hours later:", future)
```

You can use `Add()` to move forward or backward in time.

## 📆 Creating a Specific Date

```go
dob := time.Date(1995, time.January, 1, 0, 0, 0, 0, time.UTC)
fmt.Println("Date of Birth:", dob)
```

You define:  
`Year, Month, Day, Hour, Minute, Second, Nanosecond, Timezone`

## 🤔 Is One Time Before/After Another?

```go
if time.Now().After(dob) {
    fmt.Println("You're older than that date.")
}
```

Use `.Before()`, `.After()`, or `.Equal()` to compare times.

## 📋 Summary Table

| Function        | Description                |
|----------------|----------------------------|
| `time.Now()`   | Get current time 🕒        |
| `time.Sleep()` | Pause execution 💤        |
| `time.Since()` | Get time elapsed ⏱️       |
| `time.Add()`   | Add duration to time ➕    |
| `time.Format()`| Format time as string 🖊️ |
| `time.Parse()` | Convert string to time 🧾  |
| `.Before()`    | Check if time is earlier 🕰️ |
| `.After()`     | Check if time is later ⌛   |
| `.Equal()`     | Check if two times are equal ⚖️ |

## ✅ Example: Timed Task

```go
func main() {
    start := time.Now()

    fmt.Println("Doing something...")
    time.Sleep(2 * time.Second)

    elapsed := time.Since(start)
    fmt.Println("Done! Took:", elapsed)
}
```

---

## 🧠 Duration Units Cheat Sheet

| Unit              | Example                  | Meaning            |
|-------------------|--------------------------|--------------------|
| `time.Second`     | `5 * time.Second`        | 5 seconds ⏱️       |
| `time.Minute`     | `10 * time.Minute`       | 10 minutes 🕐       |
| `time.Hour`       | `2 * time.Hour`          | 2 hours ⏳          |
| `time.Millisecond`| `100 * time.Millisecond` | 0.1 seconds ⚡      |
| `time.Microsecond`| `500 * time.Microsecond` | Very small ⬇️      |
| `time.Nanosecond` | `1 * time.Nanosecond`    | Extremely tiny 🧬  |
