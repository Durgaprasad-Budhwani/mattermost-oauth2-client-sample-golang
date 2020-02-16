package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func main() {
	r := gin.Default()

	mattermostURL := "http://localhost:8065" // url for mattermost server
	conf := &oauth2.Config{
		ClientID:     "mizxg8mxfibi9cy35xc1z8m9ge", // client id
		ClientSecret: "9rixankt4ff5fqdu74dcbowimy", // client secret
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("%s/oauth/access_token", mattermostURL),
			AuthURL:  fmt.Sprintf("%s/oauth/authorize", mattermostURL),
		},
		RedirectURL: "http://localhost:8080/oauth/callback", // callback url
	}

	r.GET("/login", func(c *gin.Context) {
		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		fmt.Printf("Redirect to url: %v", url)
		c.Redirect(307, url)

	})

	r.GET("/oauth/callback", func(c *gin.Context) {
		//Use the authorization code that is pushed to the redirect
		//URL. Exchange will do the handshake to retrieve the
		//initial access token. The HTTP Client returned by
		//conf.Client will refresh the token as necessary.
		code := c.Query("code")
		tok, err := conf.Exchange(context.Background(), code, oauth2.AccessTypeOffline)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Access token: %v", tok.AccessToken)
		c.JSON(http.StatusOK, gin.H{"Access token": tok.AccessToken})
	})

	r.Run()
}
