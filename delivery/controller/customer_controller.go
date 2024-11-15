package controller

import (
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	uc usecase.CustomerUsecase
	r *gin.RouterGroup
}

// RegisterCustomer implements CustomerController.
func (cc *CustomerController) createHandler(c *gin.Context){
	var payload dto.CustomerRequestDto
	err := c.ShouldBindJSON(&payload)	
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : "Error Register Customer : " + err.Error(),
		})
		return 
	}
	if err := cc.uc.RegisterCustomer(payload); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message" : "Error Register Customer : " + err.Error(),
		})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message" : "Success Register Customer : ",
	})
}

func (cc *CustomerController) findByIdHandler(c *gin.Context){
	// Path Param
	// id := c.Param("id")

	// Query Param
	id := c.DefaultQuery("id","100")
	customer, err := cc.uc.FindCustomerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message" : "Error Get Customer : " + err.Error(),
		})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"data" : customer,
		"message" : "Success get Customer by id",
	})
}

func (cc *CustomerController) Route(){
	cc.r.POST("/customers",cc.createHandler)
	cc.r.GET("/customers",cc.findByIdHandler)
}

func NewCustomerController(usecase usecase.CustomerUsecase,rg *gin.RouterGroup) *CustomerController {
	return &CustomerController{
		uc: usecase,
		r : rg,
	}
}
