package file

import(
	"bufio"
	"fmt"
	"log"
	"os"
)

func Open(fileName string) {
    file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
    if err != nil {
        return
    }
    defer file.Close()

    // read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}