package main

import (
	"common"
	"fmt"
	//	"gamelog"
	"http"
	"strings"
	// "tcp"
)

func main() {
	fmt.Println("=>", strings.Index("(c)#蘑菇adfadsf", "(c"))

	fmt.Println(common.GetCurrPath())

	//	TestInterface()
	//	TestInterfaceSelect()

	//	test_SetStruct()
	//	test_SetMap()

	//	common.InitReflectParser()
	//	common.LoadCsv("labs\\test.csv")
	//	fmt.Println(common.G_MapCsv)
	//	fmt.Println(common.G_SliceCsv)

	//	test_OOP()
	//	testList()

	//初始化日志系统
	// gamelog.InitLogger("gitSundry", true)
	// gamelog.SetLevel(0)

	// tcp.RegTcpMsgHandler()
	// tcp.NewTcpServer(":9001", 5000)

	http.RegHttpMsgHandler()
	http.NewHttpServer(":9002")
}

func testList() {
	var list []int
	fmt.Println(len(list)) // 0
	if list == nil {       //! 判断通过哟！
		fmt.Println(list) // []
	}
	list = append(list, 22)
	list = append(list, 33)
	fmt.Println(list, list[2:]) // [22,33] []  可以填数组长度哟！！

	for i := 0; i < len(list); i++ {
		if list[i] == 22 {
			list = append(list[:i], list[i+1:]...)
			i--
		}
	}
	fmt.Println(list) // []
}
