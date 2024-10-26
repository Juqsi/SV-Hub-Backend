package user

import (
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	/*
		response := ctx.Locals("response").(response.R)
		user := User{}

		token := ctx.Locals("token").(Claims)
		/*firebaseID := token.UID
		if firebaseID == "" {
			response.Msg = MSG_FIREBASE_TOKEN
			response.Error = append(response.Error, ERROR_FIREBASE_TOKEN)
			response.send(fiber.StatusBadRequest)
			return nil
		}

		err := ctx.BodyParser(&user)
		if err != nil {
			response.Msg = MSG_DEFAULT
			response.Error = append(response.Error, "Body: JSON has false format")
			response.Error = append(response.Error, err.Error())
			response.send(fiber.StatusBadRequest)
			return nil
		}
		user.FirebaseID = firebaseID

		query := "INSERT INTO users (forename, lastname, brithday, telenum,firebaseID,image)VALUES ({forename}, {lastname}, {brithday}, {telenum},{firebaseID},{image});"
		id, err := Insert(query, "id", &user)
		if err != nil {
			if strings.Contains(err.Error(), "Error 1062") {
				response.Msg = "Der Benutzer wurde bereits erstellt"
				response.Error = append(response.Error, err.Error())
				response.send(fiber.StatusBadRequest)
				return nil
			}
			response.Msg = MSG_DEFAULT
			response.Error = append(response.Error, "Database: error with INSERT")
			response.Error = append(response.Error, err.Error())
			response.send(fiber.StatusBadRequest)
			return nil
		}
		user.Id = id.(string)
		response.Content = user
		response.send(fiber.StatusCreated)*/
	return nil
	return nil
}
