package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

// Problem: main method not waiting for the go routine to complete
// Solution: Using wait group
// Problem: This cocurrency is unbounded
// Solution: Next program
func main() {
	sc := bufio.NewScanner(os.Stdin)

	var wg sync.WaitGroup
	for sc.Scan() {
		domain := sc.Text()

		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := net.LookupHost(domain)
			if err != nil {
				return
			}

			fmt.Println(domain)
		}()
	}

	wg.Wait()
}
