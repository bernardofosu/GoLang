# Underscore (_) in GO
In Go, the underscore (_) is commonly used in loops to ignore specific values. Here's a detailed breakdown of how the underscore can be used within loops:

## 1. Ignoring the Index in a for-range Loop

When iterating over a collection (such as a slice or array), you may only care about the values and not the indices. In such cases, you can use the underscore (_) to ignore the index.
Example: Ignoring the Index in a for-range Loop
```go
package main

import "fmt"

func main() {
    names := []string{"Nana", "Kwasi", "Ofosu"}
    
    // Ignore the index, only print the value (name)
    for _, name := range names {
        fmt.Println(name)
    }
}
```
Explanation: The loop iterates over the names slice, but the index is not used, so we use _ to discard it, only accessing and printing the name.

## 2. Ignoring the Value in a for-range Loop

If you only care about the indices (or just want to iterate over the collection without using the values), you can use the underscore (_) to ignore the value.
Example: Ignoring the Value in a for-range Loop
```go
package main

import "fmt"

func main() {
    names := []string{"Nana", "Kwasi", "Ofosu"}
    
    // Ignore the value, only print the index
    for index, _ := range names {
        fmt.Println("Index:", index)
    }
}
```
Explanation: Here, the loop iterates over the names slice, but the value (name) is ignored by using _, and only the index is printed.