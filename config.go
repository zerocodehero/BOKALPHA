package main

import (
	"fmt"
	"github.com/zerocodehero/BOKALPHA/api"
)

func main() {
	abc := api.LoginParams{
		CustID:   "",
		CompID:   "",
		UserName: "",
		PassWord: "",
		Source:   "",
	}
	fmt.Println(abc.Data())
}
