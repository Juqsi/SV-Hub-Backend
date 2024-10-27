package group

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func promoteMember(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	groupID := ctx.Params("groupid")
	userID := ctx.Params("userid")
	usr := ctx.Locals("user").(user.User)
	role := ctx.Query("role")
	if !database.IsValidUUID(groupID) {
		res.Msg = response.MSG_DEFAULT
		res.AddError("GroupID isnt specified or hasnt right format")
		res.Send(fiber.StatusBadRequest)
		return nil
	}
	if !database.IsValidUUID(userID) {
		res.Msg = response.MSG_DEFAULT
		res.AddError("UserID isnt specified or hasnt right format")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	roleMember, ok := memberRoleFromString(role)
	if !ok {
		res.Msg = "Invalid role type"
		res.AddError("role cannot be converted to MemberRole")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	if changeMemberRole(groupID, userID, usr.Id, roleMember, &res) {
		return nil
	}
	res.Send(fiber.StatusOK)
	return nil
}

func changeMemberRole(groupID, userID, ownerId string, role MemberRole, res *response.Response) bool {
	query := "UPDATE `members` SET role = ? WHERE EXISTS (SELECT 1 FROM members WHERE `group` = ? AND members.role = ? AND members.user = ?) AND members.`group` = ? AND members.user = ?;"
	effectedRows, err := database.Update(query, role, groupID, MEMBER_OWNER, ownerId, groupID, userID)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.AddError("internal Database error")
		res.AddError(err.Error())
		res.Send(fiber.StatusServiceUnavailable)
		return true
	}

	if effectedRows != 1 {
		if effectedRows == 0 {
			res.Msg = response.MSG_NO_RETRY
			res.Error = append(res.Error, "No access or rights to this group or user")
			res.Send(fiber.StatusBadRequest)
			return true
		}
		res.Msg = response.MSG_DEFAULT
		res.Error = append(res.Error, strconv.Itoa(int(effectedRows))+" effected Rows")
		res.Send(fiber.StatusServiceUnavailable)
		return true
	}
	return false
}
