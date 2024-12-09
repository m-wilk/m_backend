package main

import (
	"fmt"
	"os"

	"github.com/m-wilk/w_gen/core"
	"github.com/m-wilk/w_gen/handlers"
	"github.com/m-wilk/w_gen/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	newCore := core.New()
	serve(newCore)
}

func serve(c *core.Core) {
	c.InfoLog.Println("start server")

	c.InitRepository(os.Getenv("DB_NAME"))
	defer repository.CloseDB()

	c.InitRedisClient()

	e := echo.New()
	h := handlers.Handler{Core: c}
	h.Routes(e)

	c.ErrorLog.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
