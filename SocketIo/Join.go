/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 23:14
*  @To: 加入房间
 */

package SocketIo

import (
	socketio "github.com/googollee/go-socket.io"
)

type Data struct {
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

func Join(s socketio.Conn, data *Data) {
	s.Emit("join", data)
	S.BroadcastToRoom("/", "user", "notice", s.ID())
}
