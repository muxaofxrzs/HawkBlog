package im

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type Client struct {
	conn     *websocket.Conn
	username string
	userid   int64
}

var clients = &sync.Map{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandIeWebSocket(w http.ResponseWriter, r *http.Request) {
	//升级http连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	userid := r.Context().Value("userID").(int64)
	username := r.Context().Value("userName").(string)

	// 创建 Client 结构体，并将连接和用户名保存其中
	client := &Client{
		conn:     conn,
		userid:   userid,
		username: username,
	}
	// 将连接添加到clients列表
	clients.Store(client, true)

	// 处理WebSocket连接
	for {
		// 读取消息
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("Received message:", string(msg))

		// 广播消息给所有客户端
		go broadcastMessage(messageType, msg, client)
	}

	// 连接断开时从clients列表移除该连接
	clients.Delete(client)

}

func broadcastMessage(messageType int, message []byte, sender *Client) {
	clients.Range(func(key, value interface{}) bool {
		client := key.(*Client)
		fullMessage := fmt.Sprintf("[%s] %s", sender.username, string(message))

		err := client.conn.WriteMessage(messageType, []byte(fullMessage))
		if err != nil {
			log.Println(err)
			client.conn.Close()
			clients.Delete(client)
		}
		return true
	})

}
