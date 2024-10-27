package group

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Kick(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("usr").(user.User)
	groupID := ctx.Params("groupid")
	userID := ctx.Params("userid")

	if !database.IsValidUUID(groupID) {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "GroupID isnt specified or hasnt right format")
		res.Send(fiber.StatusBadRequest)
		return nil
	}
	if !database.IsValidUUID(userID) {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "UserID isnt specified or hasnt right format")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	query := "DELETE FROM `members` WHERE EXISTS ( SELECT 1 FROM members m WHERE m.`group` = ? AND m.role = ? AND m.user = ?) and members.`group` = ? and members.user = ?;"
	effectedRows, err := database.Delete(query, groupID, MEMBER_OWNER, usr.Id, groupID, userID)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, "Error with SQL Request")
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	if effectedRows != 1 {
		res.Msg = "You are not allowed to kick this member or the member is not in the group."
		res.AddError(fmt.Errorf("not exactly 1 row effected on delete: %d effected", effectedRows).Error())
		res.Send(fiber.StatusNotFound)
		return nil
	}

	err = generateNewInvitetoken(groupID, usr)
	if err != nil {
		res.Msg = "Generiere bitte einen neuen Invitationtoken"
		res.Error = append(res.Error, err.Error())
		res.Send(fiber.StatusInternalServerError)
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
