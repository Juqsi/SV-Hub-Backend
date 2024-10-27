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

	var newPostArray []Post
	newPost := Post{}
	err := ctx.BodyParser(&newPost)
	if err != nil {
		err := ctx.BodyParser(&newPostArray)
		if err != nil {
			res.Msg = response.MSG_DEFAULT
			res.AddError("Body: JSON has false format")
			res.AddError(err.Error())
			res.Send(fiber.StatusBadRequest)
			return nil
		}
	} else {
		newPostArray[0] = newPost
	}

	for _, post := range newPostArray {

		grp, err := group.GetGroupByID(post.Group.Id)
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

		post.Id, err = database.Insert("INSERT INTO posts (creator,`group`,title,content,type,parent) VALUES (?,?,?,?,?,?,?)", "id", usr.Id, grp.Id, post.Title, post.Content, post.Type, post.Parent)
		if err != nil {
			res.Msg = response.MSG_DEFAULT
			res.AddError("INSERT ERROR ERROR")
			res.AddError(err.Error())
			res.Send(fiber.StatusBadRequest)
			return nil
		}

		resp, err := llama.DoRequest(llama.PROMPT_KEY_INFOS.String() + "Titel: " + post.Title + " \n Inhalt: " + post.Content)
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
		weaviate.InsertData(lis, "posts", post.Id)
	}
	res.Content = newPost
	res.Send(fiber.StatusCreated)
	return nil
}

func Like(ctx *fiber.Ctx) error {
	//TODO: Implement Like
	return nil
}
