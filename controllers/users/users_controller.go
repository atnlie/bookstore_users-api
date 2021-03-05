package users

import (
	"atnlie/domain/users"
	"atnlie/services"
	"atnlie/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
		restError := errors.RestErr {
			Message: "invalid json body",
			Status: http.StatusBadRequest,
			Error: "bad_request",
		}
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
	c.String(http.StatusNotImplemented, "pake aku")
}
