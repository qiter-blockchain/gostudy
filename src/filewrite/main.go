package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file_path := "D:/out_go.txt"
	// os.OpenFile(name string, flag int, perm FileMode)  参数1：文件路径，参数2：打开模式，参数3：权限控制（windows无效）
	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0777) // 最后的777在windows下没有用

	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("New content" + fmt.Sprintf("%d", i) + "\n") // 写入一行数据
	}

	writer.Flush() // 把缓存数据刷入文件中

}
