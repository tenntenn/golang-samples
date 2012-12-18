package main

import (
	"syscall"
	"log"
)

func main() {
	if err := syscall.Exec("/bin/ls", []string{"-l"}, []string{}); err != nil {
		log.Fatal(err.Error())
	}
}
