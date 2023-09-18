package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

const HEADER_LEN = 4

type Content struct {
	Method string      `json:"method"`
	NetId  string      `json:"netId"`
	Data   interface{} `json:"params"`
}

func Packet(method string, netId string, content string) []byte {
	bytes, _ := json.Marshal(Content{Method: method, NetId: netId, Data: content})
	buffer := make([]byte, HEADER_LEN+len(bytes))
	// 将buffer前面四个字节设置为包长度，大端序
	binary.BigEndian.PutUint32(buffer[0:4], uint32(len(bytes)))
	copy(buffer[4:], bytes)
	return buffer
}

func main() {
	conn, e := net.Dial("tcp", "10.7.20.5:30854")
	if e != nil {
		log.Fatal(e)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')

	//buffer := new(bytes.Buffer)
	buffer := Packet("config_req", "1168743388497420289", text)
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
