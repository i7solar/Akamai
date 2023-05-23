package main

import (
	_ "fmt"
	"github.com/SolarSystems-Software/HTTP2Client/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	_ "log"
	"time"
)

var (
	deviceList []AkamaiDevice
)

// ErrorApiCheck Checks for service erros.
func ErrorApiCheck(e error, f *fiber.Ctx) bool {
	if e != nil {
		_ = f.JSON(fiber.Error{
			Code:    400,
			Message: e.Error(),
		})
		_ = f.SendStatus(400)
		return true
	} else {
		return false
	}
}

// StartService  Heart of API, starts the service.
func StartService() {
	// Grab all devices prior to starting service.
	deviceList = ReturnAkamaiDevices(false)

	// Service init.
	service := fiber.New(fiber.Config{
		ServerHeader: "SolarSystems Akamai Web Generator",
	})

	// Force all client to use HTTPS if they ever try to use HTTP protocol.
	service.Use(func(ctx *fiber.Ctx) error {
		ctx.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		return ctx.Next()
	})

	// Health Check
	service.Get("/status", func(f *fiber.Ctx) error {
		return f.SendStatus(http.StatusOK)
	})

	v1 := service.Group("/v1")

	// /generate
	generateAkamai := v1.Group("/generate", func(ctx *fiber.Ctx) error {

		// Check target site inputted.
		if !(targetCheck(ctx)) {
			_ = ctx.JSON(fiber.Error{
				Code:    0,
				Message: "Target lookup failed, please visit target header value.",
			})
			ctx.Status(http.StatusBadRequest)
			return nil
		}

		// Check proxy inputted.
		if !checkIfValidProxy(ctx.Get("X-Proxy")) {
			_ = ctx.JSON(fiber.Error{
				Code:    0,
				Message: "Proxy format incorrect, please visit header value.",
			})
			ctx.Status(http.StatusBadRequest)
			return nil
		}

		return ctx.Next()
	})

	// (/generate route) - Timeout given of 30 seconds.
	generateAkamai.Get("/", timeout.New(HandleAkamai, 30*time.Second))

	if err := service.Listen(":80"); err != nil {
		panic(err)
	}
}
