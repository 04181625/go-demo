package user

import "github.com/gin-gonic/gin"

// @Summary 查询用户
// @Produce json
// @Param userName query string false "用户名" maxlength(20)
// @Param userCode query string false "用户代码" maxlength(20)
// @Param page query int true "查询页"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} user.User "成功"
// @Failure 500 {object} error "内部错误"
// @Router /api/v1/user [post]
func List(c *gin.Context) {

}
