package main

import (
	"bufio"
	"fmt"
	"os"
	"net"
)

func main() {
	f,_ := os.Open("ip.txt")  //Open a txt file with a list of ips
	scanner := bufio.NewScanner(f)
// Read line by line the file
	for scanner.Scan() {
		line := scanner.Text()
		addr, err := net.LookupAddr(line) // perform reverse lookup
		
		fmt.Println(line, addr, err) 
	}
}