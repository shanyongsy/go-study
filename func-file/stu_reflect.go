package funcfile

import (
	"fmt"
	"reflect"
)

type Animal struct {
}

func (m *Animal) Eat(sth string) {
	fmt.Println("Eat: " + sth)
}

func RunReflect() {
	animal := Animal{}
	value := reflect.ValueOf(&animal)
	f := value.MethodByName("Eat") //通过反射获取它对应的函数，然后通过call来调用

	if !f.IsValid() || f.Kind() != reflect.Func {
		return
	} else {
		println("f is func")
	}

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf("apple")

	f.Call(in)

}
