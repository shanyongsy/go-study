package main

import (
	"fmt"

	funcfile "github.com/shanyongsy/go-study/func-file"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// funcfile.Run()

	// log.Println("End.")

	// fmt.Println(uuid.New().String())

	// fmt.Println(time.Now().String())

	// funcfile.RunReflect()

	// funcfile.RunMyTime()

	// funcfile.RunFuncName()

	// funcfile.RunSlice()

	// funcfile.RunSayHello()

	funcfile.CallDoString()
	funcfile.CallLuaFunc1()
	funcfile.CallLuaFunc2()
	funcfile.CallLuaFunc3()
	funcfile.CallLuaFunc4()
	funcfile.RunLuaCallback()
	funcfile.RunTableMapping()

}
