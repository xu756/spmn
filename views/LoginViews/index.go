/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/18 10:59
*  @ To:
 */

package LoginViews

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/xu756/spmn/models"
	"time"
)

/*--------------数据类型------------*/
// 请求数据

type LoginData struct {
	Data  Empty  `json:"data"`  // 请求数据
	Time  string `json:"time"`  // 时间， 请求时间
	Token string `json:"token"` // 用户唯一， 128位字符串
}
type Empty struct {
	Password string `json:"password"` // 密码，登录密码  前端加密过 64位
	Username string `json:"username"` // 用户名，登录用户名 6-12位
}

// 响应数据

type ResponseData struct {
	Type    string                 `json:"type"`    // 类型， success或error
	Message string                 `json:"message"` // 消息， 提示信息
	Time    string                 `json:"time"`    // 时间， 响应时间
	Data    map[string]interface{} `json:"data"`
}

/*--------------处理函数------------*/

func Login(c *gin.Context) {
	var data LoginData
	c.BindJSON(&data)
	var user models.User
	var db = models.DB
	// 对密码进行sha256加密
	h := sha256.New()
	h.Write([]byte(data.Data.Password))
	data.Data.Password = hex.EncodeToString(h.Sum(nil))
	// 查询数据库
	db.Where("user_name = ? ", data.Data.Username).First(&user)
	// 定义响应数据
	var responseData ResponseData
	responseData.Data = make(map[string]interface{})
	responseData.Time = time.Now().Format("2006-01-02 15:04:05") // 获取当前时间
	if user.Id == 0 {
		responseData.Type = "error"
		responseData.Message = "用户名不存在"
		c.JSON(200, responseData)
		return
	}
	if user.Password != data.Data.Password {
		responseData.Type = "error"
		responseData.Message = "密码错误"

		c.JSON(200, responseData)
		return
	}
	// 生成token
	h = sha512.New()
	h.Write([]byte(user.UserName + user.Password + time.Now().String()))
	token := hex.EncodeToString(h.Sum(nil))
	// 存入数据库
	user.Token = token
	db.Save(&user)
	responseData.Type = "success"
	responseData.Message = "登录成功"
	responseData.Data["token"] = token
	responseData.Data["username"] = user.UserName
	c.JSON(200, responseData)
	return
}
