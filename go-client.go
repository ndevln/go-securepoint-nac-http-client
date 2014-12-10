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
	Transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: Transport}
	urlValues := url.Values{}
	urlValues.Set("login", "ilg2.4.59")
	urlValues.Set("password", "FWSKXFym")
	urlValues.Set("action", "authenticate")
	response, err := client.PostForm("https://controller.mobile.lan/portal_api.php", urlValues)
	if err != nil {
		log.Print(err)
	} else {
		defer response.Body.Close()
		//log.Print(response.Header)
		ioresp, _ := ioutil.ReadAll(response.Body)
		log.Print(string(ioresp))
	}
}
