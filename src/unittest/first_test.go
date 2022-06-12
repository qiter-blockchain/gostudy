package main

import (
	"testing" // 引入testing框架
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10) // 调用目标函数
	if res != 45 {
		t.Fatalf("AddUpper(10)执行错误, 期望值%d, 实际值%d\n", 55, res) // 打出错误日志
	}

	t.Logf("AddUpper(10)执行正确") // 打出正常日志
}
