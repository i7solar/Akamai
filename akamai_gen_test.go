package main

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

/* To test: Input the following command into cmd or terminal: go test -count=1 */
func TestGenerateAkamaiSensor(t *testing.T) {
	var (
		bmak Bmak
		f    *fiber.Ctx // NIL.
	)

	targetSite := "mrporter"
	deviceList := ReturnAkamaiDevices(false)

	_, _, _, _, finalCookie := GenerateCookie(f, "", targetSite, deviceList[bmak.GenerateRandomNumber(0, len(deviceList))], true)

	switch targetSite {
	case "zalando":
		if len(finalCookie) != 533 {
			t.Errorf("[FAILED] Cookie length returned %d :: Want 533", len(finalCookie))
		}
		break
	case "finishline":
		if len(finalCookie) != 533 {
			t.Errorf("[FAILED] Cookie length returned %d :: Want 533", len(finalCookie))
		}
		break
	case "jdsports":
		if len(finalCookie) != 533 {
			t.Errorf("[FAILED] Cookie length returned %d :: Want 533", len(finalCookie))
		}
		break
	case "nike":
		if len(finalCookie) != 525 {
			t.Errorf("[FAILED] Cookie length returned %d :: Want 525", len(finalCookie))
		}
		break
	case "mpv":
		if len(finalCookie) != 464 {
			t.Errorf("[FAILED] Cookie length returned %d :: Want 464", len(finalCookie))
		}
		break
	case "dsg":
		if len(finalCookie) != 553 {
			t.Errorf("[FAILED] Cookie length returned %d :: Want 553", len(finalCookie))
		}
		break

	}
}
