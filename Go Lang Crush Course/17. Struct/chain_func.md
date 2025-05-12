```go
package main

import (
	"fmt"
)

// Define a struct similar to mongo's ClientOptions
type ClientOptions struct {
	URI string
}

// Define a function like options.Client() that returns *ClientOptions
func NewClientOptions() *ClientOptions {
	return &ClientOptions{}
}

// Define a method on *ClientOptions that sets the URI
func (co *ClientOptions) ApplyURI(uri string) *ClientOptions {
	co.URI = uri
	return co
}

func main() {
	// Call function + method (just like options.Client().ApplyURI(...))
	clientOptions := NewClientOptions().ApplyURI("mongodb://localhost:27017")

	// Print the result
	fmt.Println("Mongo URI:", clientOptions.URI)
}
```

🔍 Explanation
Code	What it Represents
NewClientOptions()	Like options.Client()
ApplyURI("...")	Like .ApplyURI(...) on returned struct
ClientOptions	Like MongoDB driver's struct