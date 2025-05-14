// package main

// import "fmt"

// func main() {
// 	// var name string
// 	// var age int

// 	// fmt.Print("Enter your name and age: ")
// 	// fmt.Scan(&name, &age)
// 	// fmt.Printf("Hello, I am %s and I am %d years old.\n", name, age)

// 	var num int = 42

// 	fmt.Println("Value of num:", num)    // 42
// 	fmt.Println("Address of num:", &num) // Memory address of num
// }

// package main

// import "fmt"

// func main() {

// 	const conferenceTickets int = 50
// 	var remainingTickets uint = 50
// 	conferenceName := "Go Conference"
// 	bookings := []string{}

// 	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	// asking for user input
// 	fmt.Println("Enter Your First Name: ")
// 	fmt.Scanln(&firstName)

// 	fmt.Println("Enter Your Last Name: ")
// 	fmt.Scanln(&lastName)

// 	fmt.Println("Enter Your Email: ")
// 	fmt.Scanln(&email)

// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scanln(&userTickets)

// 	// book ticket in system
// 	remainingTickets = remainingTickets - userTickets
// 	bookings = append(bookings, firstName+" "+lastName)

// 	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
// 	fmt.Printf("These are all our bookings: %v\n", bookings)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var email string
// 	fmt.Print("Enter your email: ")
// 	// fmt.Scanln(&email)
// 	// fmt.Println("You entered:", email)
// 	n, err := fmt.Scan(&email)
// 	fmt.Println("n:", n, "err:", err)

// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	const conferenceTickets int = 50
// 	var remainingTickets uint = 50
// 	conferenceName := "Go Conference"
// 	bookings := []string{}

// 	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

// 	for {
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		// asking for user input
// 		fmt.Println("Enter Your First Name: ")
// 		fmt.Scanln(&firstName)

// 		fmt.Println("Enter Your Last Name: ")
// 		fmt.Scanln(&lastName)

// 		fmt.Println("Enter Your Email: ")
// 		fmt.Scanln(&email)

// 		fmt.Println("Enter number of tickets: ")
// 		fmt.Scanln(&userTickets)

// 		// book ticket in system
// 		remainingTickets = remainingTickets - userTickets
// 		bookings = append(bookings, firstName+" "+lastName)

// 		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
// 		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// 		// print only first names
// 		firstNames := []string{}
// 		for _, booking := range bookings {
// 			var names = strings.Fields(booking)
// 			firstNames = append(firstNames, names[0])
// 		}
// 		fmt.Printf("The first names %v\n", firstNames)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var fullName string = "   Nana    Kwasi  Ofosu   "
// 	var parts []string = strings.Fields(fullName)

// 	// Create an array to hold the parts
// 	var nameArray []string

// 	// Append each part to the array
// 	for index, name := range parts {
// 		nameArray = append(nameArray, name)
// 		fmt.Println("Word", index, ":", name)
// 	}

// 	// Print the final array
// 	fmt.Println("Final name array:", nameArray)
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// 🔧 Create a file named "example.txt"
// 	file, err := os.Create("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close() // 📌 Ensure file is closed after writing

// 	// ✍️ Write text into the file
// 	file.WriteString("Hello, Go File Handling! 🚀")

// 	fmt.Println("✅ File created and text written!")
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Start")

// 	defer fmt.Println("This runs at the END") // ⏳ This is delayed

// 	fmt.Println("Doing stuff...")

// 	// END of function → now deferred line runs
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sayHello() {
// 	fmt.Println("Hello from goroutine!")
// }

// func main() {
// 	go sayHello() // 🔄 runs in the background

// 	fmt.Println("Main function done.")

// 	time.Sleep(1 * time.Second) // ⏳ give goroutine time to finish
// 	fmt.Println("Main function finished after 1 second.")
// }

// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// func main() {
// 	router := gin.Default()

// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	// router.GET("/user", func(c *gin.Context) {
// 	// 	user := map[string]interface{}{
// 	// 		"id":   123,
// 	// 		"name": "Alice",
// 	// 		"roles": []string{
// 	// 			"admin", "user",
// 	// 		},
// 	// 	}
// 	// 	c.JSON(200, gin.H{
// 	// 		"status": "success",
// 	// 		"user":   user,
// 	// 	})
// 	// })

// 	// Load all HTML templates inside templates/ folder
// 	router.LoadHTMLGlob("templates/*")

// 	// router.LoadHTMLFiles("templates/index.html", "templates/about.html")

// 	router.GET("/welcome/:name", func(c *gin.Context) {
// 		// Get the "name" from the URL path
// 		name := c.Param("name")

// 		// Pass it to the HTML template
// 		c.HTML(200, "hello.html", gin.H{
// 			"name": name,
// 		})
// 	})

// 	router.Run(":8080")
// }

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	resp, err := http.Get("https://google.com")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println("Status:", resp.StatusCode)
// 	fmt.Println("Body:", string(body))
// }

// package main

// import (
// 	"fmt"
// )

// // Define a struct similar to mongo's ClientOptions
// type ClientOptions struct {
// 	URI string
// }

// // Define a function like options.Client() that returns *ClientOptions
// func NewClientOptions() *ClientOptions {
// 	return &ClientOptions{}
// }

// // Define a method on *ClientOptions that sets the URI
// func (co *ClientOptions) ApplyURI(uri string) *ClientOptions {
// 	co.URI = uri
// 	return co
// }

// func main() {
// 	// Call function + method (just like options.Client().ApplyURI(...))
// 	clientOptions := NewClientOptions().ApplyURI("mongodb://localhost:27017")

// 	// Print the result
// 	fmt.Println("Mongo URI:", clientOptions.URI)
// }

package main

import (
	"fmt"
)

func updateAge(agePtr *int) {
	*agePtr = 99 // 👈 attempt to update a copy
}

func main() {
	var age int
	fmt.Println("Enter your age:")

	fmt.Scanln(&age)
	fmt.Println("You are", age, "years old")

	updateAge(&age) // 👈 Passing value, not pointer
	fmt.Println("You are now", age, "years old after update")
}
