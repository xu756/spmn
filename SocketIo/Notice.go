/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 23:08
*  @Todo:
 */

package SocketIo

import (
	socketio "github.com/googollee/go-socket.io"
	"log"
)

func Notice(s socketio.Conn, msg string) {
	s.Emit("notice", "have "+msg) // 回复内容
	s.Close()
	log.Println("notice收到内容：:", msg)
}
