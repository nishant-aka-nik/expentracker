package handler

import (
	"encoding/json"
	"expentracker/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const tokeninfo string = "https://oauth2.googleapis.com/tokeninfo?id_token="

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Dashboard(c *gin.Context) {
	var userData model.TokenUser
	authToken := c.Request.Header.Get("Authorization")

	responseBody, err := getBody(tokeninfo + authToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":             "invalid_token",
			"error_description": err,
		})
	}

	json.Unmarshal(responseBody, &userData)

	c.JSON(http.StatusOK, userData)
	// c.HTML(http.StatusTemporaryRedirect, "dashboard.html", nil)
}

func getBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
