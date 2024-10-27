package user

import (
	"HexMaster/api/response"
	"HexMaster/database"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	user := User{}

	err := ctx.BodyParser(&user)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Body: JSON has false format")
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	query := "SELECT * FROM users WHERE username = ?"
	users, amount, err := database.Select[User](query, user.Username)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Database: ")
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}
	if amount != 1 {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Username or Password incorrect")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	baseValues := PBKDF2Hash{
		KeyLen:  len(user.Password),
		SaltLen: len(user.Salt),
	}
	hash, _ := hex.DecodeString(users[0].Hash)
	salt, _ := hex.DecodeString(users[0].Salt)
	if nil != baseValues.Compare(hash, salt, []byte(user.Password)) {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Username or Password incorrect")
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
