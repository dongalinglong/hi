package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

func obtainTokens() (access, refresh string, err error) {
	conf := &oauth2.Config{
		ClientID:     "4955964",
		ClientSecret: "iavSXd5s8SyRbEWkpKCJ",
		Scopes:       []string{"audio", "video", "offline"},
		RedirectURL:  "https://dll.dyndns.org:8599",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://oauth.vk.com/authorize",
			TokenURL: "https://oauth.vk.com/access_token",
		},
	}
	authUrl := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	open.Run(authUrl)

	fmt.Print("Enter code: ")
	var code string
	_, er := fmt.Scan(&code)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println(code)
	access = ""
	refresh = ""
	err = nil
	return
}

func main() {
	acs, refr, err := obtainTokens()

	fmt.Println(acs, refr, err)
}
