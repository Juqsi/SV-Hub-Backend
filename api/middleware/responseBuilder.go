package middleware

import (
	"HexMaster/api/response"
	"github.com/gofiber/fiber/v2"
)

func ResponseBuilder(ctx *fiber.Ctx) error {
	ctx.Locals("response", response.Response{
		false,
		"",
		[]string{},
		nil,
		ctx,
	})
	return ctx.Next()
}
