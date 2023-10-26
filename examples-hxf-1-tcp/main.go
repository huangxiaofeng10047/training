package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const HEADER_LEN = 4

type Content struct {
	Method    string      `json:"method"`
	NetId     string      `json:"netId"`
	Data      interface{} `json:"data"`
	SN        string      `json:"sn"`
	Timestamp int64       `json:"timestamp"`
}
type Alarm struct {
	AlarmType int    `json:"type"`
	AlarmDesc string `json:"describe"`
	AlarmCode int    `json:"code"`
}

func Packet(method string, netId string, content string) []byte {
	//当前时间戳
	timestamp := time.Now().Unix()
	//将json转为结构体
	var alarm Alarm
	err := json.Unmarshal([]byte(content), &alarm)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
	}
	bytes, _ := json.Marshal(Content{Method: method, NetId: netId, Data: alarm, SN: "21V30000110122B000113", Timestamp: timestamp})
	buffer := make([]byte, HEADER_LEN+len(bytes))
	// 将buffer前面四个字节设置为包长度，大端序
	binary.BigEndian.PutUint32(buffer[0:4], uint32(len(bytes)))
	copy(buffer[4:], bytes)
	return buffer
}

func main() {
	conn, e := net.Dial("tcp", "10.7.10.5:30854")
	if e != nil {
		log.Fatal(e)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')

	//buffer := new(bytes.Buffer)
	buffer := Packet("alarm", "1096683703008911361", text)
	conn.Write(buffer)
	//var message string = ""
	var wbuffer *bufio.Reader = bufio.NewReader(conn)
	fmt.Printf("the default size of buffered reader is %d\n", wbuffer.Size())
	fmt.Printf("The number of unread bytes in the buffer: %d\n", wbuffer.Buffered())
	buf1 := make([]byte, 8*1024*1024)
	for {
		n, err := wbuffer.Read(buf1)
		fmt.Print(string(buf1[:n]))
		//if err == io.EOF {
		//	break
		//}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("The number of unread bytes in the buffer: %d\n", wbuffer.Buffered())
		if wbuffer.Buffered() == 0 {
			break
		}
	}
	n, err := wbuffer.Read(buf1)
	fmt.Println("try again!")
	fmt.Println(n, err)
	// listen for reply
	//rbuf := bufio.NewReaderSize(conn, 8*1024*1024)
	//i := 0
	//for {
	//	context := make([]byte, 21)
	//	n, err := rbuf.Read(context)
	//	// 读取完毕，则跳出
	//	if err != nil {
	//		fmt.Println("读取完毕")
	//		break
	//	}
	//	fmt.Printf("读取内容:%s", context)
	//	if n < (int(21)) {
	//		fmt.Println("err: %s", n < (int(21)))
	//	}
	//	i++
	//}
	//conn.Close()
	//fmt.Print("Message from server: " + message)

	defer conn.Close()
}
