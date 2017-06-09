package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/levigross/grequests"
)

var flashAgent = "ShockwaveFlash/24.0.0.221"

var cookiesFile = "cookies.txt"
var timeout = time.Second * 15 // defaulf global timeout
//delay = (0, 5)  # default mininum delay between requests (random range)

var cookies []*http.Cookie

type Core struct {
	//config
	credits     int
	cookiesFile string
	timeout     time.Duration
	//delay int
	requestTime int
	urls        map[string]string

	//user info
	email        string
	passwd       string
	secretAnswer string

	//db TODO

	//session info
	session *grequests.Session
	header  map[string]string

	sku string
}

func (c *Core) Init(email, passwd, secretAnswer string) {
	c.email = email
	c.passwd = passwd
	c.secretAnswer = secretAnswer

	c.credits = 0
	c.cookiesFile = cookiesFile
	c.timeout = timeout
	//c.delay = ??
	c.requestTime = 0

	c.urls = GetURLs()
	c.header = GetChromeHeader()
}

func (c *Core) Login() error {
	//Init some param
	// secretAnswerHash := util.EAHash(c.secretAnswer)
	// gameSku := "FFA17PS4"
	// platform := "ps4"
	// sku := "FUT17WEB"
	// c.sku = sku
	// clientVersion := 1
	// skuA := "F17"

	//Create new session
	reqOption := &grequests.RequestOptions{
		Headers:        c.header,
		RequestTimeout: c.timeout,
	}
	c.session = grequests.NewSession(reqOption)

	//session load cookies

	//session save header

	//GET login
	// homeResp, err := c.session.Get(c.urls["fut_home"], nil)
	// if err != nil {
	// 	return fmt.Errorf("Get fut_home failed [%s]", err.Error())
	// }
	// var loginURL struct {
	// 	url string
	// }
	// homeResp.JSON(&loginURL)

	// log.Println(homeResp.Header)
	// log.Println(homeResp.StatusCode)
	// homeResp.DownloadToFile("./response")
	//if login != fut_home

	//POST login
	c.session.RequestOptions.Headers["Referer"] = c.urls["fut_home"]
	loginRequestOption := &grequests.RequestOptions{
		Data: map[string]string{
			"email":              c.email,
			"password":           c.passwd,
			"country":            "US", // is it important?
			"phoneNumber":        "",   // TODO: add phone code verification
			"passwordForPhone":   "",
			"gCaptchaResponse":   "",
			"isPhoneNumberLogin": "false", // TODO: add phone login
			"isIncompletePhone":  "",
			"_rememberMe":        "on",
			"rememberMe":         "on",
			"_eventId":           "submit",
		},
	}

	loginResp, err := c.session.Post(c.urls["fut_home"], loginRequestOption)
	if err != nil {
		return fmt.Errorf("Post login failed [%s]", err.Error())
	}

	//Analyse login response
	//respStr := loginResp.String()
	log.Println(loginResp.StatusCode)

	//if success false

	return nil
}
