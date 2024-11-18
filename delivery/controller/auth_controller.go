// Registrasi User

package controller

import (
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthController struct {
	uc usecase.AuthenticationUsecase
	r *gin.RouterGroup
}

func (us *AuthController) loginHandler(c *gin.Context){
	var payload dto.LoginDto
	err := c.ShouldBindJSON(&payload)	
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : "Error Login : " + err.Error(),
		})
		return 
	}
	token, err := us.uc.Login(payload.Email,payload.Password); 
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message" : "Error Login : " + err.Error(),
		})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"token" : token,
		"message" : "Success Login User",
	})
}

func (us *AuthController) Route(){
	us.r.POST("/auth/login",us.loginHandler)
}

func NewAuthController(
	usecase usecase.AuthenticationUsecase,
	rg *gin.RouterGroup,
	) *AuthController {
	return &AuthController{
		uc: usecase,
		r : rg,	
	}
}