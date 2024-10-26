package user

import (
	"HexMaster/api/response"
	"HexMaster/database"
	"github.com/gofiber/fiber/v2"
)

func Registration(ctx *fiber.Ctx) error {
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

	baseValues := PBKDF2Hash{
		KeyLen:  len(user.password),
		SaltLen: len(user.Salt),
	}

	salt, err := randomSecret(uint32(8))
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Salt generating error")
		res.Send(fiber.StatusInternalServerError)
		return nil
	}
	hashSalt, err := baseValues.GenerateHash([]byte(user.password), salt)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Hash generating error")
		res.Send(fiber.StatusInternalServerError)
		return nil
	}

	user.Salt = string(hashSalt.Salt)
	user.Hash = string(hashSalt.Hash)

	//Secure because of the own escape library, just easier to read
	id, err := database.Insert("INSERT INTO users (forename, lastname, telenum,email,username,hash,salt)VALUES ({forename},{lastname}, {telenum},{email},{username},{hash},{salt});", "id", &user)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Body: JSON has false format")
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	user.Id = id

	token, err := GenerateJWT(user)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil

	}
	user.Token = token
	res.Content = user
	res.Send(200)
	return nil
}
