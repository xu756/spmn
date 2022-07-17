/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 21:06
*  @Todo:
 */

package router

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/xu756/spmn/SocketIo"
)

var Server *socketio.Server

func InitRouter() {

	r := gin.Default()
	Server = socketio.NewServer(nil)

	r.NoRoute(func(c *gin.Context) {
		// 获取请求的url
		url := c.Request.URL.String()
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "404 not found",
			"url":  url,
		},
		)
	})
	SocketIo.GetServer(Server)
	SocketRouter(Server, r)

}
