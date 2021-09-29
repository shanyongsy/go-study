package funcfile

import "time"

func RunMyTime() {
	t := time.Now()
	t1 := t.UnixMilli()
	t2 := t.UnixMicro() / 1e6
	t3 := t.UnixMicro() / int64(time.Millisecond)
	println(t1)
	println(t2)
	println(t3)

}
