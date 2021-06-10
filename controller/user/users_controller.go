package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kotswane/bookstore_user_api/domain/users"
	"github.com/kotswane/bookstore_user_api/services"
	"github.com/kotswane/bookstore_user_api/utils/errors"
)

func getUserById(paramUserId string) (int64, *errors.RestErr) {
	userId, errUser := strconv.ParseInt(paramUserId, 10, 64)
	if errUser != nil {
		return 0, errors.NewBadRequestError("userId should be a number")
	}
	return userId, nil
}

func Get(c *gin.Context) {
	userId, errUser := getUserById(c.Param("user_id"))
	if errUser != nil {
		c.JSON(errUser.Status, errUser)
	}
	results, errResults := services.GetUser(userId)
	if errResults != nil {
		c.JSON(errResults.Status, errResults)
		return
	}

	c.JSON(http.StatusOK, results)
}

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	results, errResults := services.CreateUser(user)
	if errResults != nil {
		c.JSON(errResults.Status, errResults)
		return
	}

	c.JSON(http.StatusCreated, results)
}

func Update(c *gin.Context) {

	userId, errUser := getUserById(c.Param("user_id"))
	if errUser != nil {
		err := errors.NewBadRequestError("userId should be a number")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	isPartial := c.Request.Method == http.MethodPatch

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	user.Id = userId
	results, errResults := services.UpdateUser(isPartial, user)
	if errResults != nil {
		c.JSON(errResults.Status, errResults)
		return
	}

	c.JSON(http.StatusCreated, results)
}

func Delete(c *gin.Context) {
	userId, errUser := getUserById(c.Param("user_id"))
	if errUser != nil {
		c.JSON(errUser.Status, errUser)
		return
	}

	if errResults := services.DeleteUser(userId); errResults != nil {
		c.JSON(errResults.Status, errResults)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
