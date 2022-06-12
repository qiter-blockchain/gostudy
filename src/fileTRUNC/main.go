package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file_path := "D:/out_go.txt"
	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_TRUNC, 0777)

	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("New content " + fmt.Sprintf("%d", i) + " ....\n")
	}

	writer.Flush()

}
