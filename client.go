package main

import (
	"errors"
	"github.com/SolarSystems-Software/HTTP2Client/http"
	"github.com/SolarSystems-Software/HTTP2Client/http/cookiejar"
	"github.com/SolarSystems-Software/HTTP2Client/tlsadapter"
	utls "github.com/SolarSystems-Software/HTTP2Client/tlsproxy/utls"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/url"
	"strings"
)

// GenerateCookie - Hierachy function that handles the main operation: gen a cookie.
func GenerateCookie(f *fiber.Ctx, proxyURL string, targetSite string, device AkamaiDevice, genPixel bool) (string, string, string, string, string) {
	var (
		bmak                    Bmak
		data, cookie            string
		pixelCookie             *http.Cookie
		pixelError, cookieError error
	)

	akamaiProcessor := BuildProcessor(targetSite, device.UserAgent)
	client, clientError := constructClient(proxyURL)

	if ErrorApiCheck(clientError, f) {
		return "", "", "", "", ""
	}

	////////////////////////////////////// Init (0)
	cookie, cookieError = initCookie(client, akamaiProcessor)
	if cookieError != nil {
		log.Println("An unexpected error has occured:", cookieError)
		return "", "", "", "", ""
	}
	var err error
	//////////////////////////// First Sensor (1) Sends the endpoint the sensor data to receive the second cookie to solve challenge

	for i := 0; i < 3; i++ {

		switch i {
		case 0:
			data, err = bmak.getSensor(device, cookie, akamaiProcessor[0]+"/", false, false)
			if ErrorApiCheck(err, f) {
				return "", "", "", "", ""
			}
			cookie, _, err = executePOST("1st", client, akamaiProcessor, data)
			if ErrorApiCheck(err, f) {
				return "", "", "", "", ""
			}
			break
		case 1:
			data, err = bmak.getSensor(device, cookie, akamaiProcessor[0]+"/", true, false)
			if ErrorApiCheck(err, f) {
				return "", "", "", "", ""
			}
			cookie, _, err = executePOST("2nd", client, akamaiProcessor, data)
			if ErrorApiCheck(err, f) {
				return "", "", "", "", ""
			}
		case 2:
			data, err = bmak.getSensor(device, cookie, akamaiProcessor[0]+"/", false, true)
			if ErrorApiCheck(err, f) {
				return "", "", "", "", ""
			}
			cookie, _, err = executePOST("3rd", client, akamaiProcessor, data)
			if ErrorApiCheck(err, f) {
				return "", "", "", "", ""
			}
		}

	}

	if genPixel {
		var pixelEntity PixelEntity
		pixelCookie, pixelError = pixelEntity.GeneratePixel(bmak.device, akamaiProcessor[0], client)
		if pixelError != nil {
			log.Println("Pixel error:", pixelError)
			return bmak.device.UserAgent, data, cookie, "", ""
		} else {
			if len(pixelCookie.Value) != 0 {
				return bmak.device.UserAgent, data, cookie, pixelCookie.Value, ""
			} else {
				return bmak.device.UserAgent, data, cookie, "", ""
			}
		}
	} else {
		return bmak.device.UserAgent, data, cookie, "", ""
	}
}

// constructClient Consturcts the custom http.client.
func constructClient(proxyURL string) (*http.Client, error) {
	var (
		clientCookieJar      *cookiejar.Jar
		clientCookieJarError error
		headers              http.Header
	)

	clientCookieJar, clientCookieJarError = cookiejar.New(nil)
	if clientCookieJarError != nil {
		return nil, clientCookieJarError
	}

	if len(proxyURL) > 1 {
		proxyURL = "http://" + proxyURL
	}

	client, clientError := tlsadapter.NewClientFromSpec(clientCookieJar, utls.HelloChrome_87, utls.ClientHelloSpec{}, proxyURL, headers)

	return &client, clientError
}

// initCookie Generates the first cookie [TEMP] and starts the process of creating sensor data.
func initCookie(client *http.Client, ProcessArray []string) (string, error) {

	var (
		resp        *http.Response
		cookieValue string
		err         error
	)

	cookieValue = "-1"

	request, err := http.NewRequest("GET", ProcessArray[1], nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("accept", "*/*")
	request.Header.Add("accept-encoding", "gzip, deflate, br")
	request.Header.Add("accept-language", "en-US,en;q=0.9")
	request.Header.Add("referer", ProcessArray[0]+"/")
	request.Header.Add("sec-fetch-dest", "script")
	request.Header.Add("sec-fetch-mode", "no-cors")
	request.Header.Add("sec-fetch-site", "same-origin")
	request.Header.Set("user-agent", ProcessArray[2])

	request.HeaderOptions.H2SpecialHeadersKeyOrder = []string{
		":method",
		":authority",
		":scheme",
		":path",
	}

	request.HeaderOptions.HeadersKeyOrder = []string{
		"sec-ch-ua",
		"sec-ch-ua-mobile",
		"user-agent",
		"accept",
		"sec-fetch-site",
		"sec-fetch-mode",
		"sec-fetch-dest",
		"referer",
		"accept-encoding",
		"accept-language",
		"cookie",
	}

	resp, err = client.Do(request)
	if err != nil {
		return "", err
	}

	website, err := url.Parse(ProcessArray[0])
	if err != nil {
		return "", err
	}

	if resp != nil {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "_abck" {
				cookieValue = cookie.Value
				break
			}
		}

		if cookieValue == "-1" {
			return "-1", errors.New("cookie value wasn't changed")
		} else {
			client.Jar.SetCookies(website, resp.Cookies())
		}

	} else {
		return "-1", errors.New("cookie value wasn't changed")
	}

	// Check for the cookieValue and if it is not -1 (meaning we updated cookieValue, then return the value).
	if cookieValue != "-1" {
		return cookieValue, nil
	} else {
		return cookieValue, errors.New("cookie value wasn't changed")
	}

}

// executePOST Updates the cookie to it's final value.
func executePOST(step string, client *http.Client, ProcessArray []string, sensorData string) (string, *http.Cookie, error) {

	var (
		request     *http.Request
		cookieValue string
		httpCookie  *http.Cookie
		err         error
	)

	cookieValue = "-1"

	payloadData := strings.NewReader("{\"sensor_data\":\"" + sensorData + "\"}")
	request, err = http.NewRequest("POST", ProcessArray[1], payloadData)
	if err != nil {
		return "", nil, err
	}
	request.Header.Add("accept", "*/*")
	request.Header.Add("accept-encoding", "gzip, deflate, br")
	request.Header.Add("accept-language", "en-US,en;q=0.9")
	request.Header.Add("content-type", "text/plain;charset=UTF-8")
	request.Header.Add("origin", ProcessArray[0])
	request.Header.Add("referer", ProcessArray[0]+"/")
	request.Header.Add("sec-fetch-dest", "empty")
	request.Header.Add("sec-fetch-mode", "cors")
	request.Header.Add("sec-fetch-site", "same-origin")
	request.Header.Set("user-agent", ProcessArray[2])

	request.HeaderOptions.H2SpecialHeadersKeyOrder = []string{
		":method",
		":authority",
		":scheme",
		":path",
	}

	request.HeaderOptions.HeadersKeyOrder = []string{
		"user-agent",
		"accept",
		"origin",
		"sec-fetch-site",
		"sec-fetch-mode",
		"sec-fetch-dest",
		"referer",
		"accept-encoding",
		"accept-language",
	}

	resp, err := client.Do(request)
	if err != nil {
		return "", nil, err
	}

	website, err := url.Parse(ProcessArray[0])
	if err != nil {
		return "", nil, err
	}

	client.Jar.SetCookies(website, resp.Cookies())

	for _, cookie := range client.Jar.Cookies(website) {
		if cookie.Name == "_abck" {
			cookieValue = cookie.Value
			if step == "3rd" {
				httpCookie = cookie
				break
			}
		}
	}

	// Check for the cookieValue and if it is not -1 (meaning we updated cookieValue, then return the value).
	if cookieValue != "-1" {
		return cookieValue, httpCookie, nil
	} else {
		return "", httpCookie, nil
	}
}
