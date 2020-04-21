// WaitForServer attempts to contact the server of a URl.
// Tries for one minute using exponential back-off
// Reports an error if all attempts fail.



package main

import (
	"time"
	"fmt"
)

// func main() {
// 	fmt.Println("Running...")
// }

func WaitForSever(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %", url, timeout)
}