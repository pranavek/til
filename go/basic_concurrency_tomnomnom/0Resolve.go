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

		_, err := net.LookupHost(domain)
		if err != nil {
			continue
		}

		fmt.Println(domain)
	}
}
