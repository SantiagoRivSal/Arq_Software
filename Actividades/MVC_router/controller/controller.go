package controller

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"router/dto"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Control(c *gin.Context) {
	c.String(200, "bien!")
}
func Login(c *gin.Context) {
	var userDto dto.User
	err := c.BindJSON(&userDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if compare(userDto) == false {
		fmt.Println("Error")
		log.Error(http.StatusBadRequest)
		return
	}
	var token dto.Token
	token.Token = "4a08da1fecbb6e8b46990538c7b50b2"
	c.JSON(http.StatusAccepted, token)
}

func compare(userDto dto.User) bool {
	userCompare := "{" + userDto.Name + " ; " + userDto.Pass + "}"
	readFile, err := os.Open("login.txt")
	if err != nil {
		fmt.Println(err)
		return false
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() == userCompare {
			return true
		}
	}
	return false
}
