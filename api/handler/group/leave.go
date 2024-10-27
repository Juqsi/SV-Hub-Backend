package group

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Leave(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("user").(user.User)
	groupID := ctx.Params("groupid")

	group, err := getGroupByID(groupID)
	if err != nil {
		res.Msg = "Unable to get requested group."
		res.AddError(err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}

	member := getMemberFromGroupByUserID(*group, usr.Id)
	if member == nil {
		res.Msg = "You are not part of the requested group."
		res.AddError("usr id not in members slice of requested group")
		res.Send(fiber.StatusUnauthorized)
		return nil
	}

	if len(group.Members) == 1 {
		if err := deleteGroupById(groupID, usr); err != nil {
			res.Msg = "Unable to delete group"
			res.AddError(err.Error())
			res.Send(fiber.StatusInternalServerError)
			return nil
		}
		res.Send(200)
		return nil
	} else {
		if member.Role == "owner" && isLastOwner(group.Members, usr.Id) {
			res.Msg = "You are the last owner of the group. Please promote another member to owner before leaving the group."
			res.Send(fiber.StatusBadRequest)
			return nil
		} else {
			count, err := database.Delete("DELETE FROM members WHERE `group` = ? AND user = ?", group.Id, usr.Id)
			if err != nil {
				res.Msg = "Unable to remove from group."
				res.AddError(err.Error())
				res.Send(fiber.StatusInternalServerError)
				return nil
			}
			if count != 1 {
				res.Msg = "Logical server error occurred."
				res.AddError(fmt.Errorf("not exactly 1 row was deleted: %d were deleted", count).Error())
				res.Send(fiber.StatusInternalServerError)
				return nil
			}
			res.Send(fiber.StatusOK)
			return nil
		}
	}
	return nil
}

func isLastOwner(members []Member, currentOwnerID string) bool {
	for _, member := range members {
		if member.Id != currentOwnerID && member.Role == MEMBER_OWNER.String() {
			return false
		}
	}
	return true
}
