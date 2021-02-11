package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lindseypoche/SELU_ACM/api/domain/users"
	"github.com/lindseypoche/SELU_ACM/api/services"
	"github.com/lindseypoche/SELU_ACM/api/utils/errors/rest"
)

var (
	// UsersController ...
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Create(*gin.Context)
	Get(*gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type usersController struct{}

func getUserID(userIDParam string) (int64, rest.Err) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64) // param, base, bitSize
	if userErr != nil {
		return 0, rest.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

// Create handles POST requests and creates a new user based on the data in the request
func (c *usersController) Create(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := rest.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.GetStatus(), restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.GetStatus(), saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

// Get handles GET requests and returns a user based on the user id
func (c *usersController) Get(ctx *gin.Context) {
	// get users id from url
	userID, idErr := getUserID(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		ctx.JSON(getErr.GetStatus(), getErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Update handles PUT and PATCH requests and updates a user based on the user id
func (c *usersController) Update(ctx *gin.Context) {
	// get users id from url
	userID, idErr := getUserID(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, idErr)
		return
	}

	// the updated user's data coming in through the request
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := rest.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.GetStatus(), restErr)
		return
	}

	// user id is valid. store it into the user object
	user.ID = userID

	// if method = patch
	isPartial := ctx.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.GetStatus(), err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// Delete handles a DELETE request and deletes a user based on user id
func (c *usersController) Delete(ctx *gin.Context) {
	// get users id from url
	userID, idErr := getUserID(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userID); err != nil {
		ctx.JSON(err.GetStatus(), err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Login logs in user
func (c *usersController) Login(ctx *gin.Context) {
	// get user data from incoming request
	var request users.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		restErr := rest.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.GetStatus(), restErr)
		return
	}

	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		ctx.JSON(err.GetStatus(), err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
