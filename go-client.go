// go-client
package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	login := "ilg2.4.59"
	password := "FWSKXFym"
	passworddigest := "f76bbb94d649d9d5958882501eda86bac57842c3"

	log.Print("Disconnect started...")
	log.Print(portalapi(login, password, passworddigest, "disconnect"))
	log.Print("Connect started...")
	log.Print(portalapi(login, password, passworddigest, "authenticate"))
}

func portalapi(login string, password string, pwddigest string, action string) string {
	// login = username
	// password = password
	// pwdgigest = return value of connect, needed for disconnect (leave empty for connect)
	// action = connect, disconnect or refresh

	Transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: Transport}

	urlValues := url.Values{
		"login":           {login},
		"password":        {password},
		"password_digest": {pwddigest},
		"action":          {action},
	}

	response, err := client.PostForm("https://controller.mobile.lan/portal_api.php", urlValues)
	if err != nil {
		log.Print(err)
		return ""
	} else {
		defer response.Body.Close()
		ioresp, _ := ioutil.ReadAll(response.Body)
		return string(ioresp)
	}
}
