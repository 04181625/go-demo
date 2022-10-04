package user

import (
	_ "fa.com/wms/docs"
	"fa.com/wms/domain/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 仓储物流系统
// @version 1.0
// @description go实现仓储物流系统
// @termsOfService https://blog.csdn.net/weixin_56349119?type=blog
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/api/v1/user", user.List)

	url := ginSwagger.URL("http:localhost:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
}
