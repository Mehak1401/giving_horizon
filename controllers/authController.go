package controllers

import (
	"context"
	"encoding/json"
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

func GoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OAuth exchange failed"})
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
		return
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	var userInfo map[string]interface{}
	json.Unmarshal(data, &userInfo)

	user := models.User{
		GoogleID: userInfo["id"].(string),
		Email:    userInfo["email"].(string),
		Name:     userInfo["name"].(string),
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
