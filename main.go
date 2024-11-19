package main

import (
	"default-repo/helpers"
	dbUsecases "default-repo/libs/db/usecases"
	httpEntry "default-repo/libs/httpEntry/usecases"
	logModel "default-repo/libs/log/models"
	log "default-repo/libs/log/usecases"
	"default-repo/router"
	"encoding/json"
	"fmt"
	"runtime/debug"

	"default-repo/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	helpers.InitConfig()
	log.NewLogEntry(logrus.DebugLevel, log.GenDefaultModels())
	dbUsecases.NewGormCon()

	httpEntry.NewFacebookEntry()

}

func main() {

	app := fiber.New(fiber.Config{
		AppName:     viper.GetString("appDetail.name"),
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,

		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	return c.Status(func() int {
		// 		errDetail, ok := err.(*fiber.Error)
		// 		if ok {
		// 			return errDetail.Code
		// 		}
		// 		return c.Locals("error").(error).(*fiber.Error).Code
		// 	}()).JSON(libs.NewConnLibIns().CusErr().RecoverMsg(c, err))
		// },
	})

	// CORS
	app.Use(cors.New(cors.ConfigDefault))

	// Recover
	app.Use(recover.New(recover.Config{
		Next: func(c *fiber.Ctx) bool {
			// middleware.InitTimeStamp(c)
			return false
		},
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			c.Locals("error", fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("%v|%s", e, debug.Stack())))
		},
	}))

	// Logging
	app.Use(middleware.LoggingMiddleware)

	// authentication
	app.Use(middleware.AuthenticationMiddleware)

	// Register Router
	router.New(app)

	log.LogEntry.OnDebug(logModel.LogDebug{
		Msg: "Server is running",
	})

	err := app.Listen(fmt.Sprintf(":%s", viper.GetString("appDetail.port")))
	if err != nil {
		log.LogEntry.OnError(logModel.LogError{
			Error: err.Error(),
		})
	}
	fmt.Println("Hello World")
}
