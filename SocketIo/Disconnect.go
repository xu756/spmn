/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 23:12
*  @To: 断开连接
 */

package SocketIo

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

func Disconnect(s socketio.Conn, msg string) string {
	last := s.Context().(string)
	s.Emit("bye", msg)                     // 回复内容
	fmt.Println("断开连接============>", last) // 获取上一次的内容
	s.Close()
	return last
}
