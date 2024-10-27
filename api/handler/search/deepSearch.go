package search

import (
	"HexMaster/api/response"
	"HexMaster/llama"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func DeepSearch(ctx *fiber.Ctx) error {
	res := ctx.Locals("response").(response.Response)
	question := ctx.Query("question")
	if len(question) < 1 {
		res.Msg = "no question"
		res.Error = append(res.Error, "no question")
		res.Send(fiber.StatusBadRequest)
		return nil
	}
	resp, err := llama.DoRequest(llama.PROMPT_KEY_INFOS.String() + question)
	if err != nil {
		res.Msg = "Indexing for search failed"
		res.AddError("Cant reach llama")
		res.AddError(err.Error())
		res.Send(fiber.StatusCreated)
		return nil
	}

	lis := strings.Split(resp[1:len(resp)-1], ",")

	llama.DoRequestWithVectors(lis, question, "posts")

	return nil
}
