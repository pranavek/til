package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

// Problem: Cocurrency with just wg is unbounded
// Solution: Creat a pool of worker, who continously looks for some work to do. Use channel to communicate b/w different go routines.
func main() {

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()

			for domain := range jobs {
				_, err := net.LookupHost(domain)
				if err != nil {
					continue
				}

				fmt.Println(domain)
			}
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := sc.Text()
		jobs <- domain
	}
	close(jobs)

	wg.Wait()
}
