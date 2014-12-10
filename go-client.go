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

	portalapi("authenticate")

}

func portalapi(action string) {
	Transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: Transport}
	urlValues := url.Values{}
	urlValues.Set("login", "ilg2.4.59")
	urlValues.Set("password", "FWSKXFym")
	//urlValues.Set("action", "authenticate")
	urlValues.Set("action", action)
	urlValues.Set("password_digest", "f76bbb94d649d9d5958882501eda86bac57842c3")
	response, err := client.PostForm("https://controller.mobile.lan/portal_api.php", urlValues)
	if err != nil {
		log.Print(err)
	} else {
		defer response.Body.Close()
		ioresp, _ := ioutil.ReadAll(response.Body)
		log.Print(string(ioresp))
	}
}
