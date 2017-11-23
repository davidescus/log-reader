package worker

import (
	"bufio"
	"fmt"
	"os"
)

func TailFile(filePath string, c chan string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c <- scanner.Text()
		// c <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	close(c)
}
