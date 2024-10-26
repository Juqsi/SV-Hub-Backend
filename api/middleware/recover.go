package middleware

import (
	"HexMaster/api/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"runtime/debug"
)

func Recovery(ctx *fiber.Ctx) error {
	response := response.Response{
		Ctx:     ctx,
		Access:  true,
		Error:   []string{},
		Content: nil,
		Msg:     "Es ist ein unvorhergesehener Fehler aufgetreten, wenn es häufig passiert an den Support wenden",
	}

	defer func() {
		if r := recover(); r != nil {
			var err error
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
			response.
				Error = append(response.Error, "Panic: Recovery Done")
			response.Error = append(response.Error, err.Error())
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
			response.Send(fiber.StatusInternalServerError)
		}
	}()
	return ctx.Next()
}
