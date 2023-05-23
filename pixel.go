package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SolarSystems-Software/HTTP2Client/http"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Pixel is an extension of Akamai that is placed on high security websites. The cookie lasts for two hours and is based on metrics.
// Currently Edge is not supported, as if you go on any Akamai page, it does not activate rather than Chrome & FireFox.

////////////////

// PixelEntity Main Pixel Struct that contains all of the json values.
type PixelEntity struct {
	Link      string       // Link to Pixel
	Base64URL string       // base64 value
	AP        bool         `json:"ap"` // Always True
	BT        BtEntity     `json:"bt"`
	Baza      string       // NO JSON HERE. We find this value in the sites html and use it in our ZEntity below.
	Font      string       `json:"fonts"`
	Fh        string       `json:"fh"`
	Timing    TimingEntity `json:"timing"` //(Consists of 5 timing intervals, which increment, and calls all of the other function compute functions known as profiles.
	BP        string       `json:"bp"`     // 4 Values on Chrome, none on FireFox.
	SR        SREntity     `json:"sr"`     // (Consists of screen resolution sizes)
	Dp        DPEntity     `json:"dp"`     // (Document Properties)
	LT        string       `json:"lt"`
	/*
		Ps (Always true,true), Fp (Always false), Sp (Always false), Ieps (Always false), Av (Always false)
		// FC (Always true)
	*/
	CV  string    `json:"cv"`
	BR  string    `json:"br"`
	Z   ZEntity   `json:"z"` // Z (baza) (1) (0)
	ZH  string    `json:"zh"`
	JSV string    `json:"jsv"`
	Nav NavEntity `json:"nav"`
	Crc CRCEntity `json:"crc"` // -- {"window.chrome":"-not-existent"} for FireFox, Values set for Chrome however.
	T   string    `json:"t"`   // T Value
	U   string    `json:"u"`   // G Value
	NAP string    `json:"nap"`
}

type WindowChromeMissing struct {
	WindowChrome string `json:"window.chrome"`
}

// SREntity SR SREntity Values
type SREntity struct {
	Inner      []int `json:"inner"`
	Outer      []int `json:"outer"`
	Screen     []int `json:"screen"`
	Pageoffset []int `json:"pageOffset"`
	Avail      []int `json:"avail"`
	Size       []int `json:"size"`
	Client     []int `json:"client"`
	Colordepth int   `json:"colorDepth"`
	Pixeldepth int   `json:"pixelDepth"`
}

type Profile struct {
	Bp    int `json:"bp"`
	Sr    int `json:"sr"`
	Dp    int `json:"dp"`
	Lt    int `json:"lt"`
	Ps    int `json:"ps"`
	Cv    int `json:"cv"`
	Fp    int `json:"fp"`
	Sp    int `json:"sp"`
	Br    int `json:"br"`
	Ieps  int `json:"ieps"`
	Av    int `json:"av"`
	Z1    int `json:"z1"`
	Jsv   int `json:"jsv"`
	Nav   int `json:"nav"`
	Nap   int `json:"nap"`
	Crc   int `json:"crc"`
	Z2    int `json:"z2"`
	Z3    int `json:"z3"`
	Z4    int `json:"z4"`
	Z5    int `json:"z5"`
	Z6    int `json:"z6"`
	Fonts int `json:"fonts"`
}

// TimingEntity Timing Values
type TimingEntity struct {
	Num1    int `json:"1"`
	Num2    int `json:"2"`
	Num3    int `json:"3"`
	Num4    int `json:"4"`
	Num5    int `json:"5"`
	Num6    int `json:"6"`
	Profile struct {
		Bp    int `json:"bp"`
		Sr    int `json:"sr"`
		Dp    int `json:"dp"`
		Lt    int `json:"lt"`
		Ps    int `json:"ps"`
		Cv    int `json:"cv"`
		Fp    int `json:"fp"`
		Sp    int `json:"sp"`
		Br    int `json:"br"`
		Ieps  int `json:"ieps"`
		Av    int `json:"av"`
		Z1    int `json:"z1"`
		Jsv   int `json:"jsv"`
		Nav   int `json:"nav"`
		Nap   int `json:"nap"`
		Crc   int `json:"crc"`
		Z2    int `json:"z2"`
		Z3    int `json:"z3"`
		Z4    int `json:"z4"`
		Z5    int `json:"z5"`
		Z6    int `json:"z6"`
		Fonts int `json:"fonts"`
	} `json:"profile"`
	Main    int `json:"main"`
	Compute int `json:"compute"`
	Send    int `json:"send"`
}

// DPEntity DPEntity Values
type DPEntity struct {
	Xdomainrequest         int    `json:"XDomainRequest"`
	Createpopup            int    `json:"createPopup"`
	Removeeventlistener    int    `json:"removeEventListener"`
	Globalstorage          int    `json:"globalStorage"`
	Opendatabase           int    `json:"openDatabase"`
	Indexeddb              int    `json:"indexedDB"`
	Attachevent            int    `json:"attachEvent"`
	Activexobject          int    `json:"ActiveXObject"`
	Dispatchevent          int    `json:"dispatchEvent"`
	Addbehavior            int    `json:"addBehavior"`
	Addeventlistener       int    `json:"addEventListener"`
	Detachevent            int    `json:"detachEvent"`
	Fireevent              int    `json:"fireEvent"`
	Mutationobserver       int    `json:"MutationObserver"`
	Htmlmenuitemelement    int    `json:"HTMLMenuItemElement"`
	Int8Array              int    `json:"Int8Array"`
	Postmessage            int    `json:"postMessage"`
	Queryselector          int    `json:"querySelector"`
	Getelementsbyclassname int    `json:"getElementsByClassName"`
	Images                 int    `json:"images"`
	Compatmode             string `json:"compatMode"`
	Documentmode           int    `json:"documentMode"`
	All                    int    `json:"all"`
	Now                    int    `json:"now"`
	Contextmenu            int    `json:"contextMenu"`
}

// CRCEntity CRC values
type CRCEntity struct {
	WindowChrome struct {
		App struct {
			Isinstalled  bool `json:"isInstalled"`
			Installstate struct {
				Disabled     string `json:"DISABLED"`
				Installed    string `json:"INSTALLED"`
				NotInstalled string `json:"NOT_INSTALLED"`
			} `json:"InstallState"`
			Runningstate struct {
				CannotRun  string `json:"CANNOT_RUN"`
				ReadyToRun string `json:"READY_TO_RUN"`
				Running    string `json:"RUNNING"`
			} `json:"RunningState"`
		} `json:"app"`
		Runtime struct {
			Oninstalledreason struct {
				ChromeUpdate       string `json:"CHROME_UPDATE"`
				Install            string `json:"INSTALL"`
				SharedModuleUpdate string `json:"SHARED_MODULE_UPDATE"`
				Update             string `json:"UPDATE"`
			} `json:"OnInstalledReason"`
			Onrestartrequiredreason struct {
				AppUpdate string `json:"APP_UPDATE"`
				OsUpdate  string `json:"OS_UPDATE"`
				Periodic  string `json:"PERIODIC"`
			} `json:"OnRestartRequiredReason"`
			Platformarch struct {
				Arm    string `json:"ARM"`
				Arm64  string `json:"ARM64"`
				Mips   string `json:"MIPS"`
				Mips64 string `json:"MIPS64"`
				X8632  string `json:"X86_32"`
				X8664  string `json:"X86_64"`
			} `json:"PlatformArch"`
			Platformnaclarch struct {
				Arm    string `json:"ARM"`
				Mips   string `json:"MIPS"`
				Mips64 string `json:"MIPS64"`
				X8632  string `json:"X86_32"`
				X8664  string `json:"X86_64"`
			} `json:"PlatformNaclArch"`
			Platformos struct {
				Android string `json:"ANDROID"`
				Cros    string `json:"CROS"`
				Linux   string `json:"LINUX"`
				Mac     string `json:"MAC"`
				Openbsd string `json:"OPENBSD"`
				Win     string `json:"WIN"`
			} `json:"PlatformOs"`
			Requestupdatecheckstatus struct {
				NoUpdate        string `json:"NO_UPDATE"`
				Throttled       string `json:"THROTTLED"`
				UpdateAvailable string `json:"UPDATE_AVAILABLE"`
			} `json:"RequestUpdateCheckStatus"`
		} `json:"runtime"`
	} `json:"window.chrome"`
}

// BtEntity BT Values
type BtEntity struct {
	Charging                bool   `json:"charging"`
	ChargingTime            int    `json:"chargingTime"`
	DischargingTime         string `json:"dischargingTime"`
	Level                   int    `json:"level"`
	OnChargingChange        string `json:"onchargingchange"`
	OnChargingTimeChange    string `json:"onchargingtimechange"`
	OnDischargingTimeChange string `json:"ondischargingtimechange"`
	OnLevelChange           string `json:"onlevelchange"`
}

// NavEntity NAV Values
type NavEntity struct {
	UserAgent           string   `json:"userAgent"`
	AppName             string   `json:"appName"`
	AppCodeName         string   `json:"appCodeName"`
	AppVersion          string   `json:"appVersion"`
	AppMinorVersion     int      `json:"appMinorVersion"`
	Product             string   `json:"product"`
	ProductSub          string   `json:"productSub"`
	Vendor              string   `json:"vendor"`
	VendorSub           string   `json:"vendorSub"`
	BuildID             int      `json:"buildID"`
	Platform            string   `json:"platfom"`
	OSCPU               int      `json:"oscpu"`
	HardwareConcurrency int      `json:"hardwareConcurrency"`
	Language            string   `json:"language"`
	Languages           []string `json:"languages"`
	SystemLanguage      int      `json:"systemLanguage"`
	UserLanguage        int      `json:"userLanguage"`
	DoNotTrack          string   `json:"doNotTrack"`
	MsDoNotTrack        int      `json:"msDoNotTrack"`
	CookieEnabled       bool     `json:"cookieEnabled"`
	Geolocation         int      `json:"geolocation"`
	Vibrate             int      `json:"vibrate"`
	MaxTouchPoints      int      `json:"maxTouchPoints"`
	WebDriver           bool     `json:"webdriver"`
	Plugins             []string `json:"plugins"`
}

// ZEntity Z Values
type ZEntity struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

// Helper function for converting single unicode to single string.
func altCodeAlphabet(c string) string {
	// Decodes numbers from \x20 - \x7A
	switch c {
	case "x20":
		return " "
	case "x21":
		return "!"
	case "x22":
		return "\""
	case "x23":
		return "#"
	case "x24":
		return "$"
	case "x25":
		return "%"
	case "x26":
		return "&"
	case "x27":
		return "'"
	case "x28":
		return "("
	case "x29":
		return ")"
	case "x2a":
		return "*"
	case "x2b":
		return "+"
	case "x2c":
		return ","
	case "x2d":
		return "-"
	case "x2e":
		return "."
	case "x2f":
		return "/"
	case "x30":
		return "0"
	case "x31":
		return "1"
	case "x32":
		return "2"
	case "x33":
		return "3"
	case "x34":
		return "4"
	case "x35":
		return "5"
	case "x36":
		return "6"
	case "x37":
		return "7"
	case "x38":
		return "8"
	case "x39":
		return "9"
	case "x3a":
		return ":"
	case "x3b":
		return ";"
	case "x3c":
		return "<"
	case "x3d":
		return "="
	case "x3e":
		return ">"
	case "x3f":
		return "?"
	case "x40":
		return "@"
	case "x41":
		return "A"
	case "x42":
		return "B"
	case "x43":
		return "C"
	case "x44":
		return "D"
	case "x45":
		return "E"
	case "x46":
		return "F"
	case "x47":
		return "G"
	case "x48":
		return "H"
	case "x49":
		return "I"
	case "x4a":
		return "J"
	case "x4b":
		return "K"
	case "x4c":
		return "L"
	case "x4d":
		return "M"
	case "x4e":
		return "N"
	case "x4f":
		return "O"
	case "x50":
		return "P"
	case "x51":
		return "Q"
	case "x52":
		return "R"
	case "x53":
		return "S"
	case "x54":
		return "T"
	case "x55":
		return "U"
	case "x56":
		return "V"
	case "x57":
		return "W"
	case "x58":
		return "X"
	case "x59":
		return "Y"
	case "x5a":
		return "Z"
	case "x5b":
		return "["
	case "x5c":
		return "\\"
	case "x5d":
		return "]"
	case "x5e":
		return "^"
	case "x5f":
		return "_"
	case "x60":
		return "`"
	case "x61":
		return "a"
	case "x62":
		return "b"
	case "x63":
		return "c"
	case "x64":
		return "d"
	case "x65":
		return "e"
	case "x66":
		return "f"
	case "x67":
		return "g"
	case "x68":
		return "h"
	case "x69":
		return "i"
	case "x6a":
		return "j"
	case "x6b":
		return "k"
	case "x6c":
		return "l"
	case "x6d":
		return "m"
	case "x6e":
		return "n"
	case "x6f":
		return "o"
	case "x70":
		return "p"
	case "x71":
		return "q"
	case "x72":
		return "r"
	case "x73":
		return "s"
	case "x74":
		return "t"
	case "x75":
		return "u"
	case "x76":
		return "v"
	case "x77":
		return "w"
	case "x78":
		return "x"
	case "x79":
		return "y"
	case "x7a":
		return "z"
	case "x7b":
		return "{"
	case "x7c":
		return "|"
	case "x7d":
		return "}"
	case "x7e":
		return "~"
	}
	return ""
}

// Helper function for converting set of unicode to its string equivalent.
func altCodeMedium(char string) string {
	return altCodeAlphabet(char)
}

// Helper function known as the heart of the function for converting unicode to string.
func altCodeReader(altCode string) string {
	arr := strings.Split(altCode, "\\")
	returnStr := ""
	for _, seq := range arr {
		returnStr += altCodeMedium(seq)
	}
	return returnStr
}

// GrabPixelPage Grabs the pixel page from the requested targetURL.
func (Entity *PixelEntity) GrabPixelPage(client *http.Client, targetURL string) (string, error) {

	request, err := http.NewRequest("GET", targetURL, strings.NewReader(""))
	if err != nil {
		return "", err
	}
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Accept-Language", "en-US,en;q=0.9")
	request.Header.Add("Origin", targetURL)
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("sec-fetch-des", "empty")
	request.Header.Add("sec-fetch-mode", "cors")
	request.Header.Add("sec-fetch-site", "same-origin")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	defer func() {
		_ = request.Body.Close()
	}()

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Use a regex expression
	var pixelRegex = regexp.MustCompile(`(?m)<noscript><img src=*"(.+?)"`)
	pixelMatch := pixelRegex.FindAllString(string(body), -1)

	if len(pixelMatch) < 1 {
		return "", fmt.Errorf("pixelMatch should be at least 1 in length, not %d", len(pixelMatch))
	}

	var pixelString string
	if len(pixelMatch) >= 2 { // If our pixel match returns more than two regex matches, check for "akam" in the link. The goal is to return the pixel string containing the site. ex. https://SITE/akam/11/pixel_PIXELRANDOMVALUE
		for i := range pixelMatch {
			if strings.Contains(pixelMatch[i], "akam") { /// Execute the conditonal statement to check.
				pixelArray := strings.Split(pixelMatch[i], "?")
				pixelArray = strings.Split(pixelArray[0], "img src=\"")
				pixelString = pixelArray[1]
			}
		}
	} else {
		pixelArray := strings.Split(pixelMatch[0], "?")
		pixelArray = strings.Split(pixelArray[0], "img src=\"")
		pixelString = pixelArray[1]
	}

	Entity.Baza = Entity.GrabBaza(string(body))

	return pixelString, nil
}

// GrabGValue Grabs the G value from the targetURL.
func (Entity *PixelEntity) GrabGValue(requestBody string) string {
	var pixelEncodedStringArray [424]string

	pixelTrashArray := strings.Split(requestBody, "var _=[")
	pixelTrashArray = strings.Split(pixelTrashArray[1], "];")
	pixelEncodedTrashArray := pixelTrashArray[0]
	pixelTrashArray = strings.Split(pixelEncodedTrashArray, ",")

	for i := range pixelTrashArray {
		splitString := strings.Split(pixelTrashArray[i], "\"")[1]
		pixelEncodedStringArray[i] = splitString
		convertedPixelArrayValue := altCodeReader(pixelEncodedStringArray[i])
		if strings.Contains(convertedPixelArrayValue, "data:image") {
			Entity.Base64URL = convertedPixelArrayValue
		}
	}

	pixelTrashArray = strings.Split(requestBody, "g=_")
	pixelTrashArray = strings.Split(pixelTrashArray[1], ",")
	pixelTrashArray = strings.Split(pixelTrashArray[0], "[")
	pixelTrashArray = strings.Split(pixelTrashArray[1], "]")
	pixelArrayIndex, _ := strconv.Atoi(pixelTrashArray[0])

	return altCodeReader(pixelEncodedStringArray[pixelArrayIndex])
}

// GrabBaza Grabs the bazadebezolkohpepadr value from the targetURL.
func (Entity *PixelEntity) GrabBaza(requestBody string) string {

	var pixelRegex = regexp.MustCompile(`(?m)bazadebezolkohpepadr=*"(.+?)"`)
	BazaValue := pixelRegex.FindAllString(requestBody, -1)
	BazaValue = strings.Split(BazaValue[0], "\"")

	return BazaValue[1]
}

// InitPixelRequest Executes functions for pixel on targetURL.
func (Entity *PixelEntity) InitPixelRequest(targetSite string, client *http.Client) (bool, error) {
	pixelResource, err := Entity.GrabPixelPage(client, targetSite)
	if err != nil {
		return false, err
	}
	Entity.Link = pixelResource
	pixelResourceArray := strings.Split(pixelResource, "_")

	if len(pixelResourceArray) > 1 {
		requestURL := targetSite + "/akam/11/" + pixelResourceArray[1]
		fmt.Println("REQUEST URL -> " + requestURL)
		fmt.Println("WORKING URL -> \"https://mrporter.com/akam/11/" + pixelResourceArray[1])
		request, requestInitError := http.NewRequest("GET", requestURL, nil)
		if requestInitError != nil {
			return false, requestInitError
		}
		resp, requestError := client.Do(request)
		if requestError != nil {
			return false, requestError
		}
		body, bodyError := io.ReadAll(resp.Body)
		if bodyError != nil {
			return false, bodyError
		}
		fmt.Println("STATUS CODE -> " + strconv.Itoa(resp.StatusCode))
		fmt.Println(string(body))
		Entity.U = Entity.GrabGValue(string(body))
		return true, nil
	} else {
		return false, nil
	}

}

// SetBR Sets BR values for the struct value. (Browser)
func (Entity *PixelEntity) SetBR(device AkamaiDevice) {

	if device.browserName == "Chrome" {
		Entity.BR = "Chrome"
	} else {
		Entity.BR = "Firefox"
	}

}

// SetJSV Sets JSV values for the struct value based on Browser.
func (Entity *PixelEntity) SetJSV(device AkamaiDevice) {

	if device.browserName == "Chrome" {
		Entity.JSV = "1.7"
	} else {
		Entity.JSV = "1.5"
	}
}

// GenerateCV Sets CV value.
func (Entity *PixelEntity) GenerateCV() {
	//var byteArray [11870]byte
	//rand.Read(byteArray[:])
	hashValue := sha1.New()
	hashValue.Write([]byte(Entity.Base64URL))

	Entity.CV = hex.EncodeToString(hashValue.Sum(nil))
}

// GenerateTValue Sets T value.
func (Entity *PixelEntity) GenerateTValue() {
	byteArray := []byte(Entity.Baza)
	h := sha1.New()
	h.Write(byteArray)

	Entity.T = hex.EncodeToString(h.Sum(nil))
}

// GenerateFontHash Sets font hash value.
func (Entity *PixelEntity) GenerateFontHash() {
	byteArray := []byte(Entity.Font)
	h := sha1.New()
	h.Write(byteArray)

	Entity.Fh = hex.EncodeToString(h.Sum(nil))
}

// GenerateLT Sets LT value.
func (Entity *PixelEntity) GenerateLT() {
	var bmak Bmak
	var lt string
	lt = strconv.FormatInt(time.Now().UTC().UnixNano()-bmak.GenerateRandomNumber(8000, 11000), 10)
	lt += "-4"
	Entity.LT = lt
}

// charCodeAt similar to JS internal function.
func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}

// removeDuplicatesUnordered Removes duplicates.
func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	var result []string
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

// SetSR Sets SR value.
func (Entity *PixelEntity) SetSR(device AkamaiDevice) {
	var avail []int
	var size []int
	var screen []int
	inner := []int{device.RawDevice.Window.Innerheight, device.RawDevice.Window.Innerwidth}
	outer := []int{device.RawDevice.Window.Outerheight, device.RawDevice.Window.Outerwidth}
	client := []int{device.RawDevice.Window.Innerheight - 17, device.RawDevice.Window.Innerwidth}

	if device.RawDevice.Window.Outerheight > 1920 {
		avail = []int{1920, 1040}
		size = []int{1920, 1080}
	} else {
		avail = []int{2560, 1400}
		size = []int{2560, 1440}
	}

	screen = []int{0, 0}

	Entity.SR = SREntity{
		Inner:      inner,
		Outer:      outer,
		Screen:     screen,
		Pageoffset: []int{0, 0},
		Avail:      avail,
		Size:       size,
		Client:     client,
		Colordepth: 24,
		Pixeldepth: 24,
	}
}

// GenerateTimings Sets timing values for timing value.
func (Entity *PixelEntity) GenerateTimings(device AkamaiDevice) {

	var bmak Bmak
	// Timings, currently we'll do them up to 6. Many cases it can be 3,4,5, or even 6.
	var t1, t2, t3, t4, t5, t6 int
	var timingValues []int  // 1-6 Timings
	var profileValues []int // Profile (functions)
	var finalValues []int   // "main":X,"compute":X,"send":X

	///////////

	t1 = int(bmak.GenerateRandomNumber(10, 21))
	t2 = t1 + int(bmak.GenerateRandomNumber(90, 142))
	t3 = t2 + int(bmak.GenerateRandomNumber(100, 159))
	t4 = t3 + int(bmak.GenerateRandomNumber(90, 122))
	t5 = t4 + int(bmak.GenerateRandomNumber(98, 132))
	t6 = t5 + 107

	timingValues = append(timingValues, t1)
	timingValues = append(timingValues, t2)
	timingValues = append(timingValues, t3)
	timingValues = append(timingValues, t4)
	timingValues = append(timingValues, t5)
	timingValues = append(timingValues, t6)

	///////////////

	profileValues = append(profileValues, 0)                                    // BP should be 0.
	profileValues = append(profileValues, 0)                                    // SR should be 0.
	profileValues = append(profileValues, 0)                                    // DP should be 0.
	profileValues = append(profileValues, 0)                                    // LT should be 0.
	profileValues = append(profileValues, int(bmak.GenerateRandomNumber(1, 3))) // PS should be 0-2.
	profileValues = append(profileValues, int(bmak.GenerateRandomNumber(3, 7))) // CV should be 3-6.
	profileValues = append(profileValues, int(bmak.GenerateRandomNumber(1, 3))) // FP should be 1-2.
	profileValues = append(profileValues, 0)                                    // SP should be 0.
	profileValues = append(profileValues, 0)                                    // BR should be 0.
	profileValues = append(profileValues, 0)                                    // IEPS should be 0.
	profileValues = append(profileValues, 0)                                    // AV should be 0.
	profileValues = append(profileValues, 0)                                    // Z1 should be 2-7.
	profileValues = append(profileValues, int(bmak.GenerateRandomNumber(2, 8))) // JSV should be 0-3.
	profileValues = append(profileValues, int(bmak.GenerateRandomNumber(0, 3))) // NAV should be 0-2
	profileValues = append(profileValues, 0)                                    // NAP should be 0
	profileValues = append(profileValues, int(bmak.GenerateRandomNumber(0, 2))) //CRC should be 0-2.

	// Pick a random Z value to get the computation value of (1). (Z2-Z6)
	switch bmak.GenerateRandomNumber(2, 7) {
	case 2:
		profileValues = append(profileValues, 1)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
	case 3:
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 1)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
	case 4:
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 1)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
	case 5:
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 1)
		profileValues = append(profileValues, 0)
	case 6:
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 0)
		profileValues = append(profileValues, 1)
	}

	profileValues = append(profileValues, int(bmak.GenerateRandomNumber(4, 12))) // Fonts should be 4-11.

	// Main should be 250-295 for FireFox and 404-528 for Chrome.
	if device.browserName == "Chrome" {
		finalValues = append(finalValues, int(bmak.GenerateRandomNumber(404, 528)))
	} else {
		finalValues = append(finalValues, int(bmak.GenerateRandomNumber(250, 295)))
	}

	finalValues = append(finalValues, int(bmak.GenerateRandomNumber(11, 26)))   // Compute time should 11-26.
	finalValues = append(finalValues, int(bmak.GenerateRandomNumber(520, 599))) // Send time should be 520-599.

	Entity.Timing = TimingEntity{
		Num1: timingValues[0],
		Num2: timingValues[1],
		Num3: timingValues[2],
		Num4: timingValues[3],
		Num5: timingValues[4],
		Num6: timingValues[5],
		Profile: Profile{
			Bp:    profileValues[0],
			Sr:    profileValues[1],
			Dp:    profileValues[2],
			Lt:    profileValues[3],
			Ps:    profileValues[4],
			Cv:    profileValues[5],
			Fp:    profileValues[6],
			Sp:    profileValues[7],
			Br:    profileValues[8],
			Ieps:  profileValues[9],
			Av:    profileValues[10],
			Z1:    profileValues[11],
			Jsv:   profileValues[12],
			Nav:   profileValues[13],
			Nap:   profileValues[14],
			Crc:   profileValues[15],
			Z2:    profileValues[16],
			Z3:    profileValues[17],
			Z4:    profileValues[18],
			Z5:    profileValues[19],
			Z6:    profileValues[20],
			Fonts: profileValues[21],
		},
		Main:    finalValues[0],
		Compute: finalValues[1],
		Send:    finalValues[2],
	}

}

// ComputeBP Computes the BP value.
func (Entity *PixelEntity) ComputeBP(e string) string {

	var t int32
	t = 0

	if len(e) == 0 {
		return fmt.Sprint(t)
	}

	for n := 0; n < len(e); n++ {
		t = ((t << 5) - t) + charCodeAt(e, n)
		t &= t
	}

	return fmt.Sprint(t)
}

// GenerateBP Set the BP value.
func (Entity *PixelEntity) GenerateBP(device AkamaiDevice) {

	var bpValues []string

	if device.browserName == "Chrome" {

		for r := 0; r < len(device.Plugins); r++ {

			for o := 0; o < len(device.Plugins[r]); o++ {

				bpValues = append(bpValues, Entity.ComputeBP(device.Plugins[r]))
			}
		}

		bpValues = removeDuplicatesUnordered(bpValues)
		bpValues = append(bpValues, strconv.Itoa(749224105))

		Entity.BP = strings.Join(bpValues, ",")

	} else {
		Entity.BP = strconv.Itoa(0)
	}

}

// GenerateDP Set the DP value.
func (Entity *PixelEntity) GenerateDP() {
	Entity.Dp = DPEntity{
		Xdomainrequest:         0,
		Createpopup:            0,
		Removeeventlistener:    1,
		Globalstorage:          0,
		Opendatabase:           0,
		Indexeddb:              1,
		Attachevent:            0,
		Activexobject:          0,
		Dispatchevent:          1,
		Addbehavior:            0,
		Addeventlistener:       1,
		Detachevent:            0,
		Fireevent:              0,
		Mutationobserver:       1,
		Htmlmenuitemelement:    0,
		Int8Array:              1,
		Postmessage:            1,
		Queryselector:          1,
		Getelementsbyclassname: 1,
		Images:                 1,
		Compatmode:             "CSS1Compat",
		Documentmode:           0,
		All:                    1,
		Now:                    1,
		Contextmenu:            0,
	}
	if Entity.BR == "Chrome" {
		Entity.Dp.Opendatabase = 1
	}
}

// generateJSON Generates the JSON payload for pixel.
func (Entity *PixelEntity) generateJSON(device AkamaiDevice) string {

	var (
		formString string
		js         []byte
	)

	returnParams := url.Values{}

	/* We've got to add all our params in order. */

	// AP
	returnParams.Add("ap", "true")
	formString += returnParams.Encode() + "&"
	returnParams.Del("ap")

	// BT
	if device.browserName == "Chrome" {
		js, _ = json.Marshal(BtEntity{
			Charging:                true,
			ChargingTime:            0,
			DischargingTime:         "Infinity",
			Level:                   1,
			OnChargingChange:        "null",
			OnChargingTimeChange:    "null",
			OnDischargingTimeChange: "null",
			OnLevelChange:           "null",
		})
		returnParams.Add("bt", string(js))
		formString += returnParams.Encode() + "&"
		returnParams.Del("bt")
	} else {
		returnParams.Add("bt", "0")
		formString += returnParams.Encode() + "&"
		returnParams.Del("bt")
	}

	// Fonts
	Entity.Font = device.RawDevice.Fonts
	returnParams.Add("fonts", Entity.Font)
	formString += returnParams.Encode() + "&"
	returnParams.Del("fonts")

	// FH
	Entity.GenerateFontHash()
	returnParams.Add("fh", Entity.Fh)
	formString += returnParams.Encode() + "&"
	returnParams.Del("fh")

	// Timing
	Entity.GenerateTimings(device)
	js, _ = json.Marshal(Entity.Timing)
	returnParams.Add("timing", string(js))
	formString += returnParams.Encode() + "&"
	returnParams.Del("timing")

	// BP
	Entity.GenerateBP(device)
	js, _ = json.Marshal(Entity.BP)
	returnParams.Add("bp", string(js))
	formString += returnParams.Encode() + "&"
	returnParams.Del("bp")

	// SR
	Entity.SetSR(device)
	js, _ = json.Marshal(Entity.SR)
	returnParams.Add("sr", string(js))
	formString += returnParams.Encode() + "&"
	returnParams.Del("sr")

	// DP
	Entity.GenerateDP()
	js, _ = json.Marshal(Entity.Dp)
	returnParams.Add("dp", string(js))
	formString += returnParams.Encode() + "&"
	returnParams.Del("dp")

	//LT
	Entity.GenerateLT()
	returnParams.Add("lt", Entity.LT)
	formString += returnParams.Encode() + "&"
	returnParams.Del("lt")

	// PS
	returnParams.Add("ps", "true,true")
	formString += returnParams.Encode() + "&"
	returnParams.Del("ps")

	// CV
	Entity.GenerateCV()
	returnParams.Add("cv", Entity.CV)
	formString += returnParams.Encode() + "&"
	returnParams.Del("cv")

	// FP
	returnParams.Add("fp", "false")
	formString += returnParams.Encode() + "&"
	returnParams.Del("fp")

	// SP
	returnParams.Add("sp", "false")
	formString += returnParams.Encode() + "&"
	returnParams.Del("sp")

	// BR
	returnParams.Add("br", Entity.BR)
	formString += returnParams.Encode() + "&"
	returnParams.Del("br")

	// IEPS
	returnParams.Add("ieps", "false")
	formString += returnParams.Encode() + "&"
	returnParams.Del("ieps")

	// AV
	returnParams.Add("av", "false")
	formString += returnParams.Encode() + "&"
	returnParams.Del("av")

	// Z
	Entity.Z = ZEntity{
		A: Entity.Baza,
		B: strconv.Itoa(1),
		C: strconv.Itoa(0),
	}
	js, _ = json.Marshal(Entity.Z)
	returnParams.Add("z", string(js))
	formString += returnParams.Encode() + "&"
	returnParams.Del("z")

	// ZH
	returnParams.Add("zh", Entity.ZH)
	formString += returnParams.Encode() + "&"
	returnParams.Del("zh")

	// JSV
	returnParams.Add("jsv", Entity.JSV)
	formString += returnParams.Encode() + "&"
	returnParams.Del("jsv")

	// NAV
	var appminorversion int
	appmv := device.RawDevice.Navigator.Appminorversion
	if appmv {
		appminorversion = 1
	} else {
		appminorversion = 0
	}

	js, _ = json.Marshal(NavEntity{
		UserAgent:           device.RawDevice.Navigator.Useragent,
		AppName:             device.RawDevice.Navigator.Appname,
		AppCodeName:         device.RawDevice.Navigator.Appcodename,
		AppVersion:          device.RawDevice.Navigator.Appversion,
		AppMinorVersion:     appminorversion,
		Product:             device.RawDevice.Navigator.Product,
		ProductSub:          device.RawDevice.Navigator.Productsub,
		Vendor:              device.RawDevice.Navigator.Vendor,
		VendorSub:           "",
		BuildID:             0,
		Platform:            device.RawDevice.Navigator.Platform,
		OSCPU:               0,
		HardwareConcurrency: 24,
		Language:            device.RawDevice.Navigator.Language,
		Languages:           device.RawDevice.Navigator.Languages,
		SystemLanguage:      0,
		UserLanguage:        0,
		DoNotTrack:          "null",
		MsDoNotTrack:        0,
		CookieEnabled:       true,
		Geolocation:         1,
		Vibrate:             1,
		MaxTouchPoints:      0,
		WebDriver:           false,
		Plugins:             device.RawDevice.Navigator.Plugins,
	})
	returnParams.Add("nav", string(js))
	formString += returnParams.Encode() + "&"
	returnParams.Del("nav")

	// CRC
	if device.browserName == "Chrome" {
		js, _ = json.Marshal(CRCEntity{
			WindowChrome: device.RawDevice.Window.Chrome,
		})
	} else {
		js, _ = json.Marshal(WindowChromeMissing{WindowChrome: "-not-existent"})
	}

	returnParams.Add("crc", string(js))
	formString += returnParams.Encode() + "&"
	returnParams.Del("crc")

	// T
	Entity.GenerateTValue()
	returnParams.Add("t", Entity.T)
	formString += returnParams.Encode() + "&"
	returnParams.Del("t")

	// U
	returnParams.Add("u", Entity.U)
	formString += returnParams.Encode() + "&"
	returnParams.Del("u")

	// NAP
	returnParams.Add("nap", Entity.NAP)
	formString += returnParams.Encode() + "&"
	returnParams.Del("nap")

	// FC
	returnParams.Add("fc", "true")
	formString += returnParams.Encode()
	returnParams.Del("fc")

	return formString
}

func (Entity *PixelEntity) printPixelData() {

	fmt.Println("bt: ", Entity.BT)
	fmt.Println("fonts: ", Entity.Font)
	fmt.Println("fh:", Entity.Fh)
	fmt.Println("timing: ", Entity.Timing)
	fmt.Println("bp: ", Entity.BP)
	fmt.Println("sr: ", Entity.SR)
	fmt.Println("dp: ", Entity.Dp)
	fmt.Println("lt: ", Entity.LT)
	fmt.Println("ps: ", "true,true")
	fmt.Println("cv: ", Entity.CV)
	fmt.Println("fp: ", "false")
	fmt.Println("sp: ", "false")
	fmt.Println("br: ", Entity.BR)
	fmt.Println("ieps: ", "false")
	fmt.Println("av: ", "false")
	fmt.Println("z: ", Entity.Z)
	fmt.Println("zh: ", Entity.ZH)
	fmt.Println("jsv: ", Entity.JSV)
	fmt.Println("nav: ", Entity.Nav)
	fmt.Println("crc: ", Entity.Crc)
	fmt.Println("t: ", Entity.T)
	fmt.Println("u: ", Entity.U)
	fmt.Println("nap: ", Entity.NAP)
	fmt.Println("fc: ", "true")

}

// GeneratePixel The heart of the pixel generation, builds the JSON form data to generate a pixel cookie.
func (Entity *PixelEntity) GeneratePixel(device AkamaiDevice, targetURL string, client *http.Client) (*http.Cookie, error) {
	fmt.Println("TARGET URL -> " + targetURL)
	success, err := Entity.InitPixelRequest(targetURL, client)
	if err != nil {
		return nil, err
	}

	if success {
		Entity.SetJSV(device)
		Entity.SetBR(device)
		Entity.NAP = device.np

		var (
			request     *http.Request
			cookieValue string
			cookieRaw   *http.Cookie
		)

		cookieValue = "-1"

		pixelJSON := Entity.generateJSON(device)
		payloadData := strings.NewReader(pixelJSON)
		request, err = http.NewRequest("POST", Entity.Link, payloadData)
		if err != nil {
			return nil, err
		}
		request.Header.Add("Accept", "*/*")
		request.Header.Add("Accept-Language", "en-US,en;q=0.9")
		request.Header.Add("Content-Type", "text/plain;charset=UTF-8")
		request.Header.Add("Origin", targetURL)
		request.Header.Add("Referer", targetURL+"/")
		request.Header.Add("sec-fetch-des", "empty")
		request.Header.Add("sec-fetch-mode", "cors")
		request.Header.Add("sec-fetch-site", "same-origin")
		request.Header.Set("User-Agent", device.UserAgent)

		resp, err := client.Do(request)
		if err != nil {
			return nil, err
		}

		website, err := url.Parse(targetURL)
		if err != nil {
			return nil, err
		}

		if resp != nil {
			for _, cookie := range resp.Cookies() {
				if cookie.Name == "ak_bmsc" {
					cookieRaw = cookie
					cookieValue = cookie.Value
					break
				}
			}

			if cookieValue == "-1" {
				return nil, errors.New("error generating pixel")
			} else {
				client.Jar.SetCookies(website, resp.Cookies())
			}

		} else {
			return nil, errors.New("error generating pixel")
		}

		client.Jar.SetCookies(website, resp.Cookies())

		return cookieRaw, nil
	} else {
		return nil, nil
	}
}
