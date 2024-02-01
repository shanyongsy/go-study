package funcfile

import (
	"fmt"
	"sync"
	"time"
)

// 用于缓存消息收发记录
type MsgData struct {
	MsgID     string // 消息ID
	StartTime int64  // 发送消息时间戳 nanoseconds
	EndTime   int64  // 接到返回消息时间戳 nanoseconds
	Sus       bool   // 运行结果
	Fn        bool   // 是否完成
}

type MsgDataManager struct {
	// mutex for JSCallBackFunc
	mutex sync.Mutex
	Datas []*MsgData
}

func (m *MsgDataManager) SendMsg(msg *string) {
	defer func() {
		m.mutex.Unlock()
	}()
	m.mutex.Lock()

	m.Datas = append(m.Datas, &MsgData{MsgID: *msg, StartTime: time.Now().UnixNano()})
}

func (m *MsgDataManager) RecvMsg(msg *string, sus bool) {
	defer func() {
		m.mutex.Unlock()
	}()
	m.mutex.Lock()

	for _, it := range m.Datas {
		if it.Fn {
			continue
		}

		if it.MsgID != *msg {
			continue
		}

		it.EndTime = time.Now().UnixNano()
		it.Fn = true
		it.Sus = sus
	}
}

func RunSlice() {
	var manager MsgDataManager

	for i := 1; i < 10; i++ {
		msg := fmt.Sprintf("msg_%v", i)
		manager.SendMsg(&msg)
	}

	for i := 1; i < 10; i++ {
		msg := fmt.Sprintf("msg_%v", i)
		manager.RecvMsg(&msg, true)
	}

	data := make([]int, 10)
	data1 := data[cap(data):cap(data)]
	for _, i := range data1 {
		fmt.Println(i)
	}
}
