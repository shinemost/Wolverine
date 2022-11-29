package main

import (
	"fmt"
	"hjfu/Wolverine/utils"
)

func main() {
	if err := utils.Init(1); err != nil {
		fmt.Printf("init failed ,err:%v\n", err)
	}
	id := utils.GenId()
	fmt.Println(id)

}
