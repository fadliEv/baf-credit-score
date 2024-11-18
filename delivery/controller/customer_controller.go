package controller

import (
	"baf-credit-score/delivery/middleware"
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
    uc usecase.CustomerUsecase
    r *gin.RouterGroup
    authMiddlware middleware.AuthMiddleware
}

func (cc *CustomerController) createHandler(c *gin.Context) {
    var payload dto.CustomerRequestDto
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "Message": "Error Register Customer : " + err.Error(),
        })
        return
    }

    if err := cc.uc.RegisterCustomer(payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "Message": "Error Register Customer : " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "Message": "Success Register Customer",
    })
}

func (cc *CustomerController) listHandler(c *gin.Context) {
    customers, err := cc.uc.FindAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "Message": "Error Get List Customer : " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":    customers,
        "Message": "Success Get All Customer",
    })
}

func (cc *CustomerController) updateByIdHandler(c *gin.Context) {    
    var payload dto.CustomerRequestDto
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "Message": "Error Update Customer : " + err.Error(),
        })
        return
    }

    if err := cc.uc.UpdateCustomer(payload); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "Message": "Error Update Customer : " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "Message": "Success Update Customer",
    })
}

func (cc *CustomerController) findByIdHandler(c *gin.Context) {
    id := c.Param("id")
    customer, err := cc.uc.FindCustomerById(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "Message": "Error Find Customer : " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":    customer,
        "Message": "Success Find Customer",
    })
}

func (cc *CustomerController) deleteHandler(c *gin.Context) {
    id := c.Param("id")
    if err := cc.uc.DeleteCustomer(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "Message": "Error Delete Customer : " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "Message": "Success Delete Customer",
    })
}

func (cc *CustomerController) Route() {
    cc.r.POST("/customers",cc.authMiddlware.RequireToken("ADMIN"), cc.createHandler)
    cc.r.GET("/customers",cc.authMiddlware.RequireToken("ADMIN"), cc.listHandler)
    cc.r.GET("/customers/:id",cc.authMiddlware.RequireToken("ADMIN"), cc.findByIdHandler)
    cc.r.PUT("/customers",cc.authMiddlware.RequireToken("ADMIN"), cc.updateByIdHandler)
    cc.r.DELETE("/customers/:id",cc.authMiddlware.RequireToken("ADMIN"), cc.deleteHandler)
}

func NewCustomerController(
    usecase usecase.CustomerUsecase,
    r *gin.RouterGroup,
    authMiddleware middleware.AuthMiddleware,
    ) *CustomerController {
    return &CustomerController{
        uc: usecase,
        r:  r,
        authMiddlware: authMiddleware,
    }
}
