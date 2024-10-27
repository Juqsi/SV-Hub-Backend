package post

import (
	"HexMaster/api/handler/group"
	"HexMaster/api/handler/user"
	"HexMaster/api/response"
	"HexMaster/database"
	"HexMaster/llama"
	"HexMaster/weaviate"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Create(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	usr := ctx.Locals("user").(user.User)

	newPost := Post{}
	err := ctx.BodyParser(&newPost)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.AddError("Body: JSON has false format")
		res.AddError(err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	grp, err := group.GetGroupByID(newPost.Group.Id)
	if err != nil {
		res.Msg = "Unable to get group information."
		res.AddError(err.Error())
		res.Send(fiber.StatusInternalServerError)
		return nil
	}

	if !group.IsUserInGroup(grp, usr.Id) {
		res.Msg = "You cant access foreign groups"
		res.AddError("task should be created in a group that you dont belong to")
		res.Send(fiber.StatusBadRequest)
		return nil
	}
	newPost.Id, err = database.Insert("INSERT INTO posts (createdAt,creator,`group`,title,content,type,parent) VALUES (?,?,?,?,?,?,?)", "id", newPost.CreatedAt, usr.Id, grp.Id, newPost.Title, newPost.Content, newPost.Type, newPost.Parent)
	if err != nil {
		res.Msg = response.MSG_DEFAULT
		res.AddError("INSERT ERROR ERROR")
		res.AddError(err.Error())
		res.Send(fiber.StatusBadRequest)
		return nil
	}

	resp, err := llama.DoRequest(llama.PROMPT_KEY_INFOS.String() + "Titel: " + newPost.Title + " \n Inhalt: " + newPost.Content)
	if err != nil {
		res.Msg = "Indexing for search failed"
		res.AddError("Cant reach llama")
		res.AddError(err.Error())
		res.Send(fiber.StatusCreated)
		return nil
	}

	lis := strings.Split(resp[1:len(resp)-1], ",")
	if len(lis) < 2 {
		res.Msg = "Indexing for search failed"
		res.AddError("Llama response has wrong format")
		res.Send(fiber.StatusCreated)
		return nil
	}
	weaviate.InsertData(lis, "posts", newPost.Id)
	res.Content = newPost
	res.Send(fiber.StatusCreated)
	return nil
}

func Like(ctx *fiber.Ctx) error {
	//TODO: Implement Like
	return nil
}
