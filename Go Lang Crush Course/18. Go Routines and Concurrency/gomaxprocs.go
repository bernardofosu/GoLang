// package main

// import (
// 	"log"
// 	"runtime"
// 	"time"
// )

// func main() {
// 	// Log the number of CPUs available
// 	numCPUs := runtime.NumCPU()
// 	log.Printf("Number of CPUs available: %d\n", numCPUs)

// 	// Set GOMAXPROCS based on the number of CPUs (can be overridden)
// 	runtime.GOMAXPROCS(numCPUs)
// 	log.Printf("GOMAXPROCS set to: %d\n", numCPUs)

// 	// Start the task with multiple goroutines
// 	start := time.Now()

// 	// Launch multiple goroutines to simulate CPU work
// 	numGoroutines := 10
// 	for i := 0; i < numGoroutines; i++ {
// 		go func(id int) {
// 			// Simulate some CPU-bound work
// 			log.Printf("Goroutine %d started\n", id)
// 			for j := 0; j < 10000000; j++ { // Some CPU-bound task
// 			}
// 			log.Printf("Goroutine %d finished\n", id)
// 		}(i)
// 	}

// 	// Wait for goroutines to finish
// 	time.Sleep(2 * time.Second) // Sleep to allow goroutines to complete

// 	// Log the time taken for the work
// 	duration := time.Since(start)
// 	log.Printf("Total time taken: %s\n", duration)
// }

// package main

// import (
// 	"log"
// 	"runtime"
// 	"time"
// )

// func main() {
// 	// Log the number of CPUs available
// 	numCPUs := runtime.NumCPU()
// 	log.Printf("Number of CPUs available: %d\n", numCPUs)

// 	// Set GOMAXPROCS based on the number of CPUs (can be overridden)
// 	runtime.GOMAXPROCS(numCPUs)
// 	log.Printf("GOMAXPROCS set to: %d\n", numCPUs)

// 	// Start the task
// 	start := time.Now()

// 	// Simulate CPU work without using goroutines (sequential execution)
// 	numTasks := 10
// 	for i := 0; i < numTasks; i++ {
// 		// Simulate some CPU-bound work
// 		log.Printf("Task %d started\n", i)
// 		for j := 0; j < 10000000; j++ { // Some CPU-bound task
// 		}
// 		log.Printf("Task %d finished\n", i)
// 	}

// 	// Log the time taken for the work
// 	duration := time.Since(start)
// 	log.Printf("Total time taken: %s\n", duration)
// }

// package main

// import (
// 	"log"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// func main() {
// 	// Log the number of CPUs available
// 	numCPUs := runtime.NumCPU()
// 	log.Printf("Number of CPUs available: %d\n", numCPUs)

// 	// Set GOMAXPROCS based on the number of CPUs (can be overridden)
// 	runtime.GOMAXPROCS(numCPUs)
// 	log.Printf("GOMAXPROCS set to: %d\n", numCPUs)

// 	// Start the timer
// 	start := time.Now()

// 	// Use WaitGroup to wait for all goroutines
// 	var wg sync.WaitGroup
// 	numGoroutines := 10
// 	wg.Add(numGoroutines)

// 	// Launch multiple goroutines
// 	for i := 0; i < numGoroutines; i++ {
// 		go func(id int) {
// 			defer wg.Done() // Decrement the counter when goroutine completes
// 			log.Printf("Goroutine %d started\n", id)
// 			for j := 0; j < 10000000; j++ {
// 			}
// 			log.Printf("Goroutine %d finished\n", id)
// 		}(i)
// 	}

// 	// Wait for all goroutines to complete
// 	wg.Wait()

// 	// Measure and log time taken
// 	duration := time.Since(start)
// 	log.Printf("Total time taken: %s\n", duration)
// }