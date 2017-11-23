package main

import (
	"fmt"
	"reader-demo/worker"
	"time"
)

func main() {
	start := time.Now().Unix()
	c := make(chan string, 100)
	go worker.Parse(c)

	worker.TailFile("generated.log", c)

	fmt.Print("Time:")
	fmt.Println(time.Now().Unix() - start)

	for index, value := range worker.Datastore {
		fmt.Println(index, len(value))
	}
}
