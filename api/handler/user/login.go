package user

import (
	"HexMaster/api/response"
	"HexMaster/database"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	res := ctx.Locals("res").(response.Response)
	user := User{}

	err := ctx.BodyParser(&user)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Body: JSON has false format")
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	query := "SELECT * FROM users WHERE email = ?"
	users, amount, err := database.Select[User](query, user.Email)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Database: ")
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}
	if amount != 1 {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Email or Password incorrect")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	baseValues := PBKDF2Hash{
		KeyLen:  len(user.password),
		SaltLen: len(user.Salt),
	}
	if nil != baseValues.Compare([]byte(users[0].Hash), []byte(users[0].Salt), []byte(user.password)) {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Email or Password incorrect")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	token, err := GenerateJWT(users[0])
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil

	}
	users[0].Token = token
	res.Content = users[0]
	res.Send(200)
	return nil
}
