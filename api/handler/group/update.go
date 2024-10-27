package group

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Update(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("usr").(user.User)

	var group Group
	err := ctx.BodyParser(&group)
	if err != nil {
		res.Msg = "Melde dich nochmal an und probier es noch einmal."
		res.Error = append(res.Error, "Body: JSON has false format")
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	oldGroup, err := getGroupByID(group.Id)
	if !isUserInGroup(oldGroup, usr.Id) {
		res.Msg = "You are not part of the group."
		res.Send(fiber.StatusUnauthorized)
		return nil
	}

	colums := make([]string, 0)
	parameters := make([]any, 0)
	if len(group.Name) > 0 {
		colums = append(colums, "name")
		parameters = append(parameters, group.Name)
	}

	if len(colums) == 0 {
		res.Msg = "Nothing to change."
		res.Send(fiber.StatusOK)
		return nil
	}

	query := "UPDATE groups SET "
	for _, colum := range colums {
		query += fmt.Sprintf("%s = ?, ", colum)
	}
	query = strings.TrimSuffix(query, ", ")
	query += " WHERE id = ?;"
	parameters = append(parameters, group.Id)

	effectedRows, err := database.Update(query, parameters...)
	if err != nil {
		res.Msg = "Unable to update group."
		res.AddError(err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}

	if effectedRows != 1 {
		res.Msg = "Unable to update group."
		res.AddError(fmt.Errorf("not exactly 1 row effected: %d were effected", effectedRows).Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}

	return nil
}
