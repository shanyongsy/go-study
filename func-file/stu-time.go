package funcfile

import (
	"fmt"
	"time"
)

func RunMyTime() {
	t := time.Now()
	t1 := t.UnixMilli()
	t2 := t.UnixMicro() / 1e6
	t3 := t.UnixMicro() / int64(time.Millisecond)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

	fmt.Println("hello")

}
