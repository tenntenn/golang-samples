package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func main() {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 0
	for i:= 1; i <= 10; i++ {
		n := rnd.Intn(100)
		m := rnd.Intn(100)
		fmt.Printf("%2d : %2d + %2d = ", i, n, m)

		var ans int
		r := bufio.NewReader(os.Stdin)
		for {
			line, _, err := r.ReadLine()
			if err != nil {
				fmt.Print(">")
				continue
			}

			str := strings.Replace(string(line), "\r", "", -1)

			ans, err = strconv.Atoi(str)
			if err != nil {
				fmt.Print(">")
				continue
			}

			break
		}

		if ans == n+m {
			fmt.Println(">> correct!")
			count++
		} else {
			fmt.Println(">> wrong!!")
		}
	}

	fmt.Println("-----------------")
	fmt.Println("correct:", count)
}

