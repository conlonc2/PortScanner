package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func getBounds() (int, int, string) {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Target : ")
	target, _ := scanner.ReadString('\n')

	fmt.Print("Start Port #: ")
	lowerBound, _ := scanner.ReadString('\n')
	fmt.Print("End Port #: ")
	upperBound, _ := scanner.ReadString('\n')
	lowerBound, upperBound, target = strings.TrimSpace(lowerBound), strings.TrimSpace(upperBound), strings.TrimSpace(target)
	start, _ := strconv.Atoi(lowerBound)
	fin, _ := strconv.Atoi(upperBound)

	return start, fin, target
}

func main() {
	lower, upper, target := getBounds()
	for x := lower; x < upper; x++ {
		wg.Add(1)
		go run(x, target)
	}
	wg.Wait()

}

func run(port int, target string) {

	t := target + ":" + strconv.Itoa(port)
	con, err := net.DialTimeout("tcp", t, 10)
	if err != nil {

	} else {
		con.Close()
		fmt.Printf("Port %v is open\n", port)

	}
	wg.Done()
}
