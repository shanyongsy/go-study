package funcfile

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// people = {
//     {
//         name = "John",
//         age = 30,
//         address = {
//             city = "New York",
//             postal_code = "10001",
//         },
//         hobbies = {"reading", "coding", "traveling"},
//     },
//     {
//         name = "Tom",
//         age = 25,
//         address = {
//             city = "San Francisco",
//             postal_code = "94105",
//         },
//         hobbies = {"music", "sports", "photography"},
//     },
//     -- Add more people as needed
// }

type Person struct {
	Name    string
	Age     int
	Address Address
	Hobbies []string
}

type Address struct {
	City       string
	PostalCode string
}

func CallDoString() {
	// 创建 Lua 环境
	L := lua.NewState()
	defer L.Close()

	err := L.DoString(`print("go go go!")`)

	if err != nil {
		panic(err)
	}
}

func CallLuaFunc1() {
	// 创建 Lua 环境
	L := lua.NewState()
	defer L.Close()

	// 加载并执行 Lua 脚本
	if err := L.DoFile("script/stu_lua.lua"); err != nil {
		panic(err)
	}

	// 获取并调用 Lua 全局函数
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("sayHello"),
		NRet:    0,
		Protect: true,
	}, lua.LString("John")); err != nil {
		panic(err)
	}
}

// 加法运算
func CallLuaFunc2() {
	// 创建 Lua 环境
	L := lua.NewState()
	defer L.Close()

	// 加载并执行 Lua 脚本
	if err := L.DoFile("script/stu_lua.lua"); err != nil {
		panic(err)
	}

	// 获取并调用 Lua 全局函数
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("add"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(10), lua.LNumber(32)); err != nil {
		panic(err)
	}

	// 获取返回值
	ret := L.Get(-1)
	L.Pop(1)

	// 打印返回值
	fmt.Println("ret:", ret.String())
}

// 获取配置信息
func CallLuaFunc3() {
	// lua 文件中的配置信息
	// config = {
	// 	name = "John",
	// 	age = 30,
	// 	address = {
	// 		city = "New York",
	// 		postal_code = "10001",
	// 	},
	// 	hobbies = {"reading", "coding", "traveling"},
	// }

	// 创建 Lua 环境
	L := lua.NewState()
	defer L.Close()

	// 加载并执行 Lua 脚本
	if err := L.DoFile("script/stu_lua.lua"); err != nil {
		panic(err)
	}

	// 直接获取配置信息
	config := L.GetGlobal("config").(*lua.LTable)
	name := config.RawGetString("name").String()
	age := config.RawGetString("age").String()

	// 获取嵌套的配置信息
	address := config.RawGetString("address").(*lua.LTable)
	city := address.RawGetString("city").String()
	postalCode := address.RawGetString("postal_code").String()

	// 获取数组配置信息
	hobbies := config.RawGetString("hobbies").(*lua.LTable)
	hobbiesArr := make([]string, 0, hobbies.Len())
	var i int
	for i = 1; i <= hobbies.Len(); i++ {
		hobbiesArr = append(hobbiesArr, hobbies.RawGetInt(i).String())
	}

	// 打印配置信息
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("city:", city)
	fmt.Println("postalCode:", postalCode)
	fmt.Println("hobbies:", hobbiesArr)

}

// 除法运算，获取商和余数
func CallLuaFunc4() {
	// 创建 Lua 环境
	L := lua.NewState()
	defer L.Close()

	// 加载并执行 Lua 脚本
	if err := L.DoFile("script/stu_lua.lua"); err != nil {
		panic(err)
	}

	// 获取并调用 Lua 全局函数
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("div"),
		NRet:    2,
		Protect: true,
	}, lua.LNumber(10), lua.LNumber(3)); err != nil {
		panic(err)
	}

	// 获取返回值
	ret1 := L.Get(-1)
	ret2 := L.Get(-2)
	L.Pop(2)

	// 打印返回值
	fmt.Println("ret1:", ret1.String())
	fmt.Println("ret2:", ret2.String())
}
