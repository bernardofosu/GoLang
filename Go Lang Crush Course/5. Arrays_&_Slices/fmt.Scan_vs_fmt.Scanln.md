# 🥊 `fmt.Scan` vs `fmt.Scanln` in Go — What’s the Difference?

## 📌 Terminal Output Example
```go
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var email string
	fmt.Print("Enter your email: ")
	n, err := fmt.Scanln(&email)

	fmt.Println("n:", n, "err:", err)
	fmt.Println("Email:", email)
}
```

```sh
PS D:\GoLang> go run main.go
Enter your email: nana kwasi
n: 1 err: expected newline
PS D:\GoLang> go run main.go
Enter your email: nana kwasi
n: 1 err: <nil>
PS D:\GoLang>
```

🔍 That clears everything up. You're saying:
- 🟡 First run used fmt.Scanln(&email)
- 🟢 Second run used fmt.Scan(&email)

Same input: nana kwasi

Yet:
```sh
Scanln gave: err: expected newline
Scan gave: err: <nil>
```
- So it means **Scanln** wants to consume everything on the line until Enter (\n), and if there’s extra stuff (like another word), it gets mad and returns err: expected newline.
- Meanwhile, **Scan** just separates values by space, stores the first one, ignores the rest, and doesn't complain.


Let’s break this down clearly — because this is one of the most subtle but important differences in Go input handling!

## 🥊 fmt.Scan vs fmt.Scanln — What's the Real Difference?

| 🧩 Feature             | fmt.Scan                   | fmt.Scanln                          |
|------------------------|----------------------------|--------------------------------------|
| 📥 Reads from          | os.Stdin                   | os.Stdin                             |
| 🔤 Token separator     | Whitespace ( , \t, \n)     | Same — but with extra behavior       |
| ⛔ Stops reading at    | Enough tokens for variables | Same — but expects rest of line empty|
| 🚫 Extra input on line?| ❌ Fine — leaves in buffer  | ⚠️ Returns err: expected newline     |
| ⏎ Newline required?   | ❌ Nope                     | ✅ Yes — must reach newline          |

## 🧪 Live Example

**fmt.Scan**

```go
var email string
n, err := fmt.Scan(&email)
```

You input: `nana kwasi`

- ✅ Scan reads `nana` into `email`, ignores `kwasi` (but leaves it in the buffer)
- ⚠️ No error at all → `err: <nil>`

**fmt.Scanln**

```go
var email string
n, err := fmt.Scanln(&email)
```

You input: `nana kwasi`

- ✅ Scanln reads `nana` into `email`
- ⚠️ BUT it checks for "anything else on the same line"
- ❌ Sees `kwasi` → throws: `err: expected newline`

## 🧠 Why?

Because:

> `Scanln` wants the entire line to be consumed after reading the expected number of values.

It’s meant for **line-based input**, not just space-separated tokens.

## 📦 In Short:

| You want...                    | Use...         | Why                                |
|--------------------------------|----------------|-------------------------------------|
| Just read tokens (space-based) | `fmt.Scan`     | ✅ Ignores extra stuff              |
| Ensure input is clean          | `fmt.Scanln`   | ✅ Validates no leftover input      |
| Read full line (spaces too)    | `bufio.Reader` | ✅ For real line-by-line handling   |