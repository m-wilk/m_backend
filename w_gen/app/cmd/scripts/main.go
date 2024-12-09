package main

import (
	"fmt"
	"os"

	"github.com/m-wilk/w_gen/cmd/scripts/migrate"
	"github.com/m-wilk/w_gen/core"
	"github.com/m-wilk/w_gen/repository"
)

func main() {
	newCore := core.New()
	newCore.InfoLog.Println("enter scripts package")
	newCore.InitRepository(os.Getenv("DB_NAME"))
	defer repository.CloseDB()

	args := os.Args
	if len(args) <= 1 {
		fmt.Println("please add command name")
	}
	scriptName := args[1]
	scripts := RegisterScript()
	cb, ok := scripts[scriptName]
	if !ok {
		newCore.ErrorLog.Println("script not found", scriptName)
		os.Exit(1)
	}

	cb(newCore, args[2:]...)
}

func RegisterScript() map[string]func(c *core.Core, args ...string) {
	return map[string]func(c *core.Core, args ...string){
		"migrate-dev-user-data": migrate.MigrateDevUserData,
	}
}
