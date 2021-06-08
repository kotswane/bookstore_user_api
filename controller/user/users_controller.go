package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kotswane/bookstore_user_api/domain/user"
	"github.com/kotswane/bookstore_user_api/services"
	"github.com/kotswane/bookstore_user_api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, errUser := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if errUser != nil {
		err := errors.NewBadRequestError("Invalid Userid")
		c.JSON(err.Status, err)
	}

	results, errResults := services.GetUser(userId)
	if errResults != nil {
		c.JSON(errResults.Status, errResults)
	}

	c.JSON(http.StatusOK, results)
}

func CreateUser(c *gin.Context) {
	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json")
		c.JSON(http.StatusBadRequest, restErr)
	}

	results, errResults := services.CreateUser(user)
	if errResults != nil {
		c.JSON(errResults.Status, errResults)
	}

	c.JSON(http.StatusCreated, results)
}
