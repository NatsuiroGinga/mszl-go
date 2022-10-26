package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var level = 1
var ex = 0

func main() {
	fmt.Println("输入你的角色名字")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	name := strings.TrimSpace(input)

	fmt.Printf("角色创建成功, %s, 等级为%d\n", name, level)
	s := `你遇到一个怪物, 战斗还是逃跑
		1. 战斗
		2. 逃跑`
	fmt.Println(s)

	for {
		ans, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println(ans)
	}
}
