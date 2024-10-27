package group

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Get(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("usr").(user.User)
	groupID := ctx.Params("groupid")

	if !database.IsValidUUID(groupID) {
		res.Msg = "Die ID ist nicht im richtigen Format probiere es nochmal"
		res.Error = append(res.Error, "The id hasnt right format")
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	group, err := GetGroupByID(groupID)
	if err != nil {
		res.Msg = "Unable to get group information."
		res.AddError(err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}

	if !IsUserInGroup(group, usr.Id) {
		res.Msg = "You are not part of the requested group."
		res.AddError("user id not in members slice of requested group")
		res.Send(fiber.StatusUnauthorized)
		return nil
	}

	res.Content = group
	res.Send(fiber.StatusOK)
	return nil
}

func GetGroupByID(groupID string) (*Group, error) {
	groups, count, err := database.Select[Group]("SELECT * FROM `groups` WHERE id=?", groupID)
	if err != nil {
		return nil, err
	}
	if count != 1 {
		return nil, fmt.Errorf("no group founded with id %s", groupID)
	}
	group := groups[0]

	group.Members, err = getMembersByGroupID(groupID)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func getMembersByGroupID(groupID string) ([]Member, error) {
	members, _, err := database.Select[Member]("SELECT m.id, m.user, m.role, u.forename, u.lastname, u.username FROM members m JOIN users u ON m.user = u.id WHERE m.`group`=?", groupID)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func IsUserInGroup(group *Group, userID string) bool {
	for _, member := range group.Members {
		if member.UserID == userID {
			if member.Role != MEMBER_OWNER.String() {
				group.Invitationtoken = ""
			}
			return true
		}
	}
	return false
}

func getMemberFromGroupByUserID(group Group, userID string) *Member {
	if !IsUserInGroup(&group, userID) {
		return nil
	}
	for _, member := range group.Members {
		if member.UserID == userID {
			return &member
		}
	}
	return nil
}
