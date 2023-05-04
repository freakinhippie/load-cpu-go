package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/load"
)

// isPrime checks if a number is prime or not
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// loadCPU checks whether each number starting with 2 is
// prime or not, until it is signalled to stop
func loadCPU(stopChan chan struct{}) {
	i := 2
	for {
		select {
		case <-stopChan:
			return
		default:
			isPrime(i)
			i++
		}
	}
}

func main() {
	// Accept the duration in seconds as a command line flag
	durationPtr := flag.Int("duration", 10, "Duration for which the program should run (in seconds)")
	flag.Parse()

	// Get the number of CPUs on the system
	numCPUs := runtime.NumCPU()

	// Set the GOMAXPROCS to match the number of CPUs
	runtime.GOMAXPROCS(numCPUs)

	// Display the system load average before starting
	loadAvg, err := load.Avg()
	if err != nil {
		fmt.Println("Error getting system load average:", err)
		return
	}
	fmt.Printf("System load: %.2f, %.2f, %.2f\n", loadAvg.Load1, loadAvg.Load5, loadAvg.Load15)

	// Create a WaitGroup to sync all goroutines
	var wg sync.WaitGroup

	// Create a channel to signal all goroutines to stop
	stopChan := make(chan struct{})

	// Run the task on all CPUs
	for i := 0; i < numCPUs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			loadCPU(stopChan)
		}()
	}

	// Sleep for the specified duration
	time.Sleep(time.Duration(*durationPtr) * time.Second)

	// Display the system load average just before closing the channels
	loadAvg, err = load.Avg()
	if err != nil {
		fmt.Println("Error getting system load average:", err)
		return
	}
	fmt.Printf("System load: %.2f, %.2f, %.2f\n", loadAvg.Load1, loadAvg.Load5, loadAvg.Load15)

	// Signal all goroutines to stop
	close(stopChan)

	// Wait for all goroutines to finish
	wg.Wait()
}
