package v1

import (
	"github.com/Go-000/Week02/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取标签
// @Router /api/v1/tags [get]
func (t Tag) GetTagById(c *gin.Context) {

	//1。 初始化校验参数
	param := service.GetTagRequest{}

	tagId := c.Param("id")
	ID, err := strconv.Atoi(tagId)
	if err != nil {
		log.Println("get param is fail")
		return
	}
	param.ID = uint32(ID)
	svc := service.New(c)
	tag, errs := svc.GetTag(&param)
	if errs != nil {
		//catch errors
		log.Fatal(errs)
		return

	}
	values := gin.H{"code": 0, "msg": "sucess", "data": tag}
	c.JSON(http.StatusOK, values)
	return
}
