/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 23:18
*  @To:	向房间内的所有人员发消息
 */

package SocketIo

import (
	"fmt"
)

//type BroadcastData struct {
//	Room string `json:"room"`
//
//}

func Broadcast(c string) {
	S.BroadcastToRoom("/", "cast", "msg", "hello")
	fmt.Println("=========向房间内的所有人员发消息======>")
}
