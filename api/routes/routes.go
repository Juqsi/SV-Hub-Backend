package routes

import (
	"HexMaster/api/handler/user"
	"HexMaster/api/middleware"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const AcceptedDateTimeFormat string = "2006-01-02T15:04:05Z"

func SetupRoutes() {
	var app = fiber.New(fiber.Config{
		BodyLimit: 5 * 1024 * 1024, // MB
	})
	app.Use(cors.New())

	app.Use(requestid.New())

	//Monitor
	app.Get("/monitor", monitor.New(monitor.Config{Title: "SV-Hub"}))

	//logging(app)

	app.Use(middleware.Recovery)

	app.Use(middleware.ResponseBuilder)
	//login
	app.Post("/login", user.Login)
	app.Post("/registration", user.Registration)

	app.Use(middleware.Authentication)

	//user
	app.Post("/user", user.Create)
	app.Put("/user", user.Update)
	app.Get("/user/:id?", user.Get)
	app.Delete("/:id?", user.Delete)
	/***
	app.Put("/file/upload", handlerUpload)
	app.Get("/file/profile", handlerGetProfileImage)
	app.Get("/file/profile/:userid", handlerGetProfileFromUser)
	app.Get("/file/group/:groupid", handlerGetGroupImage)

	//user

	//groups
	app.Get("/groups", handlerGetGroups)

	//group
	app.Get("/group/:groupid", handlerGetGroup)
	app.Post("/group/:invitationtoken/join", handlerJoinGroup)
	app.Get("/group/:groupid/newtoken", handlerNewInvitationToken)
	app.Post("/group/:groupid/leave", handlerLeaveGroup)
	app.Post("/group/:groupid/kick/:userid", handlerKickMember)
	app.Delete("/group/:groupid", handlerDeleteGroup)
	app.Post("/group", handlerCreateGroup)
	app.Patch("/group", handlerUpdateGroup)
	app.Get("/group/:groupid/shopping", handlerGetShoppings)
	app.Get("/group/:groupid/todos", handlerGetTodoList)
	app.Get("/group/:groupid/finances", handlerGetFinances)

	//Todos
	app.Get("/todo/:todoid", handlerGetTodo)
	app.Patch("/todo/:todoid/done", handlerTodoDone)
	app.Delete("/todo/:todoid", handlerDeleteTodo)
	app.Post("/todo", handlerNewTodo)
	app.Patch("/todo", handlerUpdateTodo)

	//finances
	app.Post("/finance", handlerNewFinance)
	app.Patch("/finance", handlerUpdateFinance)
	app.Get("/finance/:groupid/reset", handlerResetFinances)
	app.Delete("/finance/:financeid", handlerDeleteFinance)

	// Shopping
	app.Post("/shopping", handlerCreateShopping)
	app.Delete("/shopping/:shoppingid", handlerDeleteShopping)
	app.Patch("/shopping/:shoppingid/toggle", handlerToggleShopping)
	app.Patch("/shopping", handlerUpdateShopping)
	app.Get("/shopping/:shoppingid", handlerGetShopping)
	*/
	fmt.Println("▶️ start server")
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("❌ cant Start Server probably port 3000 in use")
		panic(err)
	}
}
