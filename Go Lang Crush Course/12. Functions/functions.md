🧠 What is a Function in Go?

A function in Go is a reusable block of code that performs a specific task. It helps make your code modular, organized, and easier to maintain.

🧱 Basic Function Syntax

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return result
}
```

```go
func functionName(inputParams) (outputParams) {
    // function body
}
```

✅ Example 1: Simple Function (No Parameters, No Return)

```go
func greet() {
    fmt.Println("Welcome to the Booking App! 🛫")
}

// Usage:
greet()
```

🧮 Example 2: Function With Parameters

```go
func bookFlight(name string, destination string) {
    fmt.Printf("Booking flight for %s to %s ✈️\n", name, destination)
}

// Usage:
bookFlight("Bernard", "Paris")
```

📤 Example 3: Function With Return Value

```go
func calculatePrice(days int, ratePerDay int) int {
    total := days * ratePerDay
    return total
}

// Usage:
price := calculatePrice(3, 200)
fmt.Println("Total price:", price)
```

🔁 Example 4: Multiple Return Values

```go
func getBookingInfo() (string, int) {
    name := "Bernard"
    price := 1000
    return name, price
}

// Usage:
customer, amount := getBookingInfo()
fmt.Println("Customer:", customer)
fmt.Println("Amount:", amount)
```

🧠 Why Use Functions?

- Code reusability
- Separation of concerns
- Easy to test
- Readable and maintainable

🧪 Want to Build a Booking App With Functions?

Let me know and I’ll piece together a nice full example that uses multiple functions like:

- greetUser()
- takeBooking(name, dest)
- calculatePrice(days, rate)
- generateReceipt(name, price)

You said:
> and final functions should always be called in the main function

ChatGPT said:

Great question! And you're right on track! ✅

🧩 Should All Functions Be Called in main() in Go?

Yes — in a basic Go program, the execution starts from the main() function. So, any function you define (like bookFlight, calculatePrice, etc.) needs to be either:

- Directly called from main(), or
- Called by another function that main() calls

💡 Analogy

Think of main() as the starting point of your story. All other functions are like scenes in a movie, and the main() function decides which scenes to show and in what order.

🧪 Example: Putting It All Together

```go
package main

import (
	"fmt"
)

// Function to greet the user
func greet() {
	fmt.Println("🎉 Welcome to GoBookings!")
}

// Function to book a flight
func bookFlight(name string, destination string) {
	fmt.Printf("Booking confirmed for %s to %s ✈️\n", name, destination)
}

// Function to calculate price
func calculatePrice(days int, ratePerDay int) int {
	return days * ratePerDay
}

func main() {
	// Call all our helper functions
	greet()

	bookFlight("Bernard", "Cape Town")

	total := calculatePrice(5, 150)
	fmt.Printf("💵 Total cost for 5 days: $%d\n", total)
}
```

🧠 Final Thoughts

- You define your functions above or below main(), but they must eventually be triggered inside main() to run.
- Go won't automatically run any function unless it’s invoked in main() (or by a function that's called from main()).
- This makes your app’s execution path predictable and organized.

You said:
> if function is below the main function can it be called the main function because its declare after the main function

ChatGPT said:

🔥 Great question! And yes — in Go, you can absolutely call a function in main() even if it’s declared after main(). 💡

💡 Why It Works

Go compiles the whole file before execution, so the order of function declarations doesn't matter — as long as they're in the same package (usually main for simple apps).

✅ Example: Calling a Function Declared After main()

```go
package main

import "fmt"

func main() {
	fmt.Println("🚀 Calling function declared below...")
	sayHello()
}

func sayHello() {
	fmt.Println("👋 Hello from below!")
}
```

Output:
```
🚀 Calling function declared below...
👋 Hello from below!
```

🧠 Summary

- Functions can be declared before or after main() — Go doesn’t care.
- Just make sure they’re in the same package (usually main for entry-level programs).
- Go reads the whole file during compilation, so it knows about all functions before running main().

You said:
> because its start from the main function when its hits the function its will called it wether below or after 

ChatGPT said:

🎯 Exactly! You nailed it.

Go always starts execution from the main() function, and as soon as it hits a function call like sayHello(), it knows exactly where to find that function — whether it's declared above or below main() in the code.

🧠 Behind the Scenes:

- When you compile a Go program, the Go compiler reads the entire file (or all files in the package) and builds a map of all the declared functions.
- So when main() is called and it hits sayHello(), the compiler already knows where that function is, even if it’s further down in the file.

✅ Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("🚦 Starting the booking app...")
	checkAvailability()
}

func checkAvailability() {
	fmt.Println("🗕️ Checking if slots are available...")
}
```

The compiler doesn’t care that checkAvailability() is declared after main() — it will still run smoothly.

You said:
> what if its a variable declared below 

ChatGPT said:

🔥 Great follow-up!

Unlike functions, **variables declared *below*** their usage **cannot** be accessed. Go **reads code top-down** when it comes to variables — they must be **declared before they're used**.

⚠️ Example: Invalid Use (Variable Declared After Use)

```go
package main

import "fmt"

func main() {
	fmt.Println(name) // ❌ ERROR: undefined: name
}

var name = "Bernard"
```

➡️ Result:
```
undefined: name
```

Go will not compile this because when it gets to fmt.Println(name), it hasn’t seen name yet.

✅ Correct Use (Variable Declared Before Use)

```go
package main

import "fmt"

var name = "Bernard"

func main() {
	fmt.Println(name) // ✅ Works fine
}
```

➡️ Output:
```
Bernard
```

🧠 TL;DR

| Item        | Can be called before it's declared? |
|-------------|-------------------------------------|
| **Function** | ✅ Yes — Go compiles with full scope knowledge |
| **Variable** | ❌ No — Must be declared before usage |