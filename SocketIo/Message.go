/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 23:10
*  @Todo:	向房间内的所有人员发消息
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
