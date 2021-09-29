package funcfile

import (
	"runtime"
	"strings"
)

func getFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	name := f.Name()

	data := strings.Split(name, ".")
	return data[len(data)-1]
}

func RunFuncName() {
	println(getFuncName())

}
