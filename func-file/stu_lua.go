package funcfile

import (
	"fmt"

	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

// 参考： https://blog.csdn.net/xuehu96/article/details/126613678

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

// 无参数无返回值：打印 go go go!
func CallDoString() {
	// 创建 Lua 环境
	L := lua.NewState()
	defer L.Close()

	err := L.DoString(`print("go go go!")`)

	if err != nil {
		panic(err)
	}
}

// 无结果返回：打印 Hello
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

// 单结果返回：加法运算
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

// table 数据获取：获取配置信息
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

// 多结果返回： 除法运算，获取商和余数
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

// lua 的回调函数
func GoAdd(L *lua.LState) int {
	// 获取参数
	arg1 := L.ToInt(1)
	arg2 := L.ToInt(2)

	ret := arg1 + arg2

	// 返回值
	L.Push(lua.LNumber(ret))
	// 返回值的个数
	return 1
}

// lua 回调 Go 函数
func RunLuaCallback() {
	// 创建 Lua 环境
	L := lua.NewState()
	defer L.Close()

	// 注册 Go 函数
	L.SetGlobal("goAdd", L.NewFunction(GoAdd))

	// 加载并执行 Lua 脚本
	if err := L.DoFile("script/stu_lua_callback.lua"); err != nil {
		panic(err)
	}

	// 获取并调用 Lua 全局函数
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("callGoFunc"),
		NRet:    0,
		Protect: true,
	}, lua.LNumber(10), lua.LNumber(32)); err != nil {
		panic(err)
	}
}

// table 映射的方式获取数据
func RunTableMapping() {
	type Role struct {
		Name string
	}

	type Person struct {
		Name      string
		Age       int
		WorkPlace string
		Role      []*Role
	}

	L := lua.NewState()
	if err := L.DoString(`
		person = {
		  name = "Michel",
		  age  = "31", -- weakly input
		  work_place = "San Jose",
		  role = {
			{
			  name = "Administrator"
			},
			{
			  name = "Operator"
			}
		  }
		}
`); err != nil {
		fmt.Println(err)
		// panic(err)
	}
	var person Person
	if err := gluamapper.Map(L.GetGlobal("person").(*lua.LTable), &person); err != nil {
		fmt.Println(err)
		// panic(err)
	}
	fmt.Printf("%s %d\n", person.Name, person.Age)

}
