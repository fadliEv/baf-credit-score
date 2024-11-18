package controller

import (
	"baf-credit-score/delivery/middleware"
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"baf-credit-score/utils/common"
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
        common.SendErrorResponse(c,http.StatusBadRequest,err.Error())
        return
    }

    if err := cc.uc.RegisterCustomer(payload); err != nil {
        common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }

    common.SendSuccessResponse(c,payload,"Success Register Customer")
}

func (cc *CustomerController) listHandler(c *gin.Context) {
    customers, err := cc.uc.FindAll()
    if err != nil {
        common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }

    common.SendSuccessResponse(c,customers,"Success Get All Customer")
}

func (cc *CustomerController) updateByIdHandler(c *gin.Context) {    
    var payload dto.CustomerRequestDto
    if err := c.ShouldBindJSON(&payload); err != nil {
        common.SendErrorResponse(c,http.StatusBadRequest,err.Error())
        return
    }

    if err := cc.uc.UpdateCustomer(payload); err != nil {
        common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }

    common.SendSuccessResponse(c,payload,"Success Update Customer")
}

func (cc *CustomerController) findByIdHandler(c *gin.Context) {
    id := c.Param("id")
    customer, err := cc.uc.FindCustomerById(id)
    if err != nil {
        common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }

    common.SendSuccessResponse(c,customer,"Success Get Customer By Id")
}

func (cc *CustomerController) deleteHandler(c *gin.Context) {
    id := c.Param("id")
    if err := cc.uc.DeleteCustomer(id); err != nil {
        common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }

    common.SendSuccessResponse(c,nil,"Success Delete Customer")
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
