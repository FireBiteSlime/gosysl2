package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func Time() {

	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error message: %s\n", err)
	}
	ServerTime := time.Now().Add(response.ClockOffset)
	MachineTime := time.Now()
	fmt.Println(MachineTime)
	fmt.Println(ServerTime)

}

func main() {
	Time()
}
