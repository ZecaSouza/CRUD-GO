package controller

import (
	"fmt"
	"log"

	"github.com/ZecaSouza/CRUD-GO/src/configuration/rest_err"
	"github.com/ZecaSouza/CRUD-GO/src/configuration/rest_err/model/request"
	"github.com/ZecaSouza/CRUD-GO/src/configuration/rest_err/model/response"
	"github.com/ZecaSouza/CRUD-GO/src/configuration/rest_err/validation"
	"github.com/gin-gonic/gin"
)

func CreatdUser(c *gin.Context) {
	log.Println("Create User controller")
	var userRequest request.UserRequest

	if error := c.ShouldBindJSON(&userRequest); error != nil {
		log.Println("Error trying to marshal object error: ", error)
		errRest := validation.ValidateUserError(error)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	c.JSON(200, response.UserResponse{
		ID: 	"teste",
		Name: userRequest.Name,
		Email: userRequest.Email,
	})

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Existem campos incorretos no seu formulário"))
		
		c.JSON(restErr.Code, restErr)
		println(err.Error())
		return
	}
}