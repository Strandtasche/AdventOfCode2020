package inputloader

import (
	"os"
	"bufio"
	"log"
)

func ReadInput(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

	var lines []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}



