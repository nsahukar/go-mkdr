package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	home := os.Getenv("HOME")
	drive := fmt.Sprintf("%s/Drive", home)
	if _, err := os.Stat(drive); os.IsNotExist(err) {
		err := os.Mkdir(drive, 0700)
		if err != nil {
			log.Fatalf("Unable to create 'Drive' directory (at %s), %s\n", home, err)
		}
	}
}
