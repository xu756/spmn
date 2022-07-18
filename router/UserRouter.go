/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/17 23:45
*  @Todo:
 */

package router

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine) {
	router := r.Group("/spmn/user")
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	})

}
