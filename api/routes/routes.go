package routes

import (
	"HexMaster/api/handler/group"
	"HexMaster/api/handler/post"
	"HexMaster/api/handler/search"
	"HexMaster/api/handler/user"
	"HexMaster/api/middleware"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupRoutes() {
	var app = fiber.New(fiber.Config{
		BodyLimit: 5 * 1024 * 1024, // MB
	})
	app.Use(cors.New())

	app.Use(requestid.New())

	//Middleware
	app.Get("/monitor", monitor.New(monitor.Config{Title: "SVHub"}))
	middleware.Logging(app)
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

	//group
	app.Get("/group/:groupid", group.Get)
	app.Post("/group/:invitationtoken/join", group.Join)
	app.Get("/group/:groupid/newtoken", group.NewInviteToken)
	app.Post("/group/:groupid/leave", group.Leave)
	app.Post("/group/:groupid/kick/:userid", group.Kick)
	app.Delete("/group/:groupid", group.Delete)
	app.Post("/group", group.Create)
	app.Patch("/group", group.Update)

	//Posts
	app.Get("/post/:postid", post.Get)
	app.Delete("/post/:postid", post.Delete)
	app.Post("/post", post.Create)
	app.Patch("/post", post.Update)
	app.Post("/post/:postid/like", post.Like)

	//search
	app.Get("/search", search.DeepSearch)

	fmt.Println("▶️ start server")
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("❌ cant Start Server probably port 3000 in use")
		panic(err)
	}
}
