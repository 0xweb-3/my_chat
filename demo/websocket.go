package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 把一个 HTTP 请求 升级为 WebSocket 连接
var upgrader = websocket.Upgrader{}

func serverWs(w http.ResponseWriter, r *http.Request) {
	// 接收一个 HTTP 请求，把它升级为 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil) // nil 可选的响应头
	if err != nil {                          // 升级失败。
		log.Println("Upgrade:", err)
		return
	}

	defer conn.Close()

	for {
		// mt → message type（文本消息 TextMessage 或二进制消息 BinaryMessage）。
		//message → 消息内容（字节数组）。
		//err → 错误。
		mt, message, err := conn.ReadMessage() // 会阻塞直到有数据

		if err != nil { // 如果 err != nil（比如客户端断开连接），就跳出循环
			log.Println("read:", err)
			break
		}

		log.Printf("message:%s", message)    // 打印接收到的消息
		err = conn.WriteMessage(mt, message) //把刚收到的消息原封不动写回去。
		if err != nil {                      // 如果写入失败（例如客户端断开），则退出循环
			log.Println("Write:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", serverWs)
	fmt.Println("启动websocker。。。。。。")
	log.Fatal(http.ListenAndServe("0.0.0.0:1234", nil))
}
