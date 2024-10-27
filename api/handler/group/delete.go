package group

import (
	"HexMaster/api/handler/user"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Delete(ctx *fiber.Ctx) error {
	response := ctx.Locals("response").(Response)
	user := ctx.Locals("user").(User)
	groupID := ctx.Params("groupid")

	if err := deleteGroupById(groupID, user); err != nil {
		response.Msg = "Unable to delete group"
		response.Error.AddError(err.Error())
		response.send(fiber.StatusInternalServerError)
		return nil
	}
	response.send(200)
	return nil
}

func deleteGroupById(groupID string, user user.User) error {
	query := "DELETE g,c  FROM `groups` g JOIN calendar c ON g.calendar = c.id WHERE EXISTS (SELECT 1 FROM members INNER JOIN users ON members.user = users.id WHERE `group` = ? AND members.role = ? AND users.id = ? ) and g.id = ?;"
	effectedRows, err := Delete(query, groupID, MEMBER_OWNER, user.Id, groupID)
	if err != nil {
		return err
	}
	if effectedRows != 2 {
		return fmt.Errorf("not exactly 2 row effected on delete: %d effected", effectedRows)
	}

	return nil
}
