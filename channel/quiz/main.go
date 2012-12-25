package main

import (
	"fmt"
	"math/rand"
	"time"
)

func countdown(n uint) {
	for i := 0; i < int(n); i++ {
		fmt.Print(int(n)-i)
		<-time.After(1 * time.Second)
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Println("Start!!")
}

func main() {

	countdown(uint(5))

	ans := make(chan int)
	go func() {
		for {
			var n int
			fmt.Scanf("%d", &n)
			ans <- n
		}
	}()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 0
	for i:= 1; i <= 10; i++ {
		n := rnd.Intn(100)
		m := rnd.Intn(100)
		fmt.Printf("%2d : %2d + %2d = ", i, n, m)
		select {
		case nm := <-ans:
			if nm == n+m {
				fmt.Println(">> correct!")
				count++
			} else {
				fmt.Println(">> wrong!!")
			}
		case <-time.After(5 * time.Second):
			fmt.Println()
			fmt.Println(">> timed out")
		}
	}

	fmt.Println("-----------------")
	fmt.Println("correct:", count)
}

