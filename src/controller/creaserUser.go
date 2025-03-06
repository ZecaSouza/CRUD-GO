package controller

import (
	"github.com/ZecaSouza/CRUD-GO/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreatdUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Requisição invalida")
	c.JSON(err.Code, err)
}