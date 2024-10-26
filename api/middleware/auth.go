package middleware

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Authentication(ctx *fiber.Ctx) error {

	//response erstellen
	ctx.Locals("response", response.Response{
		true,
		"",
		[]string{},
		nil,
		ctx,
	})
	response := ctx.Locals("response").(response.Response)
	userToken := ctx.Get("Authorization", "")
	if userToken == "" {
		response.Access = false
		response.Msg = "Melde dich erneut an"
		response.Error = append(response.Error, "Authorization header is missing")
		response.Send(fiber.StatusUnauthorized)
		return nil
	}
	tmp := strings.SplitAfter(userToken, "Bearer ")
	if len(tmp) != 2 {
		response.Access = false
		response.Msg = "Melde dich erneut an"
		response.Error = append(response.Error, "Token has false format")
		response.Send(fiber.StatusUnauthorized)
		return nil
	}
	userToken = tmp[1]

	token, err := user.ValidateJWT(userToken)
	if err != nil {
		response.Access = false
		response.Error = append(response.Error, err.Error())
		response.Send(fiber.StatusUnauthorized)
		return nil
	}

	if token.ID == "" {
		response.Access = false
		response.Error = append(response.Error, "Token ID is missing")
		response.Send(fiber.StatusUnauthorized)
		return nil
	}

	ctx.Locals("token", *token)

	users, count, err := database.Select[user.User]("SELECT * FROM users WHERE id=?;", token.Id)
	if err != nil {
		response.Access = false
		response.Error = append(response.Error, err.Error())
		response.Send(fiber.StatusUnauthorized)
		return nil
	}

	if count != 1 {
		response.Access = false
		response.Error = append(response.Error, "User not found")
		response.Send(fiber.StatusUnauthorized)
	}

	ctx.Locals("user", users[0])

	return ctx.Next()
}
