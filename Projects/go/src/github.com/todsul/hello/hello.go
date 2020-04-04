package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {
	ntpTime, err := ntp.Time("time.apple.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Network time: %v\n", ntpTime)
	timeFormatted := time.Now().Local()
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)
}

