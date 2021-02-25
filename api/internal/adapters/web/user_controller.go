package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
	"github.com/lindseypoche/SELU_ACM/api/internal/utils/errors/rest"
)

// UserController is an interface for a user controller
type UserController interface {
	Create(*gin.Context)
	Get(*gin.Context)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
	// Login(ctx *gin.Context)
}

type userController struct {
	userService domain.UserService
}

// NewUserController creates a new controller for a user
func NewUserController(userService domain.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func getUserID(userIDParam string) (int64, rest.Err) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64) // param, base, bitSize
	if userErr != nil {
		return 0, rest.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

// Create handles POST requests and creates a new user based on the data in the request
func (c *userController) Create(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := rest.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.GetStatus(), restErr)
		return
	}

	result, saveErr := c.userService.CreateUser(&user)
	if saveErr != nil {
		ctx.JSON(saveErr.GetStatus(), saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

// Get handles GET requests and returns a user based on the user id
func (c *userController) Get(ctx *gin.Context) {
	// get users id from url
	userID, idErr := getUserID(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, idErr)
		return
	}

	user, getErr := c.userService.GetUser(userID)
	if getErr != nil {
		ctx.JSON(getErr.GetStatus(), getErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Update handles PUT and PATCH requests and updates a user based on the user id
// func (c *userController) Update(ctx *gin.Context) {
// 	// get users id from url
// 	userID, idErr := getUserID(ctx.Param("user_id"))
// 	if idErr != nil {
// 		ctx.JSON(http.StatusBadRequest, idErr)
// 		return
// 	}

// 	// the updated user's data coming in through the request
// 	var user domain.User
// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		restErr := rest.NewBadRequestError("invalid json body")
// 		ctx.JSON(restErr.GetStatus(), restErr)
// 		return
// 	}

// 	// user id is valid. store it into the user object
// 	user.ID = userID

// 	// if method = patch
// 	isPartial := ctx.Request.Method == http.MethodPatch

// 	result, err := c.userService.UpdateUser(isPartial, user)
// 	if err != nil {
// 		ctx.JSON(err.GetStatus(), err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, result)
// }

// // Delete handles a DELETE request and deletes a user based on user id
// func (c *userController) Delete(ctx *gin.Context) {
// 	// get users id from url
// 	userID, idErr := getUserID(ctx.Param("user_id"))
// 	if idErr != nil {
// 		ctx.JSON(http.StatusBadRequest, idErr)
// 		return
// 	}

// 	if err := c.userService.DeleteUser(userID); err != nil {
// 		ctx.JSON(err.GetStatus(), err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
// }

// // Login logs in user
// func (c *userController) Login(ctx *gin.Context) {
// 	// get user data from incoming request
// 	var request domain.LoginRequest
// 	if err := ctx.ShouldBindJSON(&request); err != nil {
// 		restErr := rest.NewBadRequestError("invalid json body")
// 		ctx.JSON(restErr.GetStatus(), restErr)
// 		return
// 	}

// 	user, err := c.userService.LoginUser(request)
// 	if err != nil {
// 		ctx.JSON(err.GetStatus(), err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, user)
// }
