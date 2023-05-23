package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type DeviceEntity struct {
	ID        string `json:"_id"`
	Navigator struct {
		Appcodename                   string   `json:"appCodeName"`
		Appname                       string   `json:"appName"`
		Appversion                    string   `json:"appVersion"`
		Language                      string   `json:"language"`
		Languages                     []string `json:"languages"`
		Useragent                     string   `json:"userAgent"`
		Vendor                        string   `json:"vendor"`
		Product                       string   `json:"product"`
		Productsub                    string   `json:"productSub"`
		Platform                      string   `json:"platform"`
		Vibrate                       bool     `json:"vibrate"`
		Getbattery                    bool     `json:"getBattery"`
		Credentials                   bool     `json:"credentials"`
		Appminorversion               bool     `json:"appMinorVersion"`
		Bluetooth                     bool     `json:"bluetooth"`
		Storage                       bool     `json:"storage"`
		Getgamepads                   bool     `json:"getGamepads"`
		Getstorageupdates             bool     `json:"getStorageUpdates"`
		Hardwareconcurrency           bool     `json:"hardwareConcurrency"`
		Mediadevices                  bool     `json:"mediaDevices"`
		Mozalarms                     bool     `json:"mozAlarms"`
		Mozconnection                 bool     `json:"mozConnection"`
		Mozislocallyavailable         bool     `json:"mozIsLocallyAvailable"`
		Mozphonenumberservice         bool     `json:"mozPhoneNumberService"`
		Msmanipulationviewsenabled    bool     `json:"msManipulationViewsEnabled"`
		Permissions                   bool     `json:"permissions"`
		Registerprotocolhandler       bool     `json:"registerProtocolHandler"`
		Requestmediakeysystemaccess   bool     `json:"requestMediaKeySystemAccess"`
		Requestwakelock               bool     `json:"requestWakeLock"`
		Sendbeacon                    bool     `json:"sendBeacon"`
		Serviceworker                 bool     `json:"serviceWorker"`
		Storewebwidetrackingexception bool     `json:"storeWebWideTrackingException"`
		Webkitgetgamepads             bool     `json:"webkitGetGamepads"`
		Webkittemporarystorage        bool     `json:"webkitTemporaryStorage"`
		Cookieenabled                 bool     `json:"cookieEnabled"`
		Javaenabled                   bool     `json:"javaEnabled"`
		Donottrack                    int      `json:"doNotTrack"`
		Plugins                       []string `json:"plugins"`
	} `json:"navigator"`
	Window struct {
		Innerheight            int     `json:"innerHeight"`
		Innerwidth             int     `json:"innerWidth"`
		Outerwidth             int     `json:"outerWidth"`
		Outerheight            int     `json:"outerHeight"`
		Devicepixelratio       float64 `json:"devicePixelRatio"`
		Addeventlistener       bool    `json:"addEventListener"`
		Xmlhttprequest         bool    `json:"XMLHttpRequest"`
		Xdomainrequest         bool    `json:"XDomainRequest"`
		Deviceorientationevent bool    `json:"DeviceOrientationEvent"`
		Devicemotionevent      bool    `json:"DeviceMotionEvent"`
		Touchevent             bool    `json:"TouchEvent"`
		Chrome                 struct {
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
		} `json:"chrome"`
		PrototypeBind   bool    `json:"prototype_bind"`
		Pointerevent    bool    `json:"PointerEvent"`
		Sessionstorage  bool    `json:"sessionStorage"`
		Localstorage    bool    `json:"localStorage"`
		Indexeddb       bool    `json:"indexedDB"`
		Filereader      bool    `json:"FileReader"`
		Htmlelement     bool    `json:"HTMLElement"`
		Webrtc          bool    `json:"webRTC"`
		Mozinnerscreeny float64 `json:"mozInnerScreenY"`
	} `json:"window"`
	Document struct {
		Documentmode string `json:"documentMode"`
		Webdriver    bool   `json:"webdriver"`
		Driver       bool   `json:"driver"`
		Selenium     bool   `json:"selenium"`
		Hidden       bool   `json:"hidden"`
		Webkithidden bool   `json:"webkitHidden"`
	} `json:"document"`
	Other struct {
		CcOn             bool `json:"CC_ON"`
		Installtrigger   bool `json:"InstallTrigger"`
		PrototypeForeach bool `json:"prototype_forEach"`
		Imul             bool `json:"imul"`
		Parseint         bool `json:"parseInt"`
		Hypot            bool `json:"hypot"`
		Value1           bool `json:"value1"`
		Xpathresult      bool `json:"XPathResult"`
	} `json:"other"`
	Performance struct {
		Timeorigin float64 `json:"timeOrigin"`
		Timing     struct {
			Connectstart               int64 `json:"connectStart"`
			Navigationstart            int64 `json:"navigationStart"`
			Loadeventend               int64 `json:"loadEventEnd"`
			Domloading                 int64 `json:"domLoading"`
			Secureconnectionstart      int64 `json:"secureConnectionStart"`
			Fetchstart                 int64 `json:"fetchStart"`
			Domcontentloadedeventstart int64 `json:"domContentLoadedEventStart"`
			Responsestart              int64 `json:"responseStart"`
			Responseend                int64 `json:"responseEnd"`
			Dominteractive             int64 `json:"domInteractive"`
			Domainlookupend            int64 `json:"domainLookupEnd"`
			Redirectstart              int   `json:"redirectStart"`
			Requeststart               int64 `json:"requestStart"`
			Unloadeventend             int   `json:"unloadEventEnd"`
			Unloadeventstart           int   `json:"unloadEventStart"`
			Domcomplete                int64 `json:"domComplete"`
			Domainlookupstart          int64 `json:"domainLookupStart"`
			Loadeventstart             int64 `json:"loadEventStart"`
			Domcontentloadedeventend   int64 `json:"domContentLoadedEventEnd"`
			Redirectend                int   `json:"redirectEnd"`
			Connectend                 int64 `json:"connectEnd"`
		} `json:"timing"`
		Navigation struct {
			Type          int `json:"type"`
			Redirectcount int `json:"redirectCount"`
		} `json:"navigation"`
	} `json:"performance"`
	Rcfp []struct {
		Value string `json:"value"`
		Rval  int    `json:"rVal"`
	} `json:"rCFP"`
	Canvas struct {
		Value1 string `json:"value1"`
		Value2 string `json:"value2"`
	} `json:"canvas"`
	FontsOptm string  `json:"fonts_optm"`
	Fonts     string  `json:"fonts"`
	SSH       string  `json:"ssh"`
	Mr        string  `json:"mr"`
	Brave     string  `json:"brave"`
	Wv        string  `json:"wv"`
	Wr        string  `json:"wr"`
	Weh       string  `json:"weh"`
	Wl        int     `json:"wl"`
	Fmh       string  `json:"fmh"`
	Fmz       float64 `json:"fmz"`
	Fpvalstr  string  `json:"fpValstr"`
	Np        string  `json:"np"`
	Screen    struct {
		Availheight int `json:"availHeight"`
		Availleft   int `json:"availLeft"`
		Availtop    int `json:"availTop"`
		Availwidth  int `json:"availWidth"`
		Colordepth  int `json:"colorDepth"`
		Height      int `json:"height"`
		Pixeldepth  int `json:"pixelDepth"`
		Width       int `json:"width"`
		Orientation struct {
		} `json:"orientation"`
	} `json:"screen"`
	Hash string `json:"hash"`
}

type AkamaiDevice struct {
	osSystem          string
	browserName       string
	chromiumSupport   bool
	bluetoothSupport  bool
	javaSupport       bool
	doNotTrackSupport bool
	productName       string
	productSub        string
	UserAgent         string
	versionNumber     string
	canvasArray       [2]string
	rValValue         string
	rCFPValue         string
	fpValString       string
	wv                string
	wr                string
	np                string
	weh               string
	wl                string
	fmh               string
	mrValue           string
	screenResolution  string
	Plugins           []string
	RawDevice         DeviceEntity
}

func parseUser(user DeviceEntity) AkamaiDevice {
	var returnPlatform, returnBrowser string
	var chromiumSup, bluetoothSup, javaSup bool

	// Linux x86_64, Win32, MacIntel returned from platform. Parse and set var.
	platform := user.Navigator.Platform
	switch platform {
	case "Linux x86_64":
		returnPlatform = "Linux"
	case "Win32":
		returnPlatform = "Mac"
	case "MacIntel":
		returnPlatform = "Windows"
	}

	// Parse UA for our browser. thx jorglen
	browser := user.Navigator.Useragent
	if strings.Contains(browser, "rv:") {
		returnBrowser = "Firefox"
	} else if strings.Contains(browser, "Edg/") {
		returnBrowser = "Edge"
	} else {
		returnBrowser = "Chrome"
	}

	// Chromium Support
	if returnBrowser == "Chrome" || returnBrowser == "Edge" {
		chromiumSup = true
	} else {
		chromiumSup = false
	}

	// Bluetooth Support
	if user.Navigator.Bluetooth == true {
		bluetoothSup = true
	} else if user.Navigator.Bluetooth == false {
		bluetoothSup = false
	}

	// Java Support
	if user.Navigator.Javaenabled == true {
		javaSup = true
	} else if user.Navigator.Javaenabled == false {
		javaSup = false
	}

	// screenRes :: availWidth, availHeight, width, height, innerWidth, innerHeight, outerWidth
	availWidth := strconv.Itoa(user.Screen.Availwidth)
	availHeight := strconv.Itoa(user.Screen.Availheight)
	width := strconv.Itoa(user.Screen.Width)
	height := strconv.Itoa(user.Screen.Height)
	innerWidth := strconv.Itoa(user.Window.Innerwidth)
	innerHeight := strconv.Itoa(user.Window.Innerheight)
	outerWidth := strconv.Itoa(user.Window.Outerwidth)
	screenRes := availWidth + "," + availHeight + "," + width + "," + height + "," + innerWidth + "," + innerHeight + "," + outerWidth

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(500) < 250 {
		user.Wl = 33
	} else {
		user.Wl = 31
	}

	return AkamaiDevice{
		osSystem:          returnPlatform,
		browserName:       returnBrowser,
		chromiumSupport:   chromiumSup,
		bluetoothSupport:  bluetoothSup,
		javaSupport:       javaSup,
		doNotTrackSupport: false,
		productName:       user.Navigator.Product,
		productSub:        user.Navigator.Productsub,
		UserAgent:         user.Navigator.Useragent,
		versionNumber:     user.Navigator.Appversion,
		canvasArray:       [2]string{user.Canvas.Value1, user.Canvas.Value2},
		rValValue:         strconv.Itoa(rand.Intn(1000)),
		rCFPValue:         strconv.Itoa(int(rand.Int31())),
		fpValString:       user.Fpvalstr,
		wv:                user.Wv,
		wr:                user.Wr,
		np:                genNP(browser), //user.Np was used prior, currently we have a function that does it for us!
		weh:               user.Weh,
		wl:                strconv.Itoa(user.Wl),
		fmh:               user.Fmh,
		mrValue:           genMR(),
		screenResolution:  screenRes,
		Plugins:           user.Navigator.Plugins,
		RawDevice:         user,
	}

}

func genNP(browserType string) string {
	var (
		napValues [20]string
		napString string
	)
	rand.Seed(time.Now().UnixNano())

	if browserType == "Chrome" || browserType == "Edge" {
		// geolocation
		napValues[0] = strconv.Itoa(1)
		//notifcations
		napValues[1] = strconv.Itoa(0)
		// push
		napValues[2] = strconv.Itoa(3)
		// midi
		napValues[3] = strconv.Itoa(2)
		// camera
		napValues[4] = strconv.Itoa(1)
		// microphone
		napValues[5] = strconv.Itoa(rand.Intn(2))
		// speaker
		napValues[6] = strconv.Itoa(4)
		// deviceInfo
		napValues[7] = strconv.Itoa(4)
		// backgroundSync
		napValues[8] = strconv.Itoa(2)
		// bluetooth
		napValues[9] = strconv.Itoa(4)
		// persisentStorage
		napValues[10] = strconv.Itoa(1)
		// ambientLightSensor
		napValues[11] = strconv.Itoa(3)
		// accelerometer
		napValues[12] = strconv.Itoa(rand.Intn(2))
		// gyroscope
		napValues[13] = napValues[12]
		// megnetometer
		napValues[14] = napValues[12]
		// clipboard
		napValues[15] = strconv.Itoa(4)
		//accessibilityEvents
		napValues[16] = strconv.Itoa(3)
		// clipboard_read
		napValues[17] = strconv.Itoa(rand.Intn(2))
		// clipboard_write
		napValues[18] = strconv.Itoa(2)
		// paymentHandler
		napValues[19] = strconv.Itoa(2)
	} else {
		/* Firefox */
		// geolocation
		napValues[0] = strconv.Itoa(1)
		//notifcations
		napValues[1] = strconv.Itoa(rand.Intn(2))
		// push
		napValues[2] = strconv.Itoa(rand.Intn(2))
		// midi
		napValues[3] = strconv.Itoa(3)
		// camera
		napValues[4] = strconv.Itoa(3)
		// microphone
		napValues[5] = strconv.Itoa(3)
		// speaker
		napValues[6] = strconv.Itoa(3)
		// deviceInfo
		napValues[7] = strconv.Itoa(3)
		// backgroundSync
		napValues[8] = strconv.Itoa(3)
		// bluetooth
		napValues[9] = strconv.Itoa(rand.Intn(1))
		// persisentStorage
		napValues[10] = strconv.Itoa(3)
		// ambientLightSensor
		napValues[11] = strconv.Itoa(3)
		// accelerometer
		napValues[12] = strconv.Itoa(3)
		// gyroscope
		napValues[13] = napValues[12]
		// megnetometer
		napValues[14] = napValues[12]
		// clipboard
		napValues[15] = strconv.Itoa(3)
		//accessibilityEvents
		napValues[16] = strconv.Itoa(3)
		// clipboard_read
		napValues[17] = strconv.Itoa(3)
		// clipboard_write
		napValues[18] = strconv.Itoa(3)
		// paymentHandler
		napValues[19] = strconv.Itoa(3)
	}

	for i := 0; i < len(napValues); i++ {
		napString += napValues[i]
	}

	return napString
}

func genMR() string {
	rand.Seed(time.Now().UnixNano())
	m1 := strconv.Itoa(rand.Intn(40-35)+35) + ","
	m2 := strconv.Itoa(rand.Intn(42-37)+37) + ","
	m3 := strconv.Itoa(rand.Intn(43-36)+36) + ","
	m4 := strconv.Itoa(rand.Intn(44-38)+38) + ","
	m5 := strconv.Itoa(rand.Intn(53-47)+47) + ","
	m6 := strconv.Itoa(rand.Intn(54-48)+48) + ","
	m7 := strconv.Itoa(rand.Intn(45-41)+41) + ","
	m8 := strconv.Itoa(rand.Intn(41-36)+36) + ","
	m9 := strconv.Itoa(rand.Intn(53-47)+47) + ","
	m10 := strconv.Itoa(rand.Intn(8-4)+4) + ","
	m11 := strconv.Itoa(rand.Intn(6-4)+4) + ","
	m12 := strconv.Itoa(rand.Intn(8-5)+5) + ","
	m13 := strconv.Itoa(rand.Intn(7-5)+5) + ","
	m14 := strconv.Itoa(rand.Intn(14-9)+9) + ","
	m15 := strconv.Itoa(rand.Intn(520-440)+440) + ","

	return m1 + m2 + m3 + m4 + m5 + m6 + m7 + m8 + m9 + m10 + m11 + m12 + m13 + m14 + m15
}

func ReturnAkamaiDevices(print bool) []AkamaiDevice {

	users := make([]DeviceEntity, 3000)
	parsedUsers := make([]AkamaiDevice, 3000)
	filename := "moddedsensordata.json"

	// Open our jsonFile, if we os.Open returns an error then handle it.
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = jsonFile.Close()
	}()

	byteArray, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(byteArray, &users); err != nil {
		panic(err)
	}

	for i, user := range users {
		parsedUsers[i] = parseUser(user)
		if print {
			parsedUsers[i].printDevice()
		}
	}

	return parsedUsers
}

func (device *AkamaiDevice) printDevice() {

	if len(device.osSystem) > 0 {
		fmt.Println("OS System: " + device.osSystem)
		fmt.Println("Browser: " + device.browserName)
		fmt.Println("Chromium Support: ", device.chromiumSupport)
		fmt.Println("Bluetooth Support: ", device.bluetoothSupport)
		fmt.Println("Java Support: ", device.javaSupport)
		fmt.Println("Do Not Track Enabled: ", device.doNotTrackSupport)
		fmt.Println("Product: " + device.productName)
		fmt.Println("Product Sub: " + device.productSub)
		fmt.Println("User-Agent: " + device.UserAgent)
		fmt.Println("Version Number: " + device.versionNumber)
		fmt.Println("Canvas Values: " + "[1] : " + device.canvasArray[0] + " || [2] : " + device.canvasArray[1])
		fmt.Println("rVAL: " + device.rValValue)
		fmt.Println("RCFP: " + device.rCFPValue)
		fmt.Println("fpValstr: " + device.fpValString)
		fmt.Println("NP : " + device.np)
		fmt.Println("WV : " + device.wv)
		fmt.Println("WEH : " + device.weh)
		fmt.Println("WL : " + device.wl)
		fmt.Println("FMH : " + device.fmh)
		fmt.Println("Screen Resolution: " + device.screenResolution)
		fmt.Println("Plugins: ")
		for i := 0; i < len(device.Plugins); i++ {
			fmt.Println(device.Plugins[i])
		}
	}

}
