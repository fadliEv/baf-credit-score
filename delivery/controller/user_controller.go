// Registrasi User

package controller

import (
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserController struct {
	uc usecase.UserUsecase
	r *gin.RouterGroup
}

func (us *UserController) createHandler(c *gin.Context){
	var payload dto.UserRequestDto
	err := c.ShouldBindJSON(&payload)	
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : "Error Register User : " + err.Error(),
		})
		return 
	}
	if err := us.uc.RegisterUser(payload); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message" : "Error Register User : " + err.Error(),
		})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message" : "Success Register User",
	})
}

func (us *UserController) Route(){
	us.r.POST("/users",us.createHandler)
}

func NewUserController(
	usecase usecase.UserUsecase,
	rg *gin.RouterGroup,
	) *UserController {
	return &UserController{
		uc: usecase,
		r : rg,	
	}
}