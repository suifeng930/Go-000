package routers

import (
	v1 "github.com/Go-000/Week02/internal/router/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	tag := v1.NewTag()
	apiv1 := engine.Group("/api/v1")
	{
		//tags
		apiv1.GET("/tags", tag.GetTagById)

	}

	return engine

}
