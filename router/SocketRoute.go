/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/15 21:11
*  @Todo:
 */

package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/xu756/spmn/SocketIo"
	"log"
	"net/http"
)

var S *socketio.Server

func SocketRouter(Server *socketio.Server, r *gin.Engine) {
	// 客户端连接  并加入一个房间
	Server.OnConnect("/", SocketIo.Connect)
	// 接收通知事件
	Server.OnEvent("/", "notice", SocketIo.Notice)
	// 消息
	Server.OnEvent("/", "msg", SocketIo.Message)
	// 客户端断开连接
	Server.OnEvent("/", "bye", SocketIo.Disconnect)
	//加入房间
	Server.OnEvent("/", "join", SocketIo.Join)
	// 连接错误
	Server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("连接错误:", e)

	})
	// 关闭连接
	Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("关闭连接：", s.ID(), reason)
	})
	go func() {
		if err := Server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err) // 启动socket.io服务
		}
	}()
	defer Server.Close()
	r.Use(GinMiddleware("http://localhost:8080"))
	r.GET("/connect/*any", gin.WrapH(Server))
	r.POST("/connect/*any", gin.WrapH(Server))
	r.StaticFS("/public", http.Dir("../asset"))
	err := r.Run(":8000") // 监听端口
	fmt.Println("启动成功")
	if err != nil {
		fmt.Println("启动失败") // 启动失败
		return
	}
}
func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}
