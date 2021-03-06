package users

import (
	"atnlie/domain/users"
	"atnlie/services"
	"atnlie/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	/*
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		fmt.Println("err -> ", err.Error())
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle json error
		fmt.Println("err -> ", err.Error())
		return
	}

	fmt.Println("user: ", user)
	fmt.Println("bytes:", string(bytes))
	fmt.Println("err: ", saveErr)

	*/
	//simplify with this
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: return bad request to caller
		//restError := errors.RestErr {
		//	Message: "invalid json body",
		//	Status: http.StatusBadRequest,
		//	Error: "bad_request",
		//}
		restError := errors.CustomBadRequestError("Invalid json body")
		c.JSON(restError.Status, restError)
		fmt.Println("err -> ", err.Error())
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		c.JSON(saveErr.Status, saveErr)
		fmt.Println("err -> ", saveErr.Message)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.CustomBadRequestError("user_id should a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.CustomBadRequestError("user_id should a number")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.CustomBadRequestError("Invalid json body")
		c.JSON(restError.Status, restError)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}