package controller

import (
	db "courseonline/db/sqlc"
	"courseonline/models"
	"courseonline/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	storedb services.Store
}

func NewCartController(store services.Store) *CartController {
	return &CartController{
		storedb: store,
	}
}

func (ct *CartController) AddToCart(c *gin.Context) {
	var payload models.CartPostReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.GetCartByUserandCourseParams{
		CartUserID:  payload.CartUserID,
		CartCoursID: payload.CartCoursID,
	}

	product, _ := ct.storedb.GetCartByUserandCourse(c, *args)

	var response = &models.CartResponse{}
	var cart = &db.Cart{}
	var err error

	if product == nil || product.CartID == 0 {
		argsAddCart := &db.CreateCartParams{
			CartUserID:  payload.CartUserID,
			CartCoursID: payload.CartCoursID,
			CartPrice:   payload.CartPrice,
			CartQty:     payload.CartQty,
		}
		// create new cart & return
		cart, err = ct.storedb.CreateCart(c, *argsAddCart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.NewError(err))
			return
		}

	} else {
		argsUpdateCart := &db.UpdateCartParams{
			CartID:  product.CartID,
			CartQty: payload.CartQty,
		}
		// update cart quantity
		cart, err = ct.storedb.UpdateCart(c, *argsUpdateCart)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.NewError(err))
			return
		}
	}

	//fetch all list product in carts
	carts, err := ct.storedb.GetCartByUserID(c, cart.CartUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrDataNotFound)
		return
	}

	response.CartID = carts[0].CartID
	response.CartUserID = carts[0].CartUserID
	response.CartCoursID = carts[0].CartCoursID

	//fill carts data to dto response
	for _, v := range carts {
		product := &models.CartCourseResponse{
			CoursID:    &v.CoursID,
			CoursName:  v.CoursName,
			CoursPrice: v.CoursPrice,
			Qty:        v.CartQty,
		}
		response.Course = append(response.Course, product)
	}
	c.JSON(http.StatusCreated, response)
}

func (ct *CartController) DeleteCart(c *gin.Context) error {
	panic("not implemented") // TODO: Implement
}

func (ct *CartController) FindCartByCustomerAndProduct(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ct *CartController) FindCartByCartUserID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Userid := int32(id)
	carts, err := ct.storedb.GetCartByUserID(c, &Userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	var response = &models.CartResponse{}
	response.CartID = carts[0].CartID
	response.CartUserID = carts[0].CartUserID
	response.CartCoursID = carts[0].CartCoursID

	//fill carts data to dto response
	for _, v := range carts {
		product := &models.CartCourseResponse{
			CoursID:    &v.CoursID,
			CoursName:  v.CoursName,
			CoursPrice: v.CoursPrice,
			Qty:        v.CartQty,
		}
		response.Course = append(response.Course, product)
	}
	c.JSON(http.StatusOK, response)
}

func (ct *CartController) FindCartByCustomerPaging(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ct *CartController) UpdateCartQty(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ct *CartController) FindCartById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Carts, err := ct.storedb.GetCartByID(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, Carts)
}

func (ct *CartController) FindAllCart(c *gin.Context) {
	Carts, err := ct.storedb.GetAllCarts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, Carts)
}

func (ct *CartController) CreateCart(c *gin.Context) {
	var payload models.OrderCoursePostReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.CreateOrderCourseParams{
		UscoPurchaseNo: payload.UscoPurchaseNo,
		UscoTax:        payload.UscoTax,
		UscoSubtotal:   payload.UscoSubtotal,
	}

	newCart, err := ct.storedb.CreateCartTx(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, newCart)
}
