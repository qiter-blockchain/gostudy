package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Y *int32
	Name string
}

// 这是一个基础的创建编码器的用例，编码数据，然后用解码器来接收
func main() {
	// 初始化编码器，并且创建一个解码器的实例
	var network bytes.Buffer        // 标准输入
	enc := gob.NewEncoder(&network) // 编码
	dec := gob.NewDecoder(&network) // 解码
	// 编码器发送数据
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// 解码器接收数据
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	// err = dec.Decode(&q)
	// if err != nil {
	// 	log.Fatal("decode error 2:", err)
	// }
	// fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)

	var q1 P
	err = dec.Decode(&q1)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d,%d}\n", q1.Name, q1.X, q1.Y, q1.Z)
}
