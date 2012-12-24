package main

import (
	"os/exec"
	"log"
	"bytes"
	"fmt"
)

func main() {
	cmd := exec.Command("ls", "-l")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}
