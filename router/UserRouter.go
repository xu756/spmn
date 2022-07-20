/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/17 23:45
*  @Todo:
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xu756/spmn/views/LoginViews"
)

func UserRouter(r *gin.Engine) {
	router := r.Group("/spmn/user")
	router.POST("login", LoginViews.Login)
}
