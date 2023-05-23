package main

import (
	"encoding/json"
	"github.com/SolarSystems-Software/HTTP2Client/http"
	"github.com/gofiber/fiber/v2"
)

func HandleAkamai(ctx *fiber.Ctx) error {
	var bmak Bmak

	// Switch case evaulates the pixel header value. (TRUE, FALSE)
	switch ctx.Get("X-Include-Pixel") {
	case "true":
		userAgent, _, finalCookie, pixelCookie, _ := GenerateCookie(ctx, ctx.Get("X-Proxy"), ctx.Get("X-Target"), deviceList[bmak.GenerateRandomNumber(0, len(deviceList))], true)

		if len(finalCookie) > 0 && len(pixelCookie) > 0 {

			ctx.Set("user-agent", userAgent)
			ctx.Set("_abck", finalCookie)
			ctx.Set("ak_bmsc", pixelCookie)
			jsonData, jsonDataError := json.Marshal("_abck & ak_bmsc generation successful")

			ErrorApiCheck(jsonDataError, ctx)
			jsonSendError := ctx.Send(jsonData)
			ErrorApiCheck(jsonSendError, ctx)

			return ctx.SendStatus(http.StatusOK)
		} else if len(finalCookie) > 0 && len(pixelCookie) == 0 {

			ctx.Set("user-agent", userAgent)
			ctx.Set("_abck", finalCookie)
			jsonData, jsonDataError := json.Marshal("_abck generation successful, site does not have pixel active (ak_bmsc)")
			ErrorApiCheck(jsonDataError, ctx)
			jsonSendError := ctx.Send(jsonData)
			ErrorApiCheck(jsonSendError, ctx)

			return ctx.SendStatus(http.StatusOK)
		} else {
			jsonData, jsonDataError := json.Marshal("Error: _abck and ak_bmsc were not generated")
			ErrorApiCheck(jsonDataError, ctx)
			jsonSendError := ctx.Send(jsonData)
			ErrorApiCheck(jsonSendError, ctx)

			return ctx.SendStatus(http.StatusBadRequest)
		}
	case "false":
		userAgent, _, finalCookie, _, _ := GenerateCookie(ctx, ctx.Get("X-Proxy"), ctx.Get("X-Target"), deviceList[bmak.GenerateRandomNumber(0, len(deviceList))], false)

		if len(finalCookie) > 0 {

			ctx.Set("user-agent", userAgent)
			ctx.Set("_abck", finalCookie)
			jsonData, jsonDataError := json.Marshal("_abck successfully generated")
			ErrorApiCheck(jsonDataError, ctx)
			jsonSendError := ctx.Send(jsonData)
			ErrorApiCheck(jsonSendError, ctx)

			return ctx.SendStatus(http.StatusOK)
		} else {

			jsonData, jsonDataError := json.Marshal("Error: _abck generation unsuccessful")
			ErrorApiCheck(jsonDataError, ctx)
			jsonSendError := ctx.Send(jsonData)
			ErrorApiCheck(jsonSendError, ctx)

			return ctx.SendStatus(http.StatusBadRequest)
		}
	}

	return nil
}
