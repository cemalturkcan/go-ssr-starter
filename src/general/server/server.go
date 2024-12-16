package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"projectesc/src/general/config"
	"projectesc/src/general/templ"
	"projectesc/src/general/util/exitcode"
	rest2 "projectesc/src/general/util/rest"
)

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			AppName:      config.AppName,
			Prefork:      config.PreFork,
			ErrorHandler: ErrorHandler,
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
			Views:        templ.New(),
			ViewsLayout:  "layouts/main",
		}),
	}
	server.RegisterFiberRoutes()

	server.App.Use(func(c *fiber.Ctx) error {
		return rest2.ErrorRes(c, rest2.NotFound, rest2.ErrorCode[rest2.NotFound])
	})

	err := server.Listen(config.Port)
	if err != nil {
		os.Exit(exitcode.ServerStartError)
	}

	return server
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Error(err)
	code, message := rest2.Error(err)
	return rest2.ErrorRes(c, code, message)
}
