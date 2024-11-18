package controller

import (
	"baf-credit-score/delivery/middleware"
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"baf-credit-score/utils/common"
	"baf-credit-score/utils/constant"
	"net/http"
	"strconv"

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
    sizeParam := c.DefaultQuery("size", "3")
    pageParam := c.DefaultQuery("page", "1")

    size, err := strconv.Atoi(sizeParam)
    if err != nil || size <= 0 {        
		common.SendErrorResponse(c,http.StatusBadRequest,"Invalid size parameter")
        return
    }
    
    page, err := strconv.Atoi(pageParam)
    if err != nil || page <= 0 {
        common.SendErrorResponse(c,http.StatusBadRequest,"Invalid page parameter")
        return
    }


    customers, paging,err := cc.uc.FindAll(size,page)
    if err != nil {
        common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }
    var customerItems []any
	for _, customer := range customers {
		customerItems = append(customerItems, customer)
	}
    common.SendPageResponse(c,customerItems,paging,"Success Get All Customer")
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
    cc.r.POST(constant.Customers,cc.authMiddlware.RequireToken(constant.ADMIN), cc.createHandler)
    cc.r.GET(constant.Customers,cc.authMiddlware.RequireToken(constant.ADMIN), cc.listHandler)
    cc.r.GET(constant.CustomersID,cc.authMiddlware.RequireToken(constant.ADMIN), cc.findByIdHandler)
    cc.r.PUT(constant.Customers,cc.authMiddlware.RequireToken(constant.ADMIN), cc.updateByIdHandler)
    cc.r.DELETE(constant.CustomersID,cc.authMiddlware.RequireToken(constant.ADMIN), cc.deleteHandler)
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
