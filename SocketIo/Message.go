/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 23:10
*  @To:	接收客户端发送的消息
 */

package SocketIo

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

type Msg struct {
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

func Message(s socketio.Conn, msg *Msg) {
	fmt.Println("消息:", msg)

	s.Emit("msg", msg)

}
