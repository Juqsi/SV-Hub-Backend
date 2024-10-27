package group

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func Join(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("user").(user.User)
	invitationtoken := ctx.Params("invitationtoken")

	if !database.IsValidUUID(invitationtoken) {
		res.Msg = "Die ID ist nicht im richigen Format probiere es nochmal"
		res.Error = append(res.Error, "The id hasnt right format")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	var groups []Group
	query := "SELECT `groups`.id FROM `groups` WHERE invitationtoken = ?;"
	groups, count, err := database.Select[Group](query, invitationtoken)
	if err != nil {
		return err
	}

	if count != 1 {
		res.Msg = "Der Token ist nicht aktuell"
		res.Error = append(res.Error, "Token is invalid")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	query = "INSERT INTO members (user, `group`) values (?,?);"
	_, err = database.Insert(query, "", usr.Id, groups[0].Id)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Cant join group")
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	group, err := GetGroupByID(groups[0].Id)
	if err != nil {
		res.Msg = "Group might not exist."
		res.AddError("unable to get group by id")
		res.AddError(err.Error())
		res.Send(fiber.StatusNotFound)
		return nil
	}

	res.Content = *group
	res.Send(fiber.StatusOK)
	return nil
}

func NewInviteToken(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("user").(user.User)
	groupID := ctx.Params("groupid")

	err := generateNewInvitetoken(groupID, usr)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	group, err := GetGroupByID(groupID)
	if err != nil {
		res.Msg = "Group might not exist."
		res.AddError("unable to get group by id")
		res.AddError(err.Error())
		res.Send(fiber.StatusNotFound)
		return nil
	}

	res.Content = *group
	res.Send(fiber.StatusOK)
	return nil
}

func generateNewInvitetoken(groupID string, user user.User) error {
	query := "UPDATE `groups` SET invitationtoken = UUID() WHERE EXISTS (SELECT 1 FROM members INNER JOIN users ON members.user = users.id WHERE members.group = `groups`.id AND members.role = ? AND users.id = ? AND `groups`.id = ?);"
	effectedRows, err := database.Update(query, MEMBER_OWNER, user.Id, groupID)
	if err != nil {
		return err
	}
	if effectedRows != 1 {
		return errors.New("cant generate new token")
	}
	return nil
}
