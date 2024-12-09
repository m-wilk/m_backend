package migrate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/m-wilk/w_gen/core"
	model "github.com/m-wilk/w_gen/models"
	usecase "github.com/m-wilk/w_gen/use-case"
)

func MigrateDevUserData(c *core.Core, args ...string) {
	registerUseCase := usecase.NewRegister(c.ErrorLog, c.Repository.UserRepository, c.RedisClient)
	jsonFile, err := os.Open("dev-user-data.json")
	if err != nil {
		c.ErrorLog.Fatalln(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		c.ErrorLog.Fatalln(err)
	}

	var users []struct {
		Username string         `json:"username"`
		Email    string         `json:"email"`
		Password string         `json:"password"`
		Role     model.UserRole `json:"role"`
	}
	if err := json.Unmarshal(byteValue, &users); err != nil {
		log.Fatalf("Parsing JSON error: %s", err)
	}
	for _, user := range users {
		// TODO add more fields - role, username ect.
		result, err := registerUseCase.Base(user.Email, user.Password)

		if err != nil {
			c.ErrorLog.Println(err)
			return
		}
		fmt.Printf("user with id %s inserted", result.ID)
	}
	c.InfoLog.Println("dev data inserted")
}
