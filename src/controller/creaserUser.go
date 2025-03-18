package controller

import (
	"fmt"

	"github.com/ZecaSouza/CRUD-GO/src/configuration/rest_err"
	"github.com/ZecaSouza/CRUD-GO/src/configuration/rest_err/model/request"
	"github.com/gin-gonic/gin"
)

func CreatdUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Existem campos incorretos no seu formulário"))
		
		c.JSON(restErr.Code, restErr)
		println(err.Error())
		return
	}
}