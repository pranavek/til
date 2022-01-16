package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		domain := sc.Text()

		// We have a problem here, the go routine won't wait for the main method to complete
		go func() {
			_, err := net.LookupHost(domain)
			if err != nil {
				return
			}

			fmt.Println(domain)
		}()
	}
}
