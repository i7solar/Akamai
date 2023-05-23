package main

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"unicode"
)

/* Checks the target website value, if we support that site return true... else return false. */
func targetCheck(c *fiber.Ctx) bool {
	switch c.Get("X-Target") {
	case "finishline":
		return true
	case "jdsports":
		return true
	case "zalando":
		return true
	case "ys":
		return true
	case "nike":
		return true
	case "offspring":
		return true
	case "mpv":
		return true
	case "bestbuy_us":
		return true
	case "asos":
		return true
	case "dsg":
		return true
	case "qvc":
		return true
	case "kickz":
		return true
	case "notebooksbilliger":
		return true
	case "evga":
		return true
	case "converse":
		return true
	case "aldi":
		return true
	case "argos":
		return true
	case "disney":
		return true
	case "gamestop":
		return true
	case "pacsun":
		return true
	case "net-a-porter":
		return true
	case "mrporter":
		return true
	case "prodirectbasketball":
		return true
	case "prodirectsoccer":
		return true
	case "fanatics":
		return true
	case "fansedge":
		return true
	case "sportsmemorabilia":
		return true
	case "target_au":
		return true
	case "footasylum_debug":
		return true
	case "nbastore":
		return true
	case "zozo":
		return true
	case "ebuyer":
		return true
	case "allstate":
		return true
	case "stoneisland":
		return true
	case "bigw":
		return true
	case "sizeer":
		return true
	case "sizeer_ro":
		return true
	case "sizeer_de":
		return true
	case "bestbuy_ca":
		return true
	case "adidas_au":
		return true
	case "zalando_ch":
		return true
	case "luisaviaroma":
		return true
	case "homedepot":
		return true
	case "panera":
		return true
	case "woolworthsrewards":
		return true
	case "kohls":
		return true
	case "macys":
		return true
	case "cabelas":
		return true
	case "scottycameron":
		return true
	case "sephora":
		return true
	case "autozone":
		return true
	case "swatch":
		return true
	case "ae":
		return true
	}
	return false
}

/* Check if the proxy is of a valid length (2) since it needs to be split. */
func checkIfValidProxy(proxyValue string) bool {

	// If the length of the proxy is 0, return false.
	if len(proxyValue) <= 1 {
		return false
	}

	// We'll assume the user is good always.

	// Split the proxy into seperate values.
	proxySplit := strings.Split(proxyValue, ":")

	// (USER) (PW@IP) (PORT)
	if len(proxySplit) == 3 {
		if strings.Contains(proxySplit[1], "@") {
			// Check for port.
			if len(proxySplit[2]) <= 5 { // The largest port is 65535, length 5.
				for _, char := range proxySplit[2] { // Range over the port value and make sure each value is a digit..
					if !unicode.IsDigit(char) {
						return false
					}
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	// (USER) (PORT)
	if len(proxySplit) == 2 {
		// Check for port.
		if len(proxySplit[1]) <= 5 { // The largest port is 65535, length 5.
			for _, char := range proxySplit[1] { // Range over the port value and make sure each value is a digit..
				if !unicode.IsDigit(char) {
					return false
				}
			}
		}
	}

	return true
}
