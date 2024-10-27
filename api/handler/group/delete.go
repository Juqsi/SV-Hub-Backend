package group

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Delete(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("user").(user.User)
	groupID := ctx.Params("groupid")

	if err := deleteGroupById(groupID, usr); err != nil {
		res.Msg = "Unable to delete group"
		res.AddError(err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}
	res.Send(200)
	return nil
}

func deleteGroupById(groupID string, user user.User) error {
	query := "DELETE g,c FROM `groups` g JOIN calendar c ON g.calendar = c.id WHERE EXISTS (SELECT 1 FROM members INNER JOIN users ON members.user = users.id WHERE `group` = ? AND members.role = ? AND users.id = ? ) and g.id = ?;"
	effectedRows, err := database.Delete(query, groupID, MEMBER_OWNER, user.Id, groupID)
	if err != nil {
		return err
	}
	if effectedRows != 2 {
		return fmt.Errorf("not exactly 2 row effected on delete: %d effected", effectedRows)
	}
	return nil
}
