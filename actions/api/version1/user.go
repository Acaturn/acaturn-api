package version1

import (
	"acaturn_api/models"
	"errors"
	"fmt"
	"github.com/gobuffalo/buffalo"
)

func FetchUsersHandler(c buffalo.Context) error {
	users := []models.User{}
	err := models.DB.All(&users)
	if err != nil {
		fmt.Print(err)
		return errors.New("an error occurred")
	}
	return c.Render(200, r.JSON(users))
}
