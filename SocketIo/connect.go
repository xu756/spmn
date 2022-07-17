/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 23:07
*  @Todo:	连接
 */

package SocketIo

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

var S *socketio.Server

func GetServer(server *socketio.Server) {
	S = server
}

func Connect(s socketio.Conn) error {
	s.SetContext(s.ID() + "连接成功")
	// 申请一个房间
	s.Join("user")
	// 向房间内的所有人员发消息
	fmt.Println("连接成功：", s.ID())
	return nil
}
