package group

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("user").(user.User)

	newGroup := Group{}
	err := ctx.BodyParser(&newGroup)
	if err != nil {
		res.Msg = "Melde dich nochmal an und probier es noch einmal."
		res.AddError("Body: JSON has false format")
		res.AddError(err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}
	query := "INSERT INTO groups (name, parent) SELECT ?, ? WHERE EXISTS (SELECT 1 FROM members WHERE user = ? AND `group` = ? AND role = ?);"
	id, err := database.Insert(query, "id", newGroup.Name, newGroup.Parent, usr.Id, newGroup.Parent, MEMBER_OWNER)
	if err != nil {
		res.Msg = "Unable to create group."
		res.AddError("error on insert new group")
		res.AddError(err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}
	if !database.IsValidUUID(id) {
		res.Msg = "You dont have the permission to create a subgroup in this group."
		res.Send(fiber.StatusBadRequest)
		return nil
	}
	newGroup.Id = id

	query = "INSERT INTO members (user,`group`,role) values (?,?,?)"
	_, err = database.Insert(query, "", usr.Id, newGroup.Id, MEMBER_OWNER)
	if err != nil {
		res.Msg = "Unable to create group."
		res.Error = append(res.Error, "Cant INSERT creator as MEMBER_OWNER")
		res.Error = append(res.Error, err.Error())

		query = "DELETE FROM `groups` WHERE `groups`.id = ?;"
		_, err = database.Delete(query, newGroup.Id)
		if err != nil {
			res.Error = append(res.Error, "Cant DELETE group")
			res.AddError(err.Error())
		} else {
			res.Error = append(res.Error, "Error group was deleted")
		}
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	group, err := GetGroupByID(newGroup.Id)
	if err != nil {
		res.Msg = "Group might not exist."
		res.AddError("unable to get group by id")
		res.AddError(err.Error())
		res.Send(fiber.StatusNotFound)
		return nil
	}
	res.Content = *group
	res.Send(fiber.StatusCreated)
	return nil
}
