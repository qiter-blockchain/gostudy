package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:/output.txt")
	if err != nil {
		fmt.Println("open file error = ", err)
		return
	}

	fmt.Println("file = ", *file)

	reader := bufio.NewReader(file) // 默认缓冲4096

	for {
		str, err := reader.ReadString('\n') // 一次读取一行
		if err == nil {
			fmt.Print(str) // reader会把分隔符\n读进去，所以不用Println
		} else if err == io.EOF { // 读到文件尾会返回一个EOF异常
			fmt.Println("文件读取完毕")
			break
		} else {
			fmt.Println("read error: ", err)
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println("close file error = ", err)
	}

}
