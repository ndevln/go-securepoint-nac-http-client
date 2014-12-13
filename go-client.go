// go-securepoint-http-client
package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type LoginData struct {
	Login     string
	Password  string
	PwdDigest string
}

func main() {
	// Setup Logfile
	logfile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Print("Failed to open log file", ":", err)
	}
	multiLogger := io.MultiWriter(logfile, os.Stdout)
	log.SetOutput(multiLogger)

	// Load UserData from login.conf (json format)
	loginData := loadUserData("login.conf")

	// Reconnect
	log.Print("Disconnect started...")
	log.Print(portalapi(loginData.Login, loginData.Password, loginData.PwdDigest, "disconnect"))
	log.Print("Connect started...")
	log.Print(portalapi(loginData.Login, loginData.Password, loginData.PwdDigest, "authenticate"))
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

func loadUserData(file string) LoginData {
	// file: filename of config file
	loginFile, fileErr := os.Open(file)
	if fileErr != nil {
		log.Fatal("Failed loading UserData:", fileErr)
	}
	jsondec := json.NewDecoder(loginFile)
	loginData := LoginData{}
	err := jsondec.Decode(&loginData)
	if err != nil {
		log.Fatal("Failed decoding json:", err)
	}
	return loginData
}
