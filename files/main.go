package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const fileName string = "file.txt"

	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	//Write on file
	size, err := f.Write([]byte("Writing stuff on file...\n"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("File created successfully! Size: %d bytes \n", size)
	f.Close()

	// Read file all at once
	file, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Read file by chunks
	file2, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	
	// Appending new lines to the same file
	file3, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	newLines := []string{"This is a new line", "This is another new line"}

	// stats, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fileSize := stats.Size()

	for _, line := range newLines {
		// var additionalLineBreaker string
		// if (fileSize > 0) {
		// 	additionalLineBreaker = "\n"
		// 	fileSize = 0
		// }
		// _, err := f2.WriteString(additionalLineBreaker + line + "\n")
		_, err := file3.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	file3.Close()

	// Delete file
	err = os.Remove(fileName)
	if err != nil {
		panic(err)
	}
}